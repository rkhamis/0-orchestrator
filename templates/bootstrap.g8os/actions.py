
def bootstrap(job):
    # discover the node that send the event
    if not hasattr(job.model, 'request'):
        raise j.exceptions.RuntimeError("can't inspect http request")

    ip, _ = job.model.request.ip
    job.logger.info("reverse ardb loopup on IP: {}".format(ip))
    mac = j.sal.nettools.getMacAddressForIp(ip)

    # create and install the node.g8os service
    client_actor = job.service.aysrepo.actorGet('g8os_client')
    node_actor = job.service.aysrepo.actorGet('node.g8os')
    grid_config = job.service.aysrepo.servicesFind(actor='grid_config')[0]
    client_args = {
        'redisAddr': ip,
    }

    job.logger.info("create g8os_client service {}".format(mac))
    client_actor.serviceCreate(instance=mac, args=client_args)

    node_args = {
        'id': mac,
        'status':'running',

        'client': mac,
        'gridConfig': grid_config.name,
    }
    job.logger.info("create node.g8os service {}".format(mac))
    node_actor.serviceCreate(instance=mac, args=node_args)
