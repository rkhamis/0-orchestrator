def install(job):
    import ipaddress
    from JumpScale.sal.g8os.Container import Container
    from JumpScale.sal.g8os.gateway.firewall import Firewall, Network
    container = Container.from_ays(job.service.parent)
    gateway = job.service.parent.consumers['gateway'][0]
    gwdata = gateway.model.data.to_dict()
    portforwards = gwdata.get('portforwards', [])
    # lets assume the public ip is the ip of the nic which has a gateway configured
    publicnetwork = None
    privatenetworks = []
    for nic in gateway.model.data.nics:
        if nic.config:
            if nic.config.gateway:
                publicnetwork = Network(nic.name, nic.config.cidr)
            else:
                privatenetworks.append(Network(nic.name, nic.config.cidr))
    firewall = Firewall(container, publicnetwork, privatenetworks, portforwards)
    firewall.apply_rules()




