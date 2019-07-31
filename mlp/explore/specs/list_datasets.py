"""MLP-2"""
from django import urls

from explore.specs import when, preset
from test import *

from utils import pretty
import datetime as dt


def given_no_datasets(e: test.Env):
    """show an actionable message"""

    e.scenario(
        when.list_datasets(),
        then.exists('main #empty-catalog-message'),
        then.none('main .dataset-info'),
    )


def given_a_dataset(t: test.Env):
    """show dataset properties"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)

    t.given_events(prj, ds)
    past = dt.datetime.utcfromtimestamp(ds.metadata.update_timestamp)
    now = t.time

    since = pretty.timedelta(now - past)

    href = urls.reverse('explore:view_project', args=[ds.project_id])
    pretty_bytes = pretty.bytes(ds.metadata.zip_bytes)
    t.scenario(
        when.list_datasets(),
        then.none('main #empty-catalog-message'),
        then.count('main .dataset-info', 1),
        then.text(f'main #ds-{ds.dataset_id} .dataset-name', ds.name),
        then.text(f'main #ds-{ds.dataset_id} .zip-size', pretty_bytes),
        then.link(f'main #ds-{ds.dataset_id} .project-link', href=href, text=prj.name),
        then.text(f'main #ds-{ds.dataset_id} .update-timestamp', since)
    )
