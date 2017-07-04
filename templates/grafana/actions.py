def apply_config(job):
    service = job.service
    from zeroos.orchestrator.sal.Container import Container

    containers = service.producers.get('container')
    container = Container.from_ays(containers[0])
    try:
        f = container.filesystem.open('/opt/grafana/conf/defaults.conf')
        template = container.filesystem.read(f)
    finally:
        container.filesystem.close(f)

    template = template.replace('3000', job.service.model.data.port)
    try:
        f = container.filesystem.open('/etc/grafana/grafana.ini', 'w')
        template = container.filesystem.write(f, template)
    finally:
        container.filesystem.close(f)
    return


def configure_datasources(job):
    import time
    import requests
    service = job.service
    grafanaclient = j.clients.grafana.get(url='http://%s:%d' % (job.service.parent.model.data.redisAddr, service.model.data.port), username='admin', password='admin')
    influxes = service.producers.get('influxdb')
    for influx in influxes:
        for i, database in enumerate(influx.model.data.databases):
            data = {
                'type': 'influxdb',
                'access': 'proxy',
                'database': database,
                'name': influx.name,
                'url': 'http://%s:%u' % (influx.parent.model.data.redisAddr, influx.model.data.port),
                'user': 'admin',
                'password': 'passwd',
                'default': True,
            }

            now = time.time()
            while time.time() - now < 10:
                try:
                    grafanaclient.addDataSource(data)
                    if len(grafanaclient.listDataSources()) == i + 1:
                        continue
                    break
                except requests.exceptions.ConnectionError:
                    time.sleep(1)
                    pass


def init(job):
    from zeroos.orchestrator.configuration import get_configuration

    service = job.service
    container_actor = service.aysrepo.actorGet('container')
    config = get_configuration(service.aysrepo)

    args = {
        'node': service.model.data.node,
        'flist': config.get(
            '0-grafana-flist', 'https://hub.gig.tech/gig-official-apps/grafana.flist'),
        'hostNetworking': True,
        'initProcesses': [{'name': 'grafana-server', 'args': ['-config', '/etc/grafana/grafana.ini']}],
    }
    cont_service = container_actor.serviceCreate(instance=service.name, args=args)
    service.consume(cont_service)


def install(job):
    service = job.service
    containers = service.producers.get('container')
    if containers:
        apply_config(job)
        j.tools.async.wrappers.sync(containers[0].executeAction('start', context=job.context))
    service.model.data.status = 'running'
    configure_datasources(job)
    service.saveAll()


def uninstall(job):
    service = job.service
    containers = service.producers.get('container')
    if containers:
        j.tools.async.wrappers.sync(containers[0].executeAction('stop', context=job.context))
        j.tools.async.wrappers.sync(containers[0].delete())
    service.delete()


def processChange(job):
    service = job.service
    args = job.model.args

    if args.pop('changeCategory') != 'dataschema' or service.model.actionsState['install'] in ['new', 'scheduled']:
        return

    if args.get('port'):
        service.model.data.port = args['port']

    containers = service.producers.get('container')
    if containers:
        container = containers[0]
        j.tools.async.wrappers.sync(container.executeAction('stop', context=job.context))
        service.model.data.status = 'halted'
        apply_config(job)
        j.tools.async.wrappers.sync(container.executeAction('start', context=job.context))
        service.model.data.status = 'running'

    service.saveAll()


def init_actions_(service, args):
    return {
        'init': [],
        'install': ['init'],
        'monitor': ['start'],
        'delete': ['uninstall'],
        'uninstall': [],
    }
