"""
Generate demo data from the web analytics domain
"""

import datetime
from dataclasses import dataclass
from typing import List
from test_api import events_pb2 as evt
from test_api import vo_pb2 as vo

from client import ml_pipelines as client

import faker


@dataclass
class Node:
    name: str
    title: str
    format: str
    file_count: int
    file_size: int
    record_count: int
    last_update: datetime.datetime
    description: str
    location_id: str
    ver: client.DatasetVersion


@dataclass
class Process:
    name: str
    title: str
    sources: List[Node]
    outputs: List[Node]


_counter = 0
f = faker.Faker()

storages = ['gcs-eu', 'gcs-us', 'vm-ds']


def _next_id():
    global _counter
    _counter += 1
    return _counter - 1


import random

random.seed(1)


def _node(title: str, format=None, size_mb=None, description=None, file_count=None, minutes=None) -> 'Node':
    name = f'analytics_n_{_next_id()}'
    n = Node(name, title, format, 0, 0, 0, None, description, 'gcp-vm1', b'')

    if 'tsv' in title:
        n.format = 'TSV'
        n.file_count = 1

        n.last_update = datetime.datetime.utcnow() - datetime.timedelta(hours=3)
        n.storage_id = 'gcp-eu'

    if 'metric' in title or 'Influx' in title:
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

api = _node(title='Analytics API')

rt2 = _node(title='InfluxDB: auction_perf')
rt3 = _node(title="InfluxDB: auction_wins")
rt4 = _node(title="InfluxDB: auction_errors", description=BID_ERROR_METRIC)

_proc([api], [rt2, rt3, rt4], name="Realtime processing")

fast = _node(title='Offline event-store', format='protobuf', size_mb=500000, minutes=4, file_count=442)
_proc([api], [fast], name='Ingestion')

s1 = _node(title="sessions_in_progress.lmdb", format='LMDB+PB')

s2 = _node(title="InfluxDB: user_sessions", format='InfluxDB', description=SESSION_METRIC)

gr = _node(title="grafana")
_proc([rt2, rt3, rt4, s2], [gr], name="reporting")

_proc([fast], [s1, s2], name="manage_sessions")

tsv1 = _node(title="lifetime_value.tsv", size_mb=10)
tsv2 = _node(title="lifecycle_category.tsv", size_mb=3)
tsv3 = _node(title="users_per_property.tsv", size_mb=10)
tsv4 = _node(title="revenue_per_property.tsv", size_mb=0.05)
tsv5 = _node(title="revenue_per_bidder.tsv", size_mb=0.05)
tsv6 = _node(title="revenue_per_device.tsv", size_mb=0.01)
tsv7 = _node(title="revenue_per_age.tsv", size_mb=0.1)
tsv8 = _node(title="revenue_per_device.tsv", size_mb=0.05)
tsv9 = _node(title="revenue_per_gender.tsv", size_mb=0.05)
tsv10 = _node(title="revenue_per_user.tsv", size_mb=20)
tsv11 = _node(title="cpm_per_age_by_property.tsv", size_mb=0.1)
tsv12 = _node(title="users_per_property_7d.tsv", size_mb=20)

_proc([fast], [tsv1, tsv2, tsv3, tsv4, tsv5, tsv6, tsv7, tsv8, tsv9, tsv10, tsv11, tsv12], name="store_scan_job")

r1 = _node(title="Revenue by property")
r2 = _node(title="Sessions by property")
r3 = _node(title="Ads by property")
r4 = _node(title="Revenue per bidder")
r5 = _node(title="Revenue per user")
r6 = _node(title="Weekly active users")
r7 = _node(title="User distribution by revenue segments")
r8 = _node(title="CPM distribution")
r9 = _node(title="Anonymous auctions")

_proc([tsv12], [r6], name="render_active_users")
_proc([tsv2], [r9], name="render_anonymous_users")
_proc([tsv4], [r1, r2, r3], name="render_proprty_stats")
_proc([tsv10], [r5], name="render_user_charts")

_proc([tsv5, tsv6], [], name="render_generic_daily_stats")
_proc([tsv10], [r7], name="render_user_groups")
_proc([tsv7], [r7], name="render_revenue_per_age")


def setup_analytics_demo(cl: client.Client):
    print("Start setting up")
    proj = cl.create_project("web_01", "Web Analytics")



    for n in _nodes:

        if not n.file_size:
            n.file_size = random.randint(10000, 100000)

        if n.description:
            n.description = n.description.strip(' \n\r')

        ds = proj.create_dataset(
            name=n.name,
            title=n.title,
            data_format=n.format,
            description=n.description,
            location_id=n.location_id,
        )

        ver = ds.get_last_version()
        stage = ver.prepare_commit()

        for i in range(n.file_count):
            count = n.file_size / n.file_count
            stage.add_file(name=f'file_{i}', records=int(n.record_count/n.file_count), size=int(count))

        inputs = []



        n.ver = stage.commit("initial commit")


    for p in _pipelines:
        inputs = [x.ver for x in p.sources]
        outputs = [x.ver for x in p.outputs]

        job = proj.create_job(name=p.name,title=p.title)

        run = job.start_run("Input run", inputs=inputs)
        run.log("Some work being done")
        run.complete()

    print("done")
