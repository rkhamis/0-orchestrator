from JumpScale import j


def install(job):
    import ipaddress
    from JumpScale.sal.g8os.Container import Container
    from JumpScale.sal.g8os.gateway.dhcp import DHCP

    container = Container.from_ays(job.service.parent)
    gateway = job.service.parent.consumers['gateway'][0]
    gwdata = gateway.model.data.to_dict()
    nics = gwdata.get('nics', [])
    dhcpservers = []

    for nic in nics:
        dhcpserver = nic.get('dhcpserver')
        if not dhcpserver:
            continue

        cidr = ipaddress.IPv4Interface(nic['config']['cidr'])
        dhcpserver['subnet'] = str(cidr.network)
        dhcpserver['gateway'] = str(cidr.ip)
        dhcpserver['interface'] = nic['name']
        dhcpservers.append(dhcpserver)

    dhcp = DHCP(container, gwdata['domain'], dhcpservers)
    dhcp.apply_config()








