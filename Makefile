.PHONY: run-infra proto deps tidy build

run-infra:
	docker-compose up -d postgres redis kafka nginx

proto:
	protoc --go_out=. --go_opt=paths=source_relative \\
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \\
	proto/match/v1/match.proto

tidy:
	cd event-api && go mod tidy
	cd live-score-engine && go mod tidy
	cd rating-worker && go mod tidy
