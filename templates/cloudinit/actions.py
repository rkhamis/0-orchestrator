def install(job):
    import json
    from JumpScale.sal.g8os.Container import Container
    from JumpScale.sal.g8os.gateway.cloudinit import CloudInit

    container = Container.from_ays(job.service.parent)
    gateway = job.service.parent.consumers['gateway'][0]

    config = {}
    for dhcp in gateway.model.data.dhcps:
        for host in dhcp.hosts:
            userdata = json.loads(host.cloudinit.userdata)
            metadata = json.loads(host.cloudinit.metadata)
            config[host.macaddress] = json.dumps({
                "meta-data": metadata,
                "user-data": userdata,
            })

    cloudinit = CloudInit(container, config)
    cloudinit.apply_config()
    cloudinit.start()
