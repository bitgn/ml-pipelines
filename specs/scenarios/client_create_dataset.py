import struct

from env import *


def given_empty_system(e: Env):
    """throw an error"""

    e.scenario(
        when.client(lambda c: c.create_dataset(name="test",project_name='non-existent')),
        then.not_found(subject_name='non-existent')
    )

def uid(i):
    return struct.pack(">llll",0,0, 0,  i)


def given_a_project(e: Env):
    "Create a new dataset"
    prj = preset.project_created(e)

    e.given_events(prj)

    e.scenario(
        when.client(lambda c: c.create_dataset(name="test", project_name=prj.name)),
        then.client_ok(uid=uid(0))
    )

def given_a_dataset(e: Env):
    "Create a new dataset"
    prj = preset.project_created(e)
    ds = preset.dataset_created(e, prj)

    e.given_events(prj, ds)

    e.scenario(
        when.client(lambda c: c.create_dataset(name=ds.name, project_name=prj.name)),
        then.name_taken(subject_uid=ds.uid, subject_name=ds.name)
    )


