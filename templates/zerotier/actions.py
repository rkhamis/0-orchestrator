def _get_client(parent):
    return j.clients.g8core.get(host=parent.model.data.redisAddr,
                                port=parent.model.data.redisPort or 6379,
                                password=parent.model.data.redisPassword or '')


def install(job):
    service = job.service
    client = _get_client(service.parent)
    client.zerotier.join(service.model.data.networkID)


def delete(job):
    service = job.service
    client = _get_client(service.parent)
    client.zerotier.leave(service.model.data.networkID)
