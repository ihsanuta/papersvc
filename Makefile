NAMESPACE = `echo papersvc`
BUILD_TIME = `date +%FT%T%z`
BUILD_VERSION = `git describe --tag`
COMMIT_HASH = `git rev-parse --short HEAD`

.PHONY: run
run: 
	@go run app.go