def install(job):
    from JumpScale.sal.g8os.Container import Container
    from JumpScale.sal.g8os.gateway.http import HTTPServer
    container = Container.from_ays(job.service.parent)
    gateway = job.service.parent.consumers['gateway'][0]
    gwdata = gateway.model.data.to_dict()
    httpproxies = gwdata.get('httpproxies', [])
    http = HTTPServer(container, httpproxies)
    http.apply_rules()


