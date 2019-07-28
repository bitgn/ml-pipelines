"""
Generate demo data from the web analytics domain
"""


import datetime
from dataclasses import dataclass
from typing import List
from proto import events_pb2 as evt



@dataclass
class Node:
    id: str
    name: str
    format: str
    file_count: int
    file_size: int
    record_count: int
    last_update: datetime.datetime
    description: str


@dataclass
class Process:
    id: str
    name: str
    sources: List[Node]
    outputs: List[Node]


_counter = 0


def _next_id():
    global _counter
    _counter += 1
    return _counter - 1


import random

random.seed(1)


def _node(name: str, format=None, size_mb=None, description=None, file_count=None, minutes=None) -> 'Node':
    id = f'analytics_n_{_next_id()}'
    n = Node(id, name, format, 0, 0, 0, None, description)

    if 'tsv' in name:
        n.format = 'TSV'
        n.file_count = 1

        n.last_update = datetime.datetime.utcnow() - datetime.timedelta(hours=3)

    if 'metric' in name or 'Inxlux' in name:
        n.format = 'INFLUX_DB'
        n.last_update = datetime.datetime.utcnow()

    if file_count:
        n.file_count = file_count

    if size_mb:
        min = int(size_mb * 0.8 * (1 << 20))
        max = int(size_mb * 1.1 * (1 << 20))

        n.file_size = random.randint(min, max)
        n.record_count = int(n.file_size / random.randint(20, 50))

    if minutes:
        n.last_update = datetime.datetime.utcnow() - datetime.timedelta(minutes=minutes)

    _nodes.append(n)
    return n


def _proc(sources: List[Node], targets: List[Node], name: str):
    id = f'analytics_p_{_next_id()}'
    p = Process(id, name, sources, targets)
    _pipelines.append(p)


SESSION_METRIC = """# User Sessions

_Hypertext Transfer Protocol_ (HTTP) is stateless: a client computer running a web browser
must establish a new _Transmission Control Protocol_ (TCP) network connection to the web 
server with each new HTTP `GET` or `POST` request. The web server, therefore, cannot 
rely on an established TCP network connection for longer than a single HTTP `GET` or `POST` operation. 

Session management is the technique used by the web developer to make the stateless HTTP 
protocol support session state. 

For example, once a user has been authenticated to the web server, the user's next HTTP request 
(`GET` or `POST`) should not cause the web server to ask for the user's account and password again. 
For a discussion of the methods used to accomplish this see _HTTP cookie_ and _Session ID_.
```"""



BID_ERROR_METRIC = """# Bid Errors

Side channel for returning any errors"""





_nodes: List[Node] = []
_pipelines: List[Process] = []

api = _node(name='Analytics API')

rt2 = _node(name='InfluxDB: auction_perf')
rt3 = _node(name="InfluxDB: auction_wins")
rt4 = _node(name="InfluxDB: auction_errors", description=BID_ERROR_METRIC)

_proc([api], [rt2, rt3, rt4], name="Realtime processing")

fast = _node(name='Offline event-store', format='protobuf', size_mb=500000, minutes=4, file_count=442)
_proc([api], [fast], name='Ingestion')

s1 = _node(name="sessions_in_progress.lmdb", format='LMDB+PB')

s2 = _node(name="InfluxDB: user_sessions", format='InfluxDB', description=SESSION_METRIC)

gr = _node(name="Grafana")
_proc([rt2, rt3, rt4, s2], [gr], name="reporting")

_proc([fast], [s1, s2], name="manage_sessions")

tsv1 = _node(name="lifetime_value.tsv", size_mb=10)
tsv2 = _node(name="lifecycle_category.tsv", size_mb=3)
tsv3 = _node(name="users_per_property.tsv", size_mb=10)
tsv4 = _node(name="revenue_per_property.tsv", size_mb=0.05)
tsv5 = _node(name="revenue_per_bidder.tsv", size_mb=0.05)
tsv6 = _node(name="revenue_per_device.tsv", size_mb=0.01)
tsv7 = _node(name="revenue_per_age.tsv", size_mb=0.1)
tsv8 = _node(name="revenue_per_device.tsv", size_mb=0.05)
tsv9 = _node(name="revenue_per_gender.tsv", size_mb=0.05)
tsv10 = _node(name="revenue_per_user.tsv", size_mb=20)
tsv11 = _node(name="cpm_per_age_by_property.tsv", size_mb=0.1)
tsv12 = _node(name="users_per_property_7d.tsv", size_mb=20)

_proc([fast], [tsv1, tsv2, tsv3, tsv4, tsv5, tsv6, tsv7, tsv8, tsv9, tsv10, tsv11, tsv12], name="store_scan_job")

r1 = _node(name="Revenue by property")
r2 = _node(name="Sessions by property")
r3 = _node(name="Ads by property")
r4 = _node(name="Revenue per bidder")
r5 = _node(name="Revenue per user")
r6 = _node(name="Weekly active users")
r7 = _node(name="User distribution by revenue segments")
r8 = _node(name="CPM distribution")
r9 = _node(name="Anonymous auctions")

_proc([tsv12], [r6], name="render_active_users")
_proc([tsv2], [r9], name="render_anonymous_users")
_proc([tsv4], [r1, r2, r3], name="render_proprty_stats")
_proc([tsv10], [r5], name="render_user_charts")

_proc([tsv5, tsv6], [], name="render_generic_daily_stats")
_proc([tsv10], [r7], name="render_user_groups")
_proc([tsv7], [r7], name="render_revenue_per_age")



def setup_analytics_demo():

    project = evt.ProjectCreated(project_id='web_01', name="Web Analytics")
    yield project



    for n in _nodes:

        meta = evt.DatasetMetadata()

        if n.format:
            meta.data_format = n.format
            meta.set_fields.append(evt.FIELD_DATA_FORMAT)

        if n.file_count:
            meta.file_count = n.file_count
            meta.set_fields.append(evt.FIELD_FILE_COUNT)

        if n.file_size:
            meta.raw_bytes = n.file_size
            meta.zip_bytes = n.file_size
            meta.set_fields.append(evt.FIELD_RAW_BYTES)
            meta.set_fields.append(evt.FIELD_ZIP_BYTES)
        if n.record_count:
            meta.record_count = n.record_count
            meta.set_fields.append(evt.FIELD_RECORD_COUNT)

        if n.last_update:
            meta.update_timestamp = int(n.last_update.timestamp())
            meta.set_fields.append(evt.FIELD_UPDATE_TIMESTAMP)

        if n.description:
            meta.description = n.description.strip(' \n\r')
            meta.set_fields.append(evt.FIELD_DESCRIPTION)

        ds = evt.DatasetCreated(dataset_id=n.id, name=n.name, project_id=project.project_id, metadata=meta)
        yield ds

    for p in _pipelines:
        inputs = [x.id for x in p.sources]
        outputs = [x.id for x in p.outputs]

        link = evt.JobAdded(job_id=p.id, job_name=p.name, inputs=inputs, outputs=outputs, project_id=project.project_id)
        yield link
