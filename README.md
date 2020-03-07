## WAT

This boilerplate should help you to replace a bash script as fast as possible - use it as a copy&paste starter, change the strings and start right away with your business logic.

It should be transparent, you should be able to understand what is happening and go as deep into the matter as you like.
It should be easy to understand, all implementations in cmd/ are small and specific so you can junk them together and chain them.

It should help you
 - **get the job done. Period.** (not teach you go concepts all over the place)
 - be documented .. right.
 - handling script parameters and options using the [cobra](https://github.com/spf13/cobra) lib
 - running cli commands on the shell with proper stdin/stdout/err handling easily: [go-exec](https://github.com/EugenMayer/go-exec)
 - running cli commands over ssh utilizing your ssh-agent/privkey/password easily: [go-sshclient](https://github.com/EugenMayer/go-sshclient)
 - transferring files from and to remote server using scp easily: [go-sshclient](https://github.com/EugenMayer/go-sshclient)
 - Include IntelliJ run configuration to run/debug the tasks right away
 

**Why not learning all the deeper concepts?**
Well you could .. but that stops most people from using golang over bash in the smaller, daily projects.
So you will learn the deeper parts every time you write something, part by part - but that happens as a side-track, while
you actually get your job done
 
## Get started

You can generate a new project for yourself using this boilerplate

This will generate a new shell-project with the name 'mycliname' in `/tmp/mycliname` - you can open the folder using your
IDE and start developing right away

```bash
./create-my-project.sh --project=mycliname --username=ghusername --strip-ssh

# or with private SCM
./create-my-project.sh --project=mycliname --username=ghusername  --host=ourprivate-scm.tld --strip-ssh
```

You most probably want to `mv` the project from `/tmp` into a folder for your scm prjects.

Where `cliname` is the name of your cli-program you plan, `ghusername` is your namespace, usually your github username
and if needed the third param your private SCM domain ( without scheme ). `--strip-ssh` means that the examples using
`ssh` and `scp` are removed including it's dependencies to have you start with a slim, stripped down project

That's read the shell output and your already can start creating.

## Included examples

I did ship `dist/mycli*` in the repo for convenience reasons for now, so you do neither need to `build` yourself nor us `curl`
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
   make build
```
 
## TODO

- maybe implement into http://yeoman.io/
