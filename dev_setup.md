# Grid REST API dev environment setup

## 1. Install jumpscale
```shell
cd $TMPDIR
rm -f install.sh
export JSBRANCH="8.2.0"
curl -k https://raw.githubusercontent.com/Jumpscale/jumpscale_core8/$JSBRANCH/install/install.sh?$RANDOM > install.sh
bash install.sh
```

## 2. Install g8core python client
```shell
pip3 install g8core
```

## 3. Start jumpscale and create a AYS repository
This command will start AYS in tmux:
```shell
ays start --bind 0.0.0.0 --debug
```

Create the AYS repository:
```shell
ays repo create --name grid --git http://github.com/user/repo
```

## 4. Clone ays templates
```
cd /opt/code/
git clone https://github.com/g8os/grid/
cd grid
git checkout {branch}
ays reload
```
Replace {branch} with the version you want to have

## 5. Install auto node discovery service:
Copy this in the blueprints directory of you AYS repo:
```yaml
bootstrap.g8os__grid1:

actions:
  - action: install
```
Then execute it:
```shell
ays blueprint
ays run create --follow
```

## 6. Start G8OS

https://github.com/g8os/home/blob/1.1.0-alpha/docs/development/qemu.md
