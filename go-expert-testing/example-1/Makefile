.PHONY: test
test:
	@go test -v .

.PHONY: coverage
coverage:
	@go test -v -coverprofile=coverage.out

.PHONY: coverage-html
coverage-html:
	@go tool cover -html=coverage.out

.PHONY: bench
bench:
	@go test -bench=. -run=^#

.PHONY: fuzz
fuzz:
	@go test -fuzz=. -run=^# -fuzztime=5s