
def install(job):
    raise NotImplementedError()

def start(job):
    raise NotImplementedError()

def stop(job):
    raise NotImplementedError()

def monitor(job):
    raise NotImplementedError()
    # check if the container is still running, try to restart if it halted
