"""MLP-???"""
from django import urls

from explore.specs import when, preset
from test import *
from utils import pretty


def given_a_dataset_with_metadata(t: test.Env):
    """show dataset properties"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)

    t.given_events(prj, ds)

    t.scenario(
        when.view_dataset(ds.dataset_id),
        then.link('main .project-link',
                  href=urls.reverse('explore:view_project', args=[ds.project_id]),
                  text=prj.name),

    )
