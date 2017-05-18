def install(job):
    service = job.service
    service.aysrepo.
    actor = service.aysrepo.actorGet("container")
    args = {
        'node': service.models.data.node,
        'flist': 'https://hub.gig.tech/gig-official-apps/gw.flist',
        'hostNetworking': True,
    }
    cont_service = actor.serviceCreate(instance=service.name., args=args)
    j.tools.async.wrappers.sync(cont_service.executeAction('install'))
    pass