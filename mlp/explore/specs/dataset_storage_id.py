
from django import urls

from explore.specs import when, preset
from test import *
from utils import pretty

from proto import events_pb2 as evt

def given_a_dataset_with_storage_id(t: test.Env):
    """show storage id on all related screens"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)

    ds.metadata.storage_id = "aws"
    ds.metadata.set_fields.append(evt.FIELD_STORAGE_ID)

    t.given_events(prj, ds)

    t.scenario(
        when.view_dataset(ds.dataset_id),
        then.text(f'main #ds-{ds.dataset_id} .storage-id', 'aws'),
    )

    t.scenario(
        when.list_datasets(),
        then.text(f'main #ds-{ds.dataset_id} .storage-id', 'aws')
    )

    t.scenario(
        when.view_project(prj.project_id),
        then.text(f'main #ds-{ds.dataset_id} .storage-id', 'aws')
    )


def given_a_dataset_without_storage_id(t: test.Env):
    """hide storage id field completely"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)

    ds.metadata.del_fields.append(evt.FIELD_STORAGE_ID)

    t.given_events(prj, ds)

    t.scenario(
        when.view_dataset(ds.dataset_id),
        then.none(f'main #ds-{ds.dataset_id} .storage-id'),
    )

    t.scenario(
        when.list_datasets(),
        then.none(f'main #ds-{ds.dataset_id} .storage-id')
    )

    t.scenario(
        when.view_project(prj.project_id),
        then.none(f'main #ds-{ds.dataset_id} .storage-id')
    )