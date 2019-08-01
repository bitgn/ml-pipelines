from django import urls

from explore.specs import when, preset
from test import *
from utils import pretty




def given_a_dataset_with_metadata(t: test.Env):
    """show dataset properties"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)

    t.given_events(prj, ds)

    href = urls.reverse('explore:view_project', args=[ds.project_id])
    t.scenario(
        when.view_dataset(ds.dataset_id),
        then.link(f'main .project-link', href=href, text=prj.name),
    )

    t.scenario(
        when.list_datasets(),
        then.link(f'main #ds-{ds.dataset_id} .project-link', href=href, text=prj.name),

    )
