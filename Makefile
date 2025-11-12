fmt:
	go fmt ./...
	goimports -w .
lint: fmt
	go vet ./cmd/api/    

