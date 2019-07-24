import typing

import lmdb

from fdb import tuple as tpl
from proto import dto_pb2 as dto
from google.protobuf import message

import utils

Tx = lmdb.Transaction


def connect(path: str):
    env = lmdb.open(path)


EVENTS = 0x00
PROJECTS = 0x01
DATASETS = 0x02
PROJECT_DATASETS = 0x03
LINEAGE = 0x04
STATS=0x05

LINEAGE_CACHE=0x06



def events_append(tx: lmdb.Transaction, val: message.Message):
    id = utils.sequential_uuid()
    tx.put(tpl.pack((EVENTS, id)), val.SerializeToString())


def project_add(tx: Tx, val: dto.ProjectData):
    tx.put(tpl.pack((PROJECTS, val.id)), val.SerializeToString())


def project_get(tx: lmdb.Transaction, project_id: str):
    bytes = tx.get(tpl.pack((PROJECTS, project_id)))
    if not bytes:
        return None
    val = dto.ProjectData()
    val.ParseFromString(bytes)
    return val


def project_list(tx: lmdb.Transaction) -> typing.Iterable[dto.ProjectData]:
    r = tpl.range((PROJECTS,))
    with tx.cursor() as curs:
        if not curs.set_range(r.start):
            return []

        for key, bytes in curs:
            if key > r.stop:
                break
            val = dto.ProjectData()
            val.ParseFromString(bytes)
            yield val


def dataset_add(tx: Tx, val: dto.DatasetData):
    key = tpl.pack((DATASETS, val.dataset_id))
    index_key = tpl.pack((PROJECT_DATASETS, val.project_id, val.dataset_id))
    tx.put(key, val.SerializeToString())
    tx.put(index_key, b'')

def dataset_get(tx: Tx, dataset_id: str) -> typing.Optional[dto.DatasetData]:
    key = tpl.pack((DATASETS, dataset_id))
    bytes = tx.get(key)
    if not bytes:
        return None
    val = dto.DatasetData()
    val.ParseFromString(bytes)
    return val

def dataset_list(tx: lmdb.Transaction, project_id: str):
    r = tpl.range((PROJECT_DATASETS, project_id))
    with tx.cursor() as curs:
        if not curs.set_range(r.start):
            return []
        for key, bs in curs:
            if key > r.stop:
                break

            index_key = tpl.unpack(key)


            bs = tx.get(tpl.pack((DATASETS, index_key[2])))
            val = dto.DatasetData()
            val.ParseFromString(bs)
            yield val


def dataset_list_all(tx: lmdb.Transaction):
    r = tpl.range((DATASETS, ))
    with tx.cursor() as curs:
        if not curs.set_range(r.start):
            return []
        for key, bytes in curs:
            if key > r.stop:
                break
            val = dto.DatasetData()
            val.ParseFromString(bytes)
            yield val



def lineage_set(tx: lmdb.Transaction, val : dto.DatasetLink):
    key = tpl.pack((LINEAGE, val.link_id))
    tx.put(key, val.SerializeToString())

def lineage_get(tx: lmdb.Transaction, lineage_id: str) -> typing.Optional[dto.DatasetLink]:
    r = tpl.pack((LINEAGE, lineage_id))
    bs = tx.get(r)
    if bs:
        val = dto.DatasetLink()
        val.ParseFromString(bs)
        return val
    return None

def stats_get(tx: lmdb.Transaction) -> dto.TenantStats:
    r = tpl.pack((STATS,))
    val = dto.TenantStats()
    data = tx.get(r)
    if data:
        val.ParseFromString(data)
    return val

def stats_put(tx: lmdb.Transaction, stats: dto.TenantStats):
    tx.put(tpl.pack((STATS,)),stats.SerializeToString())


def cache_get(tx: lmdb.Transaction, id: str) -> typing.Optional[dto.LineageCache]:
    r = tpl.pack((LINEAGE_CACHE,id))

    data = tx.get(r)
    if not data:
        return None
    val = dto.LineageCache()
    val.ParseFromString(data)
    return val




def cache_put(tx: lmdb.Transaction, id:str, cache: dto.LineageCache):
    tx.put(tpl.pack((LINEAGE_CACHE,id)),cache.SerializeToString())



def counter_get(tx: lmdb.Transaction, key):
    key = tpl.pack(key)
    val = tx.get(key)
    if not val:
        return 0
    return tpl.unpack(val)[0]


def counter_set(tx: lmdb.Transaction, key, val: int):
    key = tpl.pack(key)
    tx.put(key, tpl.pack(val))


def counter_inc(tx: Tx, key):
    val = counter_get(tx, key)
    counter_set(tx, key, val + 1)


class DB:
    env: lmdb.Environment

    def __init__(self, env):
        self.env = env
