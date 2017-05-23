def install(job):

    gateway = job.service.parent.consumers['gateway'][0]
    gwdata = gateway.model.data.to_dict()
    httpproxies = gwdata.get('httpproxies', [])
    apply_rules(job, httpproxies)


def apply_rules(job, httpproxies=None):
    from JumpScale.sal.g8os.Container import Container
    from JumpScale.sal.g8os.gateway.http import HTTPServer

    httpproxies = [] if httpproxies is None else httpproxies
    container = Container.from_ays(job.service.parent)
    http = HTTPServer(container, httpproxies)
    http.apply_rules()


def update(job):
    apply_rules(job, job.model.args["httpproxies"])
