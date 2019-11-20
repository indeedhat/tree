default: cover

cover:
	@go test -coverprofile=cover.out
	@go tool cover -html=cover.out
	@go tool cover -func=cover.out
	@rm cover.out

test:
	@go test -cover

