## WAT

This boilerplate should help you to replace a bash script as fast as possible - use it as a copy&paste starter, change the strings and start right away with your business logic.

It should be transparent, you should be able to understand what is happening and go as deep into the matter as you like.
It should be easy to understand, all implementations in cmd/ are small and specific so you can junk them together and chain them.

It should help you
 - be documented .. right.
 - handling script parameters and options using the [cobra](https://github.com/spf13/cobra) lib
 - running cli commands over ssh utilizing your ssh-agent/privkey/password [sshapi](https://github.com/EugenMayer/go-sshclient)
 - running cli commands on the shell with proper stdin/stdout/err handling using the [sshapi](https://github.com/EugenMayer/go-exec)
 - transferring files from and to remote server using scp [sshapi](https://github.com/EugenMayer/go-sshclient)
 - help you building in a CI or locally without any dependencies using the `makefile` and `Dockerfile`
 - Include IntelliJ run configuration to run/debug the tasks right away
 
## Test it

I did ship `dist/mycli*` in the repo for convinience reasons for now, so you do neither need to `build` yourself nor us `curl`
Generally depending on your use, use `mycli-macos` or `mycli-linux`

### Exec

```bash
  dist/mycli-macos myexec --cmd="echo hi"
```

### SSH
If you want to the the ssh command, just use the local `docker-compose.yml` to start a local `ssh node`

```
  docker-compose up -d
  dist/mycli-macos myssh --host=localhost --port=2301 --key=test/sshkeys/id_rsa
```

### SCP
If you want to the the ssh command, just use the local `docker-compose.yml` to start a local `ssh node`

```
  docker-compose up -d
  dist/mycli-macos myscp --host=localhost --port=2301 --key=test/sshkeys/id_rsa --file=test/dummytestfile
```

## Build it yourself

If you have docker, make it easy for yourself, you find the binaries in dist/ locally after that

```bash
   make build-docker
```

If you feel in need do this for a local manual build

```bash
   make build
   # or of course
   env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags netgo -o dist/mycli-linux mycli.go
```
 
## TODO

- maybe implement into http://yeoman.io/