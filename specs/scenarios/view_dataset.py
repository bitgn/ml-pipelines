"""MLP-???"""
from env import *


def given_a_dataset_with_metadata(t: Env):
    """show dataset properties"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)

    t.given_events(prj, ds)

    t.scenario(
        when.view_dataset(prj.name, ds.name),
        then.link('main .project-link',
                  href=urls.view_project(prj.name),
                  text=prj.name),
    )
