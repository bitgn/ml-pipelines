"""MLP-???"""
from env import *


def given_a_dataset_with_metadata(t: Env):
    """show dataset properties"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)

    t.given_events(prj, ds)

    t.scenario(
        when.view_dataset(ds.dataset_id),
        then.link('main .project-link',
                  href=urls.view_project(ds.project_id),
                  text=prj.meta.name),

    )
