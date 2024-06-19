build:
	@go build -o bin/rssagg.exe ./main.go

run: build
	@./bin/rssagg.exe

test:
	@go test -v ./...