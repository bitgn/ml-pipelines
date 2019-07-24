"""MLP-3"""

from explore.specs import when, preset
from test import *


def given_a_dataset_updated_today(t: test.Env):
    """hide stale status"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)

    preset.set_update_timestamp(t, ds.metadata, 0)

    t.given_events(prj, ds)

    t.scenario(
        when.view_project(prj.project_id),
        then.none(f'main #ds-{ds.dataset_id} .stale-status')
    )

    t.scenario(
        when.view_dataset(ds.dataset_id),
        then.none('main .stale-status')
    )

    t.scenario(
        when.list_datasets(),
        then.none(f'main #ds-{ds.dataset_id} .stale-status')
    )


def given_a_dataset_updated_week_ago(t: test.Env):
    """show stale status"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)
    preset.set_update_timestamp(t, ds.metadata, days=-7)

    t.given_events(prj, ds)

    t.scenario(
        when.view_project(prj.project_id),
        then.exists(f'main #ds-{ds.dataset_id} .stale-status')
    )

    t.scenario(
        when.view_dataset(ds.dataset_id),
        then.exists('main .stale-status')
    )

    t.scenario(
        when.list_datasets(),
        then.exists(f'main #ds-{ds.dataset_id} .stale-status')
    )
