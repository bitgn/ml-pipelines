from typing import List

from google.protobuf.message import Message

from . import events_pb2 as evt

map = {
    evt.ProjectCreated: evt.Event_ProjectCreated,
    evt.DatasetCreated: evt.Event_DatasetCreated,
    evt.DatasetUpdated: evt.Event_DatasetUpdated,
    evt.JobAdded: evt.Event_JobAdded,
    evt.ExpertAdded: evt.Event_ExpertAdded,
    evt.DatasetVersionAdded: evt.Event_DatasetVersionAdded,

}


def serialize(msgs: List[Message]):
    result = []
    for m in msgs:
        bytes = m.SerializeToString()

        typ = type(m)
        if not typ in map:
            raise KeyError(f'Unknown event type {typ.__name__}')

        envelope = evt.Event(Body=bytes, Type=map[typ])
        result.append(envelope)
    return result
