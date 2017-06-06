# Development Setup of a Zero-OS Cluster

You can either:
- [Setup your cluster using an installation script](#automated-setup) (recommended)
- [Or manually setup the cluster](#manual-setup) (not supported)

Once done, the last step is documented in [Boot your Zero-OS nodes](#boot-nodes).

## Automated setup

This is the recommended and currently the only supported option to setup a Zero-OS cluster.

In order to have a full Zero-OS cluster you need:
1. JumpScale9 development Docker container
2. Install the Zero-OS orchestrator into the container

Create the Docker container with JumpScale9 by following the documentation at https://github.com/Jumpscale/developer#jumpscale-9.

Once you have your JumpScale9 Docker container running, SSH into it and install the rest using the [`installgrid.sh`](../../scripts/installgrid.sh) script.

This script takes 3 parameters:
- BRANCH: 0-orchestrator development branch
- ZEROTIERNWID: ZeroTier network id
- ZEROTIERTOKEN: ZeroTier API token

So:
```
BRANCH="master"
ZEROTIERNWID=""
ZEROTIERTOKEN=""
./installgrid.sh $BRANCH $ZEROTIERNWID $ZEROTIERTOKEN
```
