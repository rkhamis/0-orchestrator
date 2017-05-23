from JumpScale import j


def input(job):
    for arg in ['filesystems', 'arbds']:
        if job.model.args.get(arg, []) != []:
            raise j.exceptions.Input('{} should not be set as input'.format(arg))

    domain = job.model.args.get('domain')
    if not domain:
        raise j.exceptions.Input('Domain cannot be empty.')

    nics = job.model.args.get('nics', [])
    publicnetwork = False
    privatenetwork = False
    for nic in nics:
        config = nic.get('config')
        name = nic.get('name')
        dhcp = nic.get('dhcpserver')
        cidr = None

        if not name:
            raise j.exceptions.Input('Gateway nic should have name defined.')

        if config:
            cidr = config.get('cidr')
            if config.get('gateway'):
                publicnetwork = True
            else:
                privatenetwork = True

        if (dhcp and not config) or (dhcp and config and not cidr):
            raise j.exceptions.Input('Gateway should have cidr if dhcp is defined.')

        if dhcp:
            nameservers = dhcp.get('nameservers')

            if not nameservers:
                raise j.exceptions.Input('Dhcp nameservers should have at least one nameserver.')

    if not publicnetwork:
        raise j.exceptions.Input("Gateway should have at least one Public Address (gw defined)")
    if not privatenetwork:
        raise j.exceptions.Input("Gateway should have at least one Private Address (no gw defined)")
    return job.model.args


def init(job):
    service = job.service
    containeractor = service.aysrepo.actorGet("container")
    nics = service.model.data.to_dict()['nics']  # get dict version of nics
    for nic in nics:
        if 'dhcpserver' in nic:
            nic.pop('dhcpserver')

    args = {
        'node': service.model.data.node,
        'flist': 'https://hub.gig.tech/gig-official-apps/g8osgw.flist',
        'nics': nics,
        'hostname': service.model.data.hostname,
        'hostNetworking': False,
    }
    cont_service = containeractor.serviceCreate(instance=service.name, args=args)
    service.consume(cont_service)

    args = {
        'container': service.name
    }

    # create firewall
    fwactor = service.aysrepo.actorGet('firewall')
    fwactor.serviceCreate(instance=service.name, args=args)

    # create http
    httpactor = service.aysrepo.actorGet('http')
    httpactor.serviceCreate(instance=service.name, args=args)

    # create dhcp
    dhcpactor = service.aysrepo.actorGet('dhcp')
    dhcpactor.serviceCreate(instance=service.name, args=args)

    # Start cloudinit
    cloudinitactor = service.aysrepo.actorGet("cloudinit")
    cloudinitactor.serviceCreate(instance=service.name, args=args)


def install(job):
    # nothing to do here all our children will be created by ays automagic
    pass


def processChange(job):
    service = job.service
    args = job.model.args
    category = args.pop('changeCategory')

    if category != 'dataschema':
        return

    if service.model.data.nics != args.get('nics', service.model.data.nics):
        cloudInitServ = service.aysrepo.serviceGet(role='cloudinit', instance=service.name)
        j.tools.async.wrappers.sync(cloudInitServ.executeAction('update', args={'nics': args['nics']}))

        firewallServ = service.aysrepo.serviceGet(role='firewall', instance=service.name)
        j.tools.async.wrappers.sync(firewallServ.executeAction('update', args={'data': args}))

        dhcpServ = service.aysrepo.serviceGet(role='dhcp', instance=service.name)
        j.tools.async.wrappers.sync(dhcpServ.executeAction('update', args={'data': args}))

        service.model.data.nics = args['nics']

    if args.get("httpproxies", None):
        httpServ = service.aysrepo.serviceGet(role='http', instance=service.name)
        j.tools.async.wrappers.sync(httpServ.executeAction('update', args={'httpproxies': args["httpproxies"]}))
