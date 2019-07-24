from proto import dto_pb2 as dto
from explore import config
import datetime as dt

def is_stale(dm: dto.DatasetData):
    if not dm.update_timestamp_set:
        return False

    was = dt.datetime.utcfromtimestamp(dm.update_timestamp)

    now = config.utcnow()

    stale = (now - was).days >= 3
    return stale