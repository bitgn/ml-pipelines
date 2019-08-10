from datetime import timedelta

from env import *
from api import events_pb2 as evt



def _avoid_wrapping(value):
    """
    Avoid text wrapping in the middle of a phrase by adding non-breaking
    spaces where there previously were normal spaces.
    """
    return value.replace(" ", "\xa0")





def given_a_datasets_with_various_timestamps(t: env.Env):
    """show formatted storage value"""
    prj = preset.project_created(t)





    map = {
        timedelta(): 'just now',
        timedelta(seconds=1): 'just now',
        timedelta(seconds=2): '2 seconds ago',
        timedelta(minutes=1): 'a minute ago',
        timedelta(hours=1): 'an hour ago',
        timedelta(days=1): 'a day ago',
        timedelta(days=2): '2 days ago',
        timedelta(days=355): 'a year ago',
        timedelta(days=357): 'a year ago',
        timedelta(days=30): 'a month ago',
        timedelta(days=60): '2 months ago',

    }



    conditionals = []
    conditionals.append(when.list_datasets())

    t.given_events(prj)

    for k,v in map.items():
        ds = preset.dataset_created(t, prj)

        time = t.time - k
        ds.meta.update_timestamp = int(time.timestamp())
        ds.meta.update_timestamp_state = evt.STATE.VALUE


        value = _avoid_wrapping(v)
        conditionals.append(then.text(f'main #ds-{ds.dataset_id} .update-timestamp', value),)
        t.given_events(ds)



    t.scenario(*conditionals)

