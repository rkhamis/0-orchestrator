
def install(job):
    raise NotImplementedError()
    # write configuration file

def start(job):
    raise NotImplementedError()
    # start ardb process

def stop(job):
    raise NotImplementedError()
    # stop ardb process

def monitor(job):
    raise NotImplementedError()
    # check if the process is still running, try to restart if it halted
