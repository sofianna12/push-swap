.PHONY: build test coverage fmt vet lint check clean audit benchmark

# Build both binaries
build:
	go build -o push-swap ./cmd/push-swap
	go build -o checker ./cmd/checker

# Run all tests with race detector and verbose output
test:
	go test -race -v ./...

# Run tests with coverage report and HTML output
coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html

# Fix formatting in place
fmt:
	gofmt -w .

# Static analysis
vet:
	go vet ./...

# Extended lint with golangci-lint
lint:
	golangci-lint run ./...

# Run fmt + vet + lint together
check: fmt vet lint

# Remove build artifacts
clean:
	rm -f push-swap checker coverage.out coverage.html

# Run the official audit checklist (requires built binaries)
audit: build
	@echo "==> Audit: no args"
	@./push-swap 2>&1 | wc -c | grep -q "^0$$" && echo "  PASS: no output" || echo "  FAIL: unexpected output"
	@echo "==> Audit: non-integer"
	@./push-swap "0 one 2 3" 2>&1 | grep -q "Error" && echo "  PASS: Error printed" || echo "  FAIL: no Error"
	@echo "==> Audit: duplicate"
	@./push-swap "1 2 2 3" 2>&1 | grep -q "Error" && echo "  PASS: Error printed" || echo "  FAIL: no Error"
	@echo "==> Audit: n=6 op count"
	@COUNT=$$(./push-swap "2 1 3 6 5 8" | wc -l); \
	 [ "$$COUNT" -lt 9 ] && echo "  PASS: $$COUNT ops (<9)" || echo "  FAIL: $$COUNT ops (>=9)"
	@echo "==> Audit: pipe to checker"
	@ARG="4 67 3 87 23"; RESULT=$$(./push-swap "$$ARG" | ./checker "$$ARG"); \
	 [ "$$RESULT" = "OK" ] && echo "  PASS: OK" || echo "  FAIL: $$RESULT"

# Measure operation counts for n=5, n=6, n=100
benchmark: build
	@echo "==> n=5 (target: <12 ops)"
	@for i in $$(seq 5); do \
	   ARG=$$(shuf -i 1-100 -n 5 | tr '\n' ' '); \
	   COUNT=$$(./push-swap "$$ARG" | wc -l); \
	   RESULT=$$(./push-swap "$$ARG" | ./checker "$$ARG"); \
	   echo "  Run $$i: $$COUNT ops → $$RESULT | $$ARG"; \
	 done
	@echo "==> n=6 (target: <9 ops)"
	@for i in $$(seq 5); do \
	   ARG=$$(shuf -i 1-100 -n 6 | tr '\n' ' '); \
	   COUNT=$$(./push-swap "$$ARG" | wc -l); \
	   RESULT=$$(./push-swap "$$ARG" | ./checker "$$ARG"); \
	   echo "  Run $$i: $$COUNT ops → $$RESULT | $$ARG"; \
	 done
	@echo "==> n=100 (target: <700 ops)"
	@for i in $$(seq 5); do \
	   ARG=$$(shuf -i 1-10000 -n 100 | tr '\n' ' '); \
	   COUNT=$$(./push-swap "$$ARG" | wc -l); \
	   RESULT=$$(./push-swap "$$ARG" | ./checker "$$ARG"); \
	   echo "  Run $$i: $$COUNT ops → $$RESULT"; \
	 done
