def input(job):
    # make sure we always consume all the filesystems used in the mounts property
    args = job.model.args
    mounts = args.get('mounts', [])
    if 'filesystems' not in args:
        args['filesystems'] = []
    filesystems = args['filesystems']
    for mount in mounts:
        if mount['filesystem'] not in filesystems:
            args['filesystems'].append(mount['filesystem'])

    return args


def install(job):
    job.service.model.data.status = "halted"

def start(job):
    from JumpScale.sal.g8os.Container import Container

    container = Container.from_ays(job.service)
    container.start()

    if container.is_running():
        job.service.model.data.status = "running"
    else:
        raise j.exceptions.RuntimeError("container didn't started")


def stop(job):
    from JumpScale.sal.g8os.Container import Container
    container = Container.from_ays(job.service)
    container.stop()

    if not container.is_running():
        job.service.model.data.status = "halted"
    else:
        raise j.exceptions.RuntimeError("container didn't stopped")


def monitor(job):
    service = job.service
    from JumpScale.sal.g8os.Container import Container

    container = Container.from_ays(job.service)
    container.stop()

    if service.model.actionsState['install'] == 'ok':
        if not container.is_running():
            try:
                job.logger.warning("container {} not running, trying to restart".format(service.name))
                service.model.dbobj.state = 'error'
                container.start()

                if container.is_running():
                    service.model.dbobj.state = 'ok'
            except:
                job.logger.error("can't restart container {} not running".format(service.name))
                service.model.dbobj.state = 'error'
