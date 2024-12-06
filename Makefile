build:
	@go build -o bin/go_bank_jwt

run: build
	@./bin/go_bank_jwt

test: 
	@go test -v	./...

