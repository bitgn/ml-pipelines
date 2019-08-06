from env import *


def given_a_dataset_with_location_id(t: Env):
    """show location id on all related screens"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)

    ds.metadata.location_id = "aws"
    ds.metadata.set_fields.append(evt.FIELD_LOCATION_ID)

    t.given_events(prj, ds)

    t.scenario(
        when.view_dataset(ds.dataset_id),
        then.text(f'main #ds-{ds.dataset_id} .location-id', 'aws'),
    )

    t.scenario(
        when.list_datasets(),
        then.text(f'main #ds-{ds.dataset_id} .location-id', 'aws')
    )

    t.scenario(
        when.view_project(prj.project_id),
        then.text(f'main #ds-{ds.dataset_id} .location-id', 'aws')
    )


def given_a_dataset_without_location_id(t: Env):
    """hide location id field completely"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)

    ds.metadata.del_fields.append(evt.FIELD_LOCATION_ID)

    t.given_events(prj, ds)

    t.scenario(
        when.view_dataset(ds.dataset_id),
        then.none(f'main #ds-{ds.dataset_id} .location-id'),
    )

    t.scenario(
        when.list_datasets(),
        then.none(f'main #ds-{ds.dataset_id} .location-id')
    )

    t.scenario(
        when.view_project(prj.project_id),
        then.none(f'main #ds-{ds.dataset_id} .location-id')
    )