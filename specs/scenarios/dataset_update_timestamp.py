from datetime import timedelta

from env import *
from test_api import events_pb2 as evt


def _avoid_wrapping(value):
    """
    Avoid text wrapping in the middle of a phrase by adding non-breaking
    spaces where there previously were normal spaces.
    """
    return value.replace(" ", "\xa0")


def given_a_datasets_with_various_timestamps(t: env.Env):
    """show formatted storage value"""
    prj = preset.project_created(t)

    assertions = {
        timedelta(): 'just now',
        timedelta(seconds=1): 'just now',
        timedelta(seconds=2): '2 seconds ago',
        timedelta(minutes=1): 'a minute ago',
        timedelta(hours=1): 'an hour ago',
        timedelta(days=1): 'a day ago',
        timedelta(days=2): '2 days ago',
        timedelta(days=365): 'a year ago',
        timedelta(days=380): 'a year ago',
        timedelta(days=30): 'a month ago',
        timedelta(days=60): '2 months ago',
        timedelta(days=7): 'a week ago',
        timedelta(days=14): '2 weeks ago',
    }

    conditionals = []
    conditionals.append(when.list_datasets())

    t.given_events(prj)

    for k, v in assertions.items():
        ds = preset.dataset_created(t, prj)

        time = t.time - k
        ds.meta.update_timestamp = int(time.timestamp())
        ds.meta.update_timestamp_set = True

        value = _avoid_wrapping(v)
        conditionals.append(then.text(f'main #ds-{ds.dataset_id} .update-timestamp', value, hint=k))
        t.given_events(ds)

    t.scenario(*conditionals)


def given_a_dataset_with_update_timestamp(t: env.Env):
    """show formatted value on all screens"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)

    ds.meta.update_timestamp = int((t.time - timedelta(days=1)).timestamp())
    ds.meta.update_timestamp_set = True

    value = _avoid_wrapping('a day ago')

    t.given_events(prj, ds)
    t.scenario(
        when.view_dataset(ds.dataset_id),
        then.text(f'main #ds-{ds.dataset_id} .update-timestamp', value),
    )

    t.scenario(
        when.list_datasets(),
        then.text(f'main #ds-{ds.dataset_id} .update-timestamp', value),
    )

    t.scenario(
        when.view_project(prj.project_id),
        then.text(f'main #ds-{ds.dataset_id} .update-timestamp', value),
    )


def given_a_dataset_without_update_timestamp(t: Env):
    """show 0 bytes"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)

    ds.meta.update_timestamp = 0
    ds.meta.update_timestamp_set = True

    t.given_events(prj, ds)



    t.scenario(
        when.view_dataset(ds.dataset_id),
        then.text(f'main #ds-{ds.dataset_id} .update-timestamp', 'never'),
    )

    t.scenario(
        when.list_datasets(),
        then.text(f'main #ds-{ds.dataset_id} .update-timestamp', ''),
    )

    t.scenario(
        when.view_project(prj.project_id),
        then.text(f'main #ds-{ds.dataset_id} .update-timestamp', ''),
    )
