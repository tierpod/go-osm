.PHONY: test
test:
	go test -cover ./...

.PHONY: lint
lint:
	find ./ -type f -name '*.go' | xargs gofmt -l -e
	go vet ./...

.PHONY: doc
doc:
	godoc -http :6060
