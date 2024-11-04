run:
	@go run ./cmd/workerpool/main.go

build:
	@go build -o workerpool ./cmd/workerpool/main.go

test:
	@cd ./pkg/workerpool && go test