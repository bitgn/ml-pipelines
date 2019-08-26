from env import *
from test_api import events_pb2 as evt


def given_a_populated_project(t: env.Env):
    """show formatted total storage value and entity count"""
    prj = preset.project_created(t)

    ds1 = preset.dataset_created(t, prj)
    ds2 = preset.dataset_created(t, prj)

    ds1.meta.storage_bytes = 1000
    ds2.meta.storage_bytes = 2000

    t.given_events(prj, ds1, ds2,
                   preset.job_added(t, prj),
                   preset.job_added(t, prj),
                   preset.job_added(t, prj)
                   )



    t.scenario(
        when.list_projects(),
        then.text(f'main #prj-{prj.uid.hex()} .storage-size', nbsp('3.0 kB'), title='3000'),
        then.text(f'main #prj-{prj.uid.hex()} .dataset-count', "2"),
        then.text(f'main #prj-{prj.uid.hex()} .job-count', "3")
    )


def given_an_empty_project(t: env.Env):
    """show proper zeroes"""
    prj = preset.project_created(t)
    t.given_events(prj)

    t.scenario(
        when.list_projects(),
        then.text(f'main #prj-{prj.uid.hex()} .storage-size', nbsp('0 bytes'), title='0'),
        then.text(f'main #prj-{prj.uid.hex()} .dataset-count', "0"),
        then.text(f'main #prj-{prj.uid.hex()} .job-count', "0")
    )