build:
	go build tasty

test:
	go test ./... --cover

release:
	env GOOS=darwin GOARCH=amd64 go build tasty && mv tasty tasty-darwin-amd64
	env GOOS=darwin GOARCH=arm64 go build tasty && mv tasty tasty-darwin-arm64
	env GOOS=linux GOARCH=amd64 go build tasty && cp tasty tasty-linux-amd64 && mv tasty tasty-linux-x86_64
	env GOOS=linux GOARCH=arm64 go build tasty && cp tasty tasty-linux-arm64 && mv tasty tasty-linux-aarch64
