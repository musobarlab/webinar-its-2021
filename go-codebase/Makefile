.PHONY : test docker coverage clean format lint

PACKAGES = $(shell go list ./... | grep -v -e . | tr '\n' ',')

# Docker Build
docker: Dockerfile
	docker build -t go-codebase:latest .

format:
	find . -name "*.go" -not -path "./vendor/*" -not -path ".git/*" | xargs gofmt -s -d -w

# binary
go-codebase-osx:
	GOOS=darwin GOARCH=amd64 go build -ldflags '-s -w' -o $@ gitlab.com/Wuriyanto/go-codebase/cmd/go-codebase

go-codebase-linux:
	GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o $@ gitlab.com/Wuriyanto/go-codebase/cmd/go-codebase

go-codebase64.exe:
	GOOS=windows GOARCH=amd64 go build -ldflags '-s -w' -o $@ gitlab.com/Wuriyanto/go-codebase/cmd/go-codebase

go-codebase32.exe:
	GOOS=windows GOARCH=386 go build -ldflags '-s -w' -o $@ gitlab.com/Wuriyanto/go-codebase/cmd/go-codebase

go-codebase-windows: go-codebase64.exe go-codebase32.exe

go-codebase: go-codebase-osx go-codebase-linux go-codebase-windows

# unit test & calculate code coverage
test:
	if [ -f coverage.txt ]; then rm coverage.txt; fi;
	@echo ">> running unit test and calculate coverage"
	@go test ./... -cover -coverprofile=coverage.out -covermode=count -coverpkg=$(PACKAGES)
	@go tool cover -func=coverage.out

# Coverage with HTML output
cover-html:
	if [ -f coverages/coverage-all.out ]; then rm coverages/coverage-all.out; fi;
	@echo "mode: count" > coverages/coverage-all.out
	@go test ./... -cover -coverprofile=coverages/coverage-all.out -covermode=count -coverpkg=$(PACKAGES)
	go tool cover -html=coverages/coverage-all.out -o coverages/index.html

cover: clean-coverfile coverage.txt

lint-prepare: 
	@echo "Installing golangci-lint"
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.23.8

lint:
	./bin/golangci-lint run \
		--exclude-use-default=false \
		--enable=golint \
		--enable=gocyclo \
		--enable=goconst \
		--enable=unconvert \
		--disable=errcheck \
		--enable=dupl \
		./...

clean-coverfile:
	if [ -f coverage.txt ]; then rm coverage.txt; fi;

coverage.txt:
	@echo ">> get coverprofile"
	go test ./... -race -cover -coverprofile=coverage.txt -covermode=atomic; echo "coverage.txt generated"

clean:
	[ -f go-codebase-osx ] && rm go-codebase-osx || true
	[ -f go-codebase-linux ] && rm go-codebase-linux || true
	[ -f go-codebase32.exe ] && rm go-codebase32.exe || true
	[ -f go-codebase64.exe ] && rm go-codebase64.exe || true
	[ -f coverage.txt ] && rm coverage.txt || true
	[ -f coverages/coverage-all.out ] && rm coverages/coverage-all.out || true
	[ -f coverages/coverage.out ] && rm coverages/coverage.out || true
	[ -f coverages/index.html ] && rm coverages/index.html || true