from JumpScale import j


def input(job):
    if job.model.args.get("mountpoint", "") != "":
        raise j.exceptions.Input("Mountpoint should not be set as input")


def get_filesystem(service):
    nodeservice = service.parent.parent.parent
    poolname = service.parent.parent.name
    fsname = str(service.parent.model.data.name)
    node = j.sal.g8os.get_node(
        addr=nodeservice.model.data.redisAddr,
        port=nodeservice.model.data.redisPort,
        password=nodeservice.model.data.redisPassword or None,
    )
    pool = node.storagepools.get(poolname)
    return pool.get(fsname)


def install(job):
    name = str(job.service.model.data.name)
    fs = get_filesystem(job.service)
    try:
        fs.get(name)
    except ValueError:
        fs.create(name)


def delete(job):
    name = str(job.service.model.data.name)
    fs = get_filesystem(job.service)
    snapshot = fs.get(name)
    snapshot.delete()


def rollback(job):
    name = str(job.service.model.data.name)
    fs = get_filesystem(job.service)
    snapshot = fs.get(name)
    snapshot.rollback()
