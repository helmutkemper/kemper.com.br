GO ?= go

.PHONY: tidy
##
tidy:
	@$(GO) mod tidy

.PHONY: build
##
build:
	@$(GO) build -o ../build/site.so ./main.go
