from JumpScale import j


def input(job):
    ays_repo = job.service.aysrepo
    services = ays_repo.servicesFind(actor=job.service.model.dbobj.actorName)
    name = services[0].name if len(services) == 1 else ''

    if not services or name == job.service.name:
        return
    else:
        raise j.exceptions.RuntimeError('Repo can\'t contain multiple configuration services')


def processChange(job):
    service = job.service
    args = job.model.args
    category = args.pop('changeCategory')
    if category == 'dataschema':
        service.model.data.configurations = args.get('configurations', service.model.data.configurations)
        service.saveAll()
