#!/bin/bash
set -e

if ! grep 'xenial universe' /etc/apt/sources.list; then
    echo "deb http://archive.ubuntu.com/ubuntu xenial universe" >> /etc/apt/sources.list
fi

apt-get update
apt-get install -y git curl net-tools

export JSBRANCH="8.2.0"
cd /tmp
curl -k https://raw.githubusercontent.com/Jumpscale/jumpscale_core8/$JSBRANCH/install/install.sh?$RANDOM > install.sh
bash install.sh

pip3 install g8core

js 'x = j.tools.cuisine.local; x.apps.atyourservice.install()'
js 'x = j.tools.cuisine.local; x.apps.caddy.build(); x.apps.caddy.install()'

cd /
mkdir /tmp/archives
tar -czf /tmp/archives/grid.tar.gz --exclude tmp/archives --exclude sys --exclude dev --exclude proc *

