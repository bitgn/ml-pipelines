from env import *
from test_api import events_pb2 as evt

def given_a_dataset_with_location_id(t: Env):
    """show location id on all related screens"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)

    ds.meta.location_id = "aws"
    ds.meta.location_id_set = True

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

    ds.meta.location_id_set = False

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