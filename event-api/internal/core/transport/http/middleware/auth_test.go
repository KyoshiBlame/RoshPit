package middleware

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/KyoshiBlame/RoshPit/event-api/internal/core/auth"
	"github.com/KyoshiBlame/RoshPit/event-api/internal/core/domain"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

// --- мок TokenParser ---

type mockTokenParser struct {
	mock.Mock
}

func (m *mockTokenParser) ParseToken(tokenStr string) (*auth.Claims, error) {
	args := m.Called(tokenStr)
	claims, _ := args.Get(0).(*auth.Claims)
	return claims, args.Error(1)
}

// --- шпион: финальный хендлер, который запоминает, что до него дошли ---

type spy struct {
	called bool
	userID string
	role   domain.UserRole
}

func (s *spy) handler(c *gin.Context) {
	s.called = true
	s.userID = GetUserID(c)
	s.role = GetUSerRole(c)
	c.Status(http.StatusOK)
}

// perform прогоняет один запрос через роутер с заданной цепочкой хендлеров.
func perform(header string, handlers ...gin.HandlerFunc) *httptest.ResponseRecorder {
	r := gin.New()
	r.GET("/", handlers...)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	if header != "" {
		req.Header.Set("Authorization", header)
	}

	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)
	return rec
}

// withRole кладёт значение в контекст под ключом роли — имитирует отработавший Auth.
func withRole(val any) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(ctxUserRoleKey, val)
		c.Next()
	}
}

func TestRequireRoles(t *testing.T) {
	tests := []struct {
		name       string
		setRole    gin.HandlerFunc // nil — роль в контекст не кладём
		allowed    []domain.UserRole
		wantStatus int
		wantNext   bool
	}{
		{
			name:       "no role in context",
			allowed:    []domain.UserRole{domain.UserRoleAdmin},
			wantStatus: http.StatusUnauthorized,
		},
		{
			name:       "role has wrong type",
			setRole:    withRole("admin"), // string, а не domain.UserRole
			allowed:    []domain.UserRole{domain.UserRoleAdmin},
			wantStatus: http.StatusUnauthorized,
		},
		{
			name:       "role not allowed",
			setRole:    withRole(domain.UserRoleViewer),
			allowed:    []domain.UserRole{domain.UserRoleAdmin, domain.UserRoleJudge},
			wantStatus: http.StatusForbidden,
		},
		{
			name:       "role allowed",
			setRole:    withRole(domain.UserRoleJudge),
			allowed:    []domain.UserRole{domain.UserRoleAdmin, domain.UserRoleJudge},
			wantStatus: http.StatusOK,
			wantNext:   true,
		},
		{
			name:       "empty allow list denies everyone",
			setRole:    withRole(domain.UserRoleAdmin),
			allowed:    nil,
			wantStatus: http.StatusForbidden,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := &spy{}

			var chain []gin.HandlerFunc
			if tc.setRole != nil {
				chain = append(chain, tc.setRole)
			}
			chain = append(chain, RequireRoles(tc.allowed...), s.handler)

			rec := perform("", chain...)

			assert.Equal(t, tc.wantStatus, rec.Code)
			assert.Equal(t, tc.wantNext, s.called)
		})
	}
}

func TestAuth(t *testing.T) {
	validClaims := &auth.Claims{UserID: "user-1", Role: domain.UserRoleJudge}

	tests := []struct {
		name       string
		header     string
		setupMock  func(m *mockTokenParser) // nil — парсер вызываться не должен
		wantStatus int
		wantNext   bool
		wantUserID string
		wantRole   domain.UserRole
	}{
		{
			name:       "no header",
			wantStatus: http.StatusUnauthorized,
		},
		{
			name:       "wrong scheme",
			header:     "Basic abc",
			wantStatus: http.StatusUnauthorized,
		},
		{
			name:   "parser returns error",
			header: "Bearer bad",
			setupMock: func(m *mockTokenParser) {
				m.On("ParseToken", "bad").Return(nil, errors.New("expired")).Once()
			},
			wantStatus: http.StatusUnauthorized,
		},
		{
			name:   "valid token",
			header: "Bearer abc",
			setupMock: func(m *mockTokenParser) {
				m.On("ParseToken", "abc").Return(validClaims, nil).Once()
			},
			wantStatus: http.StatusOK,
			wantNext:   true,
			wantUserID: "user-1",
			wantRole:   domain.UserRoleJudge,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := &mockTokenParser{}
			if tc.setupMock != nil {
				tc.setupMock(m)
			}
			s := &spy{}

			rec := perform(tc.header, Auth(m), s.handler)

			assert.Equal(t, tc.wantStatus, rec.Code)
			assert.Equal(t, tc.wantNext, s.called)
			assert.Equal(t, tc.wantUserID, s.userID)
			assert.Equal(t, tc.wantRole, s.role)

			if tc.setupMock == nil {
				m.AssertNotCalled(t, "ParseToken", mock.Anything)
			}
			m.AssertExpectations(t)
		})
	}
}

// Связка Auth → RequireRoles: проверяет, что Auth кладёт роль туда,
// откуда её читает RequireRoles. Именно такой тест ловит путаницу ключей.
func TestAuthThenRequireRoles(t *testing.T) {
	m := &mockTokenParser{}
	m.On("ParseToken", "abc").
		Return(&auth.Claims{UserID: "u1", Role: domain.UserRoleAdmin}, nil).Once()
	s := &spy{}

	rec := perform("Bearer abc", Auth(m), RequireRoles(domain.UserRoleAdmin), s.handler)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, s.called)
	m.AssertExpectations(t)
}
