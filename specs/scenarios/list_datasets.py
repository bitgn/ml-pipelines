"""MLP-2"""
from env import *
import datetime as dt
from evil import pretty


def given_no_datasets(e: Env):
    """show an actionable message"""

    e.scenario(
        when.list_datasets(),
        then.exists('main #empty-catalog-message'),
        then.none('main .dataset-info'),
    )


def given_a_dataset(t: Env):
    """show dataset properties"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)

    t.given_events(prj, ds)
    past = dt.datetime.utcfromtimestamp(ds.metadata.update_timestamp)
    now = t.time

    since = pretty.timedelta(now - past)

    href = urls.view_project(ds.project_id)
    t.scenario(
        when.list_datasets(),
        then.none('main #empty-catalog-message'),
        then.count('main .dataset-info', 1),
        then.text(f'main #ds-{ds.dataset_id} .dataset-name', ds.name),
        then.link(f'main #ds-{ds.dataset_id} .project-link', href=href, text=prj.name),
        then.text(f'main #ds-{ds.dataset_id} .update-timestamp', since)
    )
