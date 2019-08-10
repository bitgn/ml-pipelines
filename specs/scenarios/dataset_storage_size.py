from env import *
from evil import pretty
from api import events_pb2 as evt


def _avoid_wrapping(value):
    """
    Avoid text wrapping in the middle of a phrase by adding non-breaking
    spaces where there previously were normal spaces.
    """
    return value.replace(" ", "\xa0")



def given_a_datasets_with_various_storage_sizes(t: env.Env):
    """show formatted storage value"""
    prj = preset.project_created(t)



    map = {
        0: '0 bytes',
        1: '1 byte',
        1000: '1.0 kB',
        1000000: '1.0 MB',
        1100000: '1.1 MB',
        3000000000: '3.0 GB',
        4100000000000: '4.1 TB',
    }


    conditionals = []
    conditionals.append(when.list_datasets())

    t.given_events(prj)

    for k,v in map.items():
        ds = preset.dataset_created(t, prj)
        ds.meta.storage_bytes = k
        value = _avoid_wrapping(v)
        conditionals.append(then.text(f'main #ds-{ds.dataset_id} .zip-size', value),)
        t.given_events(ds)



    t.scenario(*conditionals)






def given_a_dataset_with_storage_size(t: env.Env):
    """show formatted size on all screens"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)

    ds.meta.storage_bytes = 1100
    size = _avoid_wrapping('1.1 kB')

    t.given_events(prj, ds)
    t.scenario(
        when.view_dataset(ds.dataset_id),
        then.text(f'main #ds-{ds.dataset_id} .zip-size', size),
    )

    t.scenario(
        when.list_datasets(),
        then.text(f'main #ds-{ds.dataset_id} .zip-size', size),
    )

    t.scenario(
        when.view_project(prj.project_id),
        then.text(f'main #ds-{ds.dataset_id} .zip-size', size),
    )




def given_a_dataset_without_storage_size(t: Env):
    """show 0 bytes"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)

    ds.meta.storage_bytes_state = evt.DELETE

    t.given_events(prj, ds)

    size = _avoid_wrapping('0 bytes')

    t.scenario(
        when.view_dataset(ds.dataset_id),
        then.text(f'main #ds-{ds.dataset_id} .zip-size', size),
    )

    t.scenario(
        when.list_datasets(),
        then.text(f'main #ds-{ds.dataset_id} .zip-size', size),
    )

    t.scenario(
        when.view_project(prj.project_id),
        then.text(f'main #ds-{ds.dataset_id} .zip-size', size),
    )