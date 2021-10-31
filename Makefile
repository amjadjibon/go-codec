VERSION = "0.0.1"
version:
	@echo $(version)>VERSION
	@echo "package constant\n\n//Version constant of the service\nconst Version = \"$(version)\"">constant/version.go
	@git add VERSION
	@git add constant/version.go
	@git commit -m "v$(version)"
	@git tag -a "v$(version)" -m "v$(version)"
	@git push origin
	@git push origin "v$(version)"

test:
	go test -count=1 -race ./... -v

gofmt:
	go fmt ./...

gocyclo:
	gocyclo .

update-module:
	go get -v github.com/vmihailenco/msgpack/v5

tidy:
	go mod tidy -v

vendor:
	go mod vendor

staticcheck:
	go install honnef.co/go/tools/cmd/staticcheck@latest
	staticcheck -go 1.7 ./...

install-module:
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.42.1

cover:
	go test ./... -coverprofile=cover.out -v

cover-html:
	go tool cover -html=cover.out

