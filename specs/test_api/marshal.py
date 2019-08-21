from typing import List

from google.protobuf.message import Message

from . import events_pb2 as evt

map = {
    evt.ProjectCreated: evt.Event_ProjectCreated,
    evt.DatasetCreated: evt.Event_DatasetCreated,
    evt.DatasetUpdated: evt.Event_DatasetUpdated,
    evt.JobAdded: evt.Event_JobAdded,
    evt.ExpertAdded: evt.Event_ExpertAdded,
}


def serialize(msgs: List[Message]):
    result = []
    for m in msgs:
        bytes = m.SerializeToString()

        envelope = evt.Event(Body=bytes, Type=map[type(m)])
        result.append(envelope)
    return result
