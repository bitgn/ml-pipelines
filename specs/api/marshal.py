from typing import List

from google.protobuf.message import Message

from . import events_pb2 as evt

map = {
    evt.ProjectCreated: evt.Type.Event_ProjectCreated,
    evt.DatasetCreated: evt.Type.Event_DatasetCreated,
    evt.DatasetUpdated: evt.Type.Event_DatasetUpdated,
    evt.Event_JobAdded: evt.Type.Event_JobAdded,
}


def serialize(msgs: List[Message]):
    result = []
    for m in msgs:
        bytes = m.SerializeToString()

        envelope = evt.Event(Body=bytes, Type=map[type(m)])
        result.append(envelope)
    return result
