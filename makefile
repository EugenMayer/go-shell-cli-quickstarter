ci-build:
	rm -rf dist
	mkdir -p dist
	dep ensure -v
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags netgo -o dist/mycli-linux mycli.go
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -tags netgo -o dist/mycli-macos mycli.go
	# inject version if it exists
	if [ ! -z "${VERSION}" ]; then mv dist/mycli-linux  dist/mycli-linux-${VERSION}; fi
	if [ ! -z "${VERSION}" ]; then mv dist/mycli-macos  dist/mycli-macos-${VERSION}; fi
	chmod +x dist/*

# local docker based build, like in concourse or a ci
build-docker:
	#docker stop drupalwikihost > /dev/null 2>&1 || true
	docker build --rm -t localbuild .
	# copy the dist to our local dist folder
	docker run -v `pwd`/dist:/dist --rm --entrypoint /bin/sh localbuild -c 'cp /go/src/github.com/eugenmayer/go-antibash-boilerplate/dist/mycli* /dist/'

build:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags netgo -o dist/mycli-linux mycli.go
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -tags netgo -o dist/mycli-macos mycli.go

init:
	dep ensure