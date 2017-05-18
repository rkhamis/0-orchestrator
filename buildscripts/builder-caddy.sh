#!/bin/sh
set -e

TARGET=/tmp/caddy
url="https://caddyserver.com/download/linux/amd64"

rm -rf $TARGET
mkdir -p $TARGET
mkdir -p $TARGET/bin
wget "$url" -O "${TARGET}/caddy.tar.gz"
tar xf $TARGET/caddy.tar.gz -C $TARGET/bin caddy
mkdir -p /tmp/archives/
tar czf /tmp/archives/caddy.tar.gz -C $TARGET bin