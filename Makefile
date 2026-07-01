.PHONY: ci fmt fuzz test vet

GO_MODULES := $(shell find . -name go.mod -exec dirname {} \;)

ci: fmt vet test

fmt:
	@test -z "$$(gofmt -l $$(find . -name '*.go'))"

test:
	@set -e; for dir in $(GO_MODULES); do \
		echo "== $$dir =="; \
		(cd "$$dir" && go test ./...); \
	done

vet:
	@set -e; for dir in $(GO_MODULES); do \
		echo "== $$dir =="; \
		(cd "$$dir" && go vet ./...); \
	done

fuzz:
	cd property_based_tests && go test -fuzz=FuzzRomanNumeralRoundTrip -fuzztime=5s
