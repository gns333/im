#!/usr/bin/env bash

if [ ! -f install.sh ]; then
	echo 'install must be run within its container folder'
	exit 1
fi

CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"

gofmt -w src

go install service-center

go install session-node

export GOPATH="$OLDGOPATH"

echo 'install finished!'
