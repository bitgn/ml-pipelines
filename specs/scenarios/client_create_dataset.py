import struct

from env import *


def given_empty_system(e: Env):
    """throw an error"""
    e.scenario(
        when.client(lambda c: c['test'].create_dataset(name="test")),
        then.not_found(subject_name='non-existent')
    )



def given_a_project(e: Env):
    "Create a new dataset"
    prj = preset.project_created(e)

    e.given_events(prj)

    e.scenario(
        when.client(lambda c: c[prj.name].create_dataset(name="test")),
        then.client_ok(uid=uid(1))
    )

def given_a_dataset(e: Env):
    "Create a new dataset"
    prj = preset.project_created(e)
    ds = preset.dataset_created(e, prj)

    e.given_events(prj, ds)

    e.scenario(
        when.client(lambda c: c[prj.name].create_dataset(name=ds.name)),
        then.already_exists(subject_uid=ds.uid, subject_name=ds.name)
    )


