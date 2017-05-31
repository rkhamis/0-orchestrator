class Config:
    def __init__(self, ays_repo):
        services = ays_repo.servicesFind(actor='configuration')
        if len(services) > 1:
            raise RuntimeError('Multiple configuration services found')

        configurations = services[0].model.data.to_dict()['configurations'] if services else []
        self.configurations = {conf['key']: conf['value'] for conf in configurations}
