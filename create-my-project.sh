#!/bin/zsh

set -e

# The name of your cli like "mytool"
HOST=github.com
STRIP_SSH=false

for arg in "$@"
do
    case $arg in
        -u=*|--username=*)
            NAMESPACE="${arg#*=}"
            ;;
        -p=*|--project=*)
            NAME="${arg#*=}"
            ;;
        -h=*|--host=*)
            HOST="${arg#*=}"
            ;;
        -s|--strip-ssh)
            echo "Removing all ssh/scp dependencies - stripping down"
            STRIP_SSH=true
            ;;
        *)
            echo "Unknown argument: \"${arg}\"" >&2
            exit 1
            ;;
    esac
done

# usually the github username or the SCM namespace
# your SCM url
PROJECT_FQDN="$HOST/$NAMESPACE/$NAME"
echo "Project FQDN: $PROJECT_FQDN"

DEST=$(mktemp -d "/tmp/$NAME-shell.XXXXXXXXX")
cp -fr ./* "$DEST"

cd $DEST

rm create-my-project.sh LICENSE
mv mycli.go "$NAME".go
mv go-shell-cli-quickstarter.iml go-"$NAME".iml

mv README.your.example.md README.md

sed -i -e "s|github.com/eugenmayer/go-shell-cli-quickstarter|$PROJECT_FQDN|g" go.mod
sed -i -e "s|github.com/eugenmayer/go-shell-cli-quickstarter|$PROJECT_FQDN|g" $NAME.go
sed -i -e "s|github.com/eugenmayer/go-shell-cli-quickstarter|$PROJECT_FQDN|g" cmd/*.go

sed -i -e "s|mycli|$NAME|g" Makefile

if [ "$STRIP_SSH" = "true" ]; then
  rm -f cmd/myscp.go
  rm -f cmd/myssh.go
  rm -f docker-compose.yml
  rm -f go.sum
  rm -fr test
fi

go mod tidy
echo "You find your cli project named $NAME under $DEST"
ls -la
