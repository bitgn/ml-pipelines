from env import *
from evil import pretty

def given_a_dataset_with_storage_size(t: env.Env):
    """show storage id on all related screens"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)

    t.given_events(prj, ds)

    size = pretty.bytes(ds.metadata.zip_bytes)

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

    ds.metadata.del_fields.append(evt.FIELD_ZIP_BYTES)

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