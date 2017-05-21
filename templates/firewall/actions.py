def install(job):
    import ipaddress
    from JumpScale.sal.g8os.Container import Container
    from JumpScale.sal.g8os.gateway.firewall import Firewall
    container = Container.from_ays(job.service.parent)
    gateway = job.service.parent.consumers['gateway'][0]
    gwdata = gateway.model.data.to_dict()
    portforwards = gwdata.get('portforwards', [])
    # lets assume the public ip is the ip of the nic which has a gateway configured
    publicip = None
    privatenetwork = None
    for nic in gateway.model.data.nics:
        if nic.config:
            if nic.config.gateway:
                publicip = str(ipaddress.IPv4Interface(nic.config.cidr).ip)
            else:
                privatenetwork = str(ipaddress.IPv4Interface(nic.config.cidr).network)
    firewall = Firewall(container, publicip, privatenetwork, portforwards)
    firewall.apply_rules()




