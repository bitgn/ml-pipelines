import struct

from env import *
from client import ml_pipelines as client




def extend_last_version(project, dataset: str):
    def _(c: client.Client):
        ver = c.get_project(project).get_dataset(dataset).get_last_version()
        staging = ver.prepare_commit()
        staging.add_file(name='some', records=1, size=2)
        return staging.commit("New dataset")

    return _


def given_no_project(e: Env):
    """throw an error"""
    e.scenario(
        when.client(extend_last_version('none', 'none')),
        then.not_found(subject_name='project')
    )


def given_no_dataset(e: Env):
    """throw an error"""
    prj = preset.project_created(e)
    e.given_events(prj)

    e.scenario(
        when.client(extend_last_version(prj.name, 'dataset')),
        then.not_found(subject_name='dataset')
    )


def given_empty_dataset(e: Env):
    """create a new version"""

    prj = preset.project_created(e)
    ds = preset.dataset_created(e, prj)

    e.given_events(prj, ds)

    e.scenario(
        when.client(extend_last_version(prj.name, ds.name)),
        then.client_ok(uid=uid(1))
    )

def given_existing_version(e: Env):
    prj = preset.project_created(e)
    ds = preset.dataset_created(e, prj)
    ver = preset.dataset_version_added(e, ds)

    e.given_events(prj, ds, ver)

    e.scenario(
        when.client(extend_last_version(prj.name, ds.name)),
        then.client_ok(uid=uid(1))
    )
