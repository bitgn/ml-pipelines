from env import *
from test_api import vo_pb2 as vo


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
        ver = preset.dataset_version_added(t, ds)

        del ver.items[:]
        ver.items.append(vo.DatasetItem(name='file', storage_bytes=k, records=1, uid=t.next_uid()))

        value = _avoid_wrapping(v)
        conditionals.append(then.text(f'main #ds-{ds.uid.hex()} .zip-size', value),)
        t.given_events(ds, ver)



    t.scenario(*conditionals)






def given_a_dataset_with_a_single_version(t: env.Env):
    """show formatted size on all screens"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)

    ver = preset.dataset_version_added(t, ds)
    del ver.items[:]
    ver.items.append(vo.DatasetItem(name='file', storage_bytes=1100, records=1, uid=t.next_uid()))

    size = _avoid_wrapping('1.1 kB')
    title="1100"

    t.given_events(prj, ds, ver)
    t.scenario(
        when.view_dataset(ds.project_name, ds.name),
        then.text(f'main #ds-{ds.uid.hex()} .zip-size', size, title=title),
    )

    t.scenario(
        when.list_datasets(),
        then.text(f'main #ds-{ds.uid.hex()} .zip-size', size, title=title),
    )

    t.scenario(
        when.view_project(prj.name),
        then.text(f'main #ds-{ds.uid.hex()} .zip-size', size, title=title),
    )


def given_a_dataset_with_multiple_additive_versions(t: env.Env):
    """render final size"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)

    v1 = preset.dataset_version_added(t, ds)
    del v1.items[:]
    v1.items.append(vo.DatasetItem(name='file1', storage_bytes=1100, records=1, uid=t.next_uid()))

    v2 = preset.dataset_version_added(t, ds, v1)
    del v2.items[:]
    v2.items.append(vo.DatasetItem(name='file2', storage_bytes=2200, records=1, uid=t.next_uid()))




    t.given_events(prj, ds, v1, v2)
    t.scenario(
        when.view_dataset(ds.project_name, ds.name),
        then.text(f'main #ds-{ds.uid.hex()} .zip-size', title="3300"),
    )

    t.scenario(
        when.list_datasets(),
        then.text(f'main #ds-{ds.uid.hex()} .zip-size', title="3300"),
    )

    t.scenario(
        when.view_project(prj.name),
        then.text(f'main #ds-{ds.uid.hex()} .zip-size', title="3300"),
    )

def given_a_dataset_with_multiple_mutating_versions(t: env.Env):

    """show formatted size on all screens"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)

    v1 = preset.dataset_version_added(t, ds)
    del v1.items[:]
    item1 = vo.DatasetItem(name='file1', storage_bytes=1100, records=1, uid=t.next_uid())
    item2 = vo.DatasetItem(name='file2', storage_bytes=2200, records=1, uid=t.next_uid())

    v1.items.extend([item1,item2])

    v2 = preset.dataset_version_added(t, ds, v1)
    del v2.items[:]
    del v2.remove[:]

    v2.remove.append(item1)

    t.given_events(prj, ds, v1, v2)
    t.scenario(
        when.view_dataset(ds.project_name, ds.name),
        then.text(f'main #ds-{ds.uid.hex()} .zip-size', title="2200"),
    )

    t.scenario(
        when.list_datasets(),
        then.text(f'main #ds-{ds.uid.hex()} .zip-size', title="2200"),
    )

    t.scenario(
        when.view_project(prj.name),
        then.text(f'main #ds-{ds.uid.hex()} .zip-size', title="2200"),
    )



def given_a_dataset_without_storage_size(t: Env):
    """show 0 bytes"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)



    t.given_events(prj, ds)

    size = _avoid_wrapping('0 bytes')

    t.scenario(
        when.view_dataset(ds.project_name, ds.name),
        then.text(f'main #ds-{ds.uid.hex()} .zip-size', size),
    )

    t.scenario(
        when.list_datasets(),
        then.text(f'main #ds-{ds.uid.hex()} .zip-size', size),
    )

    t.scenario(
        when.view_project(prj.name),
        then.text(f'main #ds-{ds.uid.hex()} .zip-size', size),
    )