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

CLOUDINIT=github.com/0-complexity/cloud-init-server

go get -v -d $CLOUDINIT

cd $GOPATH/src/$CLOUDINIT

git fetch origin "${branch}:${branch}" || true
git checkout "${branch}" || true

mkdir -p build/bin
mkdir -p build/etc/cloud-init
CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o build/bin/cloud-init-server .

mkdir -p /tmp/archives/
tar -czf "/tmp/archives/cloud-init-server-${branch}.tar.gz" -C $GOPATH/src/$CLOUDINIT/build .
