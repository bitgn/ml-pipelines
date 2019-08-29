import struct

from env import *


def given_no_project(e: Env):
    """throw an error"""

    e.scenario(
        when.client(lambda c: c.add_dataset_version(project_name="project", dataset_name="dataset", items=())),
        then.not_found(subject_name='project')
    )


def given_no_dataset(e: Env):
    """throw an error"""

    prj = preset.project_created(e)
    e.given_events(prj)

    e.scenario(
        when.client(lambda c: c.add_dataset_version(project_name=prj.name, dataset_name="dataset", items=())),
        then.not_found(subject_name='dataset')
    )

def given_empty_dataset(e: Env):
    """create a new version"""

    prj = preset.project_created(e)
    ds = preset.dataset_created(e, prj)

    e.given_events(prj, ds)

    items = [
        cl.DatasetItem('some', 1, 2)

    ]

    e.scenario(
        when.client(lambda c: c.add_dataset_version(project_name=ds.project_name, dataset_name=ds.name, items=items)),
        then.client_ok(uid=uid(1))
    )

    e.scenario(
        when.client(lambda c: c.add_dataset_version(project_name=ds.project_name, dataset_name=ds.name, items=items, parent_uid=uid(7))),
        then.not_found(uid(7))
    )

def given_existing_version(e: Env):
    prj = preset.project_created(e)
    ds = preset.dataset_created(e, prj)
    ver = preset.dataset_version_added(e, ds)

    e.given_events(prj, ds, ver)

    items = [
        cl.DatasetItem('some', 1, 2)

    ]

    e.scenario(
        when.client(lambda c: c.add_dataset_version(project_name=ds.project_name, dataset_name=ds.name, items=items, parent_uid=ver.uid)),
        then.client_ok(uid=uid(1))
    )
