"""MLP-?????????"""
from env import *


def given_an_empty_project(t: Env):
    "show dataset placeholder"
    prj = preset.project_created(t)
    t.given_events(prj)

    t.scenario(
        when.view_project(prj.name),

        then.none("main .dataset-info")
    )


def given_a_project_with_a_dataset(t: Env):
    "show dataset in the list"
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)
    t.given_events(prj, ds)

    t.scenario(
        when.view_project(prj.name),
        then.count('main .dataset-info', 1),
        then.link(f'main #ds-{ds.uid.hex()} .dataset-link',
                  href=urls.view_dataset(ds.project_name, ds.name),
                  text=ds.name),
    )
