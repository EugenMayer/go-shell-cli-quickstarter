init:
	go mod tidy
	go mod verify
	go mod vendor

update:
	go get -u
	go mod tidy

build:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags netgo -o dist/mycli-linux mycli.go
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -tags netgo -o dist/mycli-macos mycli.go

test:
	env CGO_ENABLED=0
