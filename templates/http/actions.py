def install(job):
    import ipaddress
    from JumpScale.sal.g8os.Container import Container
    from JumpScale.sal.g8os.gateway.http import HTTPServer
    container = Container.from_ays(job.service.parent)
    gateway = job.service.parent.consumers['gateway'][0]
    gwdata = gateway.model.data.to_dict()
    httpproxies = gwdata.get('httpproxies', [])
    nics = gwdata.get('nics', [])
    # we add some proxies specially for cloud-init
    for nic in nics:
        if 'dhcpserver' in nic and 'config' in nic and 'cidr' in nic['config']:
            gwip = str(ipaddress.IPv4Interface(nic['config']['cidr']).ip)
            httpproxies.append({
                'host': gwip,
                'destinations': ['http://127.0.0.1:8080'],
                'types': ['http']}
            )

    http = HTTPServer(container, httpproxies)
    http.apply_rules()


