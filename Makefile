VERSION  ?= 0.1
BINARIES := bin/convert-latlong bin/unpack-metatile bin/convert-path
GITHASH  := $(shell git rev-parse --short HEAD)
FULLVER  := $(VERSION)-git.$(shell git rev-parse --abbrev-ref HEAD).$(shell git rev-parse --short HEAD)
LDFLAGS  := -ldflags "-X main.version=$(FULLVER)"

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

$(BINARIES): test
	go build -v $(LDFLAGS) -o $@ cmd/$(notdir $@)/*.go

.PHONY: clean
clean:
	rm -f bin/*
	rm -f ./pprof
