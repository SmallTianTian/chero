install:
	@go install ./cmd/chero

generate: install
	@go generate ./...