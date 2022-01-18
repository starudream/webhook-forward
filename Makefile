GO      ?= GO111MODULE=on go
VERSION ?= $(shell git describe --exact-match --tags HEAD 2>/dev/null || echo "latest")
GITHASH ?= $(shell git rev-parse --short HEAD)

LDFLAGS := -s -w
LDFLAGS += -X "github.com/go-sdk/lib/app.VERSION=$(VERSION)"
LDFLAGS += -X "github.com/go-sdk/lib/app.GITHASH=$(GITHASH)"
LDFLAGS += -X "github.com/go-sdk/lib/app.BUILT=$(shell date +%FT%T%z)"

build:
	@$(MAKE) tidy
	CGO_ENABLED=0 $(GO) build -ldflags '$(LDFLAGS)' -o bin/app cmd/*

run:
	@$(MAKE) tidy
	CGO_ENABLED=1 $(GO) run -race -ldflags '$(LDFLAGS)' cmd/*

tidy:
	$(GO) mod tidy -go=1.16 && $(GO) mod tidy -go=1.17

upx:
	upx bin/app
