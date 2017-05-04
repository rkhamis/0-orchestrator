from enum import Enum


class EnumVMStatus(Enum):
    running = "running"
    halted = "halted"
    paused = "paused"
    halting = "halting"
    migrating = "migrating"
