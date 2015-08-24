#!/usr/bin/env bash

if [ ! -f install.sh ]; then
	echo 'test must be run within its container folder'
	exit 1
fi

CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"

go test command 

export GOPATH="$OLDGOPATH"

echo 'test finished!'
