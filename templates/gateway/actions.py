def input(job):
    for arg in ['filesystems', 'arbds']:
        if job.model.args.get(arg, []) != []:
            raise j.exceptions.Input("{} should not be set as input".format(arg))

    nics = job.model.args.get('nics', [])
    publicnetwork = False
    privatenetwork = False
    for nic in nics:
        config = nic.get('config')
        if config:
            if config.get('gateway'):
                publicnetwork = True
            else:
                privatenetwork = True
    if not publicnetwork:
        raise j.exceptions.Input("Gateway should have atleast one Public Address (gw defined)")
    if not privatenetwork:
        raise j.exceptions.Input("Gateway should have atleast one Private Address (no gw defined)")
    return job.model.args


def init(job):
    service = job.service
    containeractor = service.aysrepo.actorGet("container")
    nics = service.model.data.to_dict()['nics']  # get dict version of nics
    args = {
        'node': service.model.data.node,
        'flist': 'https://hub.gig.tech/gig-official-apps/g8osgw.flist',
        'nics': nics,
        'hostname': service.model.data.hostname,
        'hostNetworking': False,
    }
    cont_service = containeractor.serviceCreate(instance=service.name, args=args)
    service.consume(cont_service)

    # create firewall
    fwactor = service.aysrepo.actorGet("firewall")
    args = {
        'container': service.name
    }
    fwactor.serviceCreate(instance=service.name, args=args)


def install(job):
    # nothing to do here all our children will be created by ays automagic