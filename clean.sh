#!/usr/bin/env bash

if [ ! -f clean.sh ]; then
	echo 'clean must be run within its container folder'
	exit 1
fi

CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"

go clean -i -x service-center 

export GOPATH="$OLDGOPATH"

echo 'clean finished!'
