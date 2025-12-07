.PHONY: all
all: build

.PHONY: gen
gen:
	go generate ./...

.PHONY: build
build: gen
	go build ./...

.PHONY: update
update:
	@echo "Updating protocol definitions..."
	@mkdir -p cmd/cdpgen/protodef
	curl -sSL https://github.com/ChromeDevTools/devtools-protocol/raw/master/json/browser_protocol.json -o cmd/cdpgen/protodef/browser_protocol.json
	curl -sSL https://github.com/ChromeDevTools/devtools-protocol/raw/master/json/js_protocol.json -o cmd/cdpgen/protodef/js_protocol.json
	@echo 'Done. Run "make gen" to regenerate bindings.'

.PHONY: test
test:
	go test ./...

.PHONY: test-race
test-race:
	go test -race ./...

.PHONY: test-browser
test-browser:
	go test . ./session -browser
	go test . ./session -browser -race

.PHONY: lint
lint:
	go vet ./...

.PHONY: fmt
fmt:
	gofmt -s -w .
	gofumpt -w .

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: clean
clean:
	go clean ./...
	rm -f cmd/cdpgen/cdpgen
