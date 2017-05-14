#!/bin/bash
set -e
source $(dirname $0)/tools.sh
ensure_go

branch="master"
echo $1

if [ "$1" != "" ]; then
    branch="$1"
fi

apt-get update
apt-get install -y curl git

BLOCKSTOR=$GOPATH/src/github.com/g8os/blockstor

go get -u -v -d github.com/g8os/blockstor/nbdserver
go get -u -v -d github.com/g8os/blockstor/cmd/copyvdisk

cd $BLOCKSTOR
mkdir -p build/bin

git fetch origin || true
git checkout -B "${branch}" origin/${branch} || true

cd $BLOCKSTOR/nbdserver
CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o ../build/bin/nbdserver

cd $BLOCKSTOR/cmd/copyvdisk
CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o ../../build/bin/copyvdisk

mkdir -p /tmp/archives/
tar -czf "/tmp/archives/blockstor-${branch}.tar.gz" -C $BLOCKSTOR/build .
