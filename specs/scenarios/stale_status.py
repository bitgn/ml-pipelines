"""MLP-3"""
from datetime import timedelta

from env import *



def given_a_dataset_never_updated(t: Env):
    """hide stale status"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)

    t.given_events(prj, ds)

    t.scenario(
        when.view_project(prj.name),
        then.none(f'main #ds-{ds.uid.hex()} .stale-status')
    )

    t.scenario(
        when.view_dataset(ds.project_name, ds.name),
        then.none('main .stale-status')
    )

    t.scenario(
        when.list_datasets(),
        then.none(f'main #ds-{ds.uid.hex()} .stale-status')
    )

    t.scenario(
        when.search_datasets("stale"),
        then.none(f'main #ds-{ds.uid.hex()}')
    )


def given_a_dataset_updated_today(t: Env):
    """hide stale status"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)

    ver = preset.dataset_version_added(t, ds)

    t.given_events(prj, ds, ver)

    t.scenario(
        when.view_project(prj.name),
        then.none(f'main #ds-{ds.uid.hex()} .stale-status')
    )

    t.scenario(
        when.view_dataset(ds.project_name, ds.name),
        then.none('main .stale-status')
    )

    t.scenario(
        when.list_datasets(),
        then.none(f'main #ds-{ds.uid.hex()} .stale-status')
    )

    t.scenario(
        when.search_datasets("stale"),
        then.none(f'main #ds-{ds.uid.hex()}')
    )


def given_a_dataset_updated_week_ago(t: Env):
    """show stale status"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)
    ver = preset.dataset_version_added(t, ds, ts=timedelta(days=-7))

    t.given_events(prj, ds, ver)

    t.scenario(
        when.view_project(prj.name),
        then.exists(f'main #ds-{ds.uid.hex()} .stale-status')
    )

    t.scenario(
        when.view_dataset(prj.name, ds.name),
        then.exists('main .stale-status')
    )

    t.scenario(
        when.list_datasets(),
        then.exists(f'main #ds-{ds.uid.hex()} .stale-status')
    )

    t.scenario(
        when.search_datasets("stale"),
        then.exists(f'main #ds-{ds.uid.hex()} .stale-status')
    )
