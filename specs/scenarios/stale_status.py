"""MLP-3"""
from env import *


def given_a_dataset_updated_today(t: Env):
    """hide stale status"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)

    preset.set_update_timestamp(t, ds.meta, 0)

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

    t.scenario(
        when.search_datasets("stale"),
        then.none(f'main #ds-{ds.dataset_id}')
    )


def given_a_dataset_updated_week_ago(t: Env):
    """show stale status"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)
    preset.set_update_timestamp(t, ds.meta, days=-7)

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

    t.scenario(
        when.search_datasets("stale"),
        then.exists(f'main #ds-{ds.dataset_id} .stale-status')
    )
