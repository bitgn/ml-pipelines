from env import *
from evil import pretty
from api import events_pb2 as evt

def given_a_dataset_with_storage_size(t: env.Env):
    """show storage id on all related screens"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)

    size = pretty.bytes(ds.meta.storage_bytes)

    t.given_events(prj, ds)


    t.scenario(
        when.view_dataset(ds.dataset_id),
        then.text(f'main #ds-{ds.dataset_id} .zip-size', size),
    )

    t.scenario(
        when.list_datasets(),
        then.text(f'main #ds-{ds.dataset_id} .zip-size', size),
    )

    t.scenario(
        when.view_project(prj.project_id),
        then.text(f'main #ds-{ds.dataset_id} .zip-size', size),
    )


def given_a_dataset_without_storage_size(t: Env):
    """show 0 bytes"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)

    ds.meta.storage_bytes_state = evt.DELETE

    t.given_events(prj, ds)

    size = pretty.bytes(0)

    t.scenario(
        when.view_dataset(ds.dataset_id),
        then.text(f'main #ds-{ds.dataset_id} .zip-size', size),
    )

    t.scenario(
        when.list_datasets(),
        then.text(f'main #ds-{ds.dataset_id} .zip-size', size),
    )

    t.scenario(
        when.view_project(prj.project_id),
        then.text(f'main #ds-{ds.dataset_id} .zip-size', size),
    )