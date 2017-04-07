from JumpScale import j


def input(job):
    for arg in ['filesystems', 'arbds']:
        if job.model.args.get(arg, []) != []:
            raise j.exceptions.Input("{} should not be set as input".format(arg))


def get_node(node_service):
    return j.sal.g8os.get_node(
        addr=node_service.model.data.redisAddr,
        port=node_service.model.data.redisPort,
        password=node_service.model.data.redisPassword or None,
    )

def get_cluster(job):
    service = job.service

    nodes = []
    for node_service in service.producers['node']:
        nodes.append(get_node(node_service))

    return j.sal.g8os.create_storagecluster(
        label=service.model.data.label,
        nodes=nodes,
        disk_type=str(service.model.data.diskType),
        nr_server=service.model.data.nrServer,
        has_slave=service.model.data.hasSlave,
    )


def init(job):
    job.service.model.data.status = 'empty'

def install(job):
    service = job.service

    job.service.model.data.status = 'deploying'
    cluster = get_cluster(job)
    cluster.ays.create(service.aysrepo)

def start(job):
    service = job.service

    cluster = get_cluster(job)
    cluster.ays.create(service.aysrepo)

    for storage_server in cluster.storage_servers:
        storage_server.start()
        ardb_service = storage_server.ardb.ays.get(service.aysrepo)
        ardb_service.model.data.bind = storage_server.ardb.bind

    job.service.model.data.status = 'ready'


def delete(job):
    # since we consume all the ardb, this will be called once everything is deleted
    job.service.model.data.status = 'empty'


def addStorageServer(job):
    raise NotImplementedError()


def reoveStorageServer(job):
    raise NotImplementedError()
