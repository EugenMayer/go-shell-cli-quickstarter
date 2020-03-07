#!/bin/zsh

set -e

NAME=${1}
NAMESPACE=${2}
HOST=${3:-github.com}

PROJECT_FQDN="$HOST/$NAMESPACE/$NAME"
echo "Project FQDN: $PROJECT_FQDN"

DEST=$(mktemp -d "/tmp/$NAME-shell.XXXXXXXXX")
cp -fr ./* "$DEST"

cd $DEST

rm create-my.sh
mv mycli.go $NAME.go
mv README.your.example.md README.md

sed -i -e "s|github.com/eugenmayer/go-shell-cli-quickstarter|$PROJECT_FQDN|g" go.mod
sed -i -e "s|github.com/eugenmayer/go-shell-cli-quickstarter|$PROJECT_FQDN|g" $NAME.go
sed -i -e "s|github.com/eugenmayer/go-shell-cli-quickstarter|$PROJECT_FQDN|g" cmd/*.go

sed -i -e "s|mycli|$NAME|g" Makefile

echo "You find your cli project named $NAME under $DEST"
