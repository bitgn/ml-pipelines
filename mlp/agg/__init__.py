"""
Contains domain logic (in form of event-driven aggregate) for the project
"""


from agg import service
import lmdb


def open(path:str):
    return service.Service(lmdb.open(path=path))
