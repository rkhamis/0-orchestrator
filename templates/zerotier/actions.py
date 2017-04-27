def _get_client(parent):
    return j.clients.g8core.get(host=parent.model.data.redisAddr,
                                port=parent.model.data.redisPort or 6379,
                                password=parent.model.data.redisPassword or '')


def install(job):
    service = job.service
    client = _get_client(service.parent)
    client.zerotier.join(service.model.data.nwid)

    for network in client.zerotier.list():
        if network['id'] == service.model.data.nwid:
            service.model.data.allowDefault = network['allowDefault']
            service.model.data.allowGlobal = network['allowGlobal']
            service.model.data.allowManaged = network['allowManaged']
            service.model.data.allowDefault = network['allowDefault']
            service.model.data.assignedAddresses = network['assignedAddresses']
            service.model.data.bridge = network['bridge']
            service.model.data.broadcastEnabled = network['broadcastEnabled']
            service.model.data.dhcp = network['dhcp']
            service.model.data.mac = network['mac']
            service.model.data.mtu = network['mtu']
            service.model.data.name = network['name']
            service.model.data.netconfRevision = network['netconfRevision']
            service.model.data.portDeviceName = network['portDeviceName']
            service.model.data.portError = network['portError']
            service.model.data.routes = network['routes']
            service.model.data.status = network['status']
            service.model.data.type = network['type'].lower()
            break

def delete(job):
    service = job.service
    client = _get_client(service.parent)
    client.zerotier.leave(service.model.data.networkID)
