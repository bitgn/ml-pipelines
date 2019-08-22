"""MLP-?????????"""
from env import *


def given_an_empty_project(t: Env):
    "show dataset placeholder"
    prj = preset.project_created(t)
    t.given_events(prj)

    t.scenario(
        when.view_project(prj.project_id),

        then.none("main .dataset-info")
    )


def given_a_project_with_a_dataset(t: Env):
    "show dataset in the list"
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)
    t.given_events(prj, ds)

    t.scenario(
        when.view_project(prj.project_id),
        then.count('main .dataset-info', 1),
        then.text(f'main #ds-{ds.dataset_id} .dataset-link', ds.meta.name),
        then.link(f'main #ds-{ds.dataset_id} .dataset-link',
                  href=urls.view_dataset(ds.dataset_id),
                  text=ds.meta.name),
    )
