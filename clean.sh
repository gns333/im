#!/usr/bin/env bash

if [ ! -f clean.sh ]; then
	echo 'clean must be run within its container folder'
	exit 1
fi

CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"

go clean -i service-center

go clean -i session-node

rm -rf ./bin/

rm -rf ./pkg/

export GOPATH="$OLDGOPATH"

echo 'clean finished!'
