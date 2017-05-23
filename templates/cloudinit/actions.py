def install(job):
    gateway = job.service.parent.consumers['gateway'][0]
    gwdata = gateway.model.data.to_dict()

    cloudinit = config_cloud_init(job, gwdata.get("nics", []))
    if not cloudinit.is_running():
        cloudinit.start()


def config_cloud_init(job, nics=None):
    import json
    from JumpScale.sal.g8os.gateway.cloudinit import CloudInit
    from JumpScale.sal.g8os.Container import Container

    container = Container.from_ays(job.service.parent)
    nics = [] if nics is None else nics
    config = {}
    for nic in nics:
        for host in nic["dhcpserver"]["hosts"]:
            userdata = json.loads(host["cloudinit"]["userdata"])
            metadata = json.loads(host["cloudinit"]["metadata"])
            config[host.macaddress] = json.dumps({
                "meta-data": metadata,
                "user-data": userdata,
            })

    cloudinit = CloudInit(container, config)
    cloudinit.apply_config()
    return cloudinit


def update(job):
    if job.model.args.get("nics", None):
        config_cloud_init(job, job.model.args["nics"])
