#!/bin/zsh

set -e
NAME=${1}

DEST=$(mktemp -d "/tmp/$NAME-shell.XXXXXXXXX")
cp -fr ./* "$DEST"

cd $DEST
pwd

rm create-my.sh
mv mycli.go $NAME.go

ls -la
