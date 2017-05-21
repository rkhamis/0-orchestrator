def init(job):
    service = job.service
    actor = service.aysrepo.actorGet("container")
    nics = service.model.data.to_dict()['nics'] # get dict version of nics
    args = {
        'node': service.model.data.node,
        'flist': 'https://hub.gig.tech/gig-official-apps/gw.flist',
        'nics': nics,
        'hostname': service.model.data.hostname,
        'hostNetworking': False,
    }
    cont_service = actor.serviceCreate(instance=service.name, args=args)
    service.consume(cont_service)


def install(job):
    # nothing to do here all our children will be created by ays automagic