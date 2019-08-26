from env import *
from test_api import events_pb2 as evt

def given_a_dataset_with_location_id(t: Env):
    """dataset headers are downgraded by one"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)

    ds.meta.description = "# HELLO"
    ds.meta.description_set = True

    t.given_events(prj, ds)

    t.scenario(
        when.view_dataset(ds.project_name, ds.name),
        then.text(f'main #ds-{ds.uid.hex()} .description h4', 'HELLO'),
    )
