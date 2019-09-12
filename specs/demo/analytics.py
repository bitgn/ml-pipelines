"""
Generate demo data from the web analytics domain
"""

import datetime as dt

from random import randint
from typing import List
from test_api import events_pb2 as evt
from test_api import vo_pb2 as vo

from client import ml_pipelines as client

import faker



N_API = 'api'
N_ES_OFFLINE='es-offline'
N_ES_BUFFER='es-buffer'

rint = randint


def import_into_fast_storage(prj: client.Project, i):

    job = prj.get_job('ingest-fast')

    input = prj.systems.get('event-store')

    output = prj.get_dataset('fast-store')


    t=int((dt.datetime.utcnow() - dt.timedelta(days=105-i)).timestamp())


    # import into the fast storage
    run = job.start_run(inputs=[input], timestamp=t)

    try:
        run.log(f"Starting run {i}", timestamp=t)
        run.log_version_info()


        staging = output.get_last_version().prepare_commit()
        staging.add_input(run)


        staging.add_file(f'file_{i}', records=rint(10000, 20000), size=rint(10000000, 80000000))
        ver = staging.commit(timestamp=t)

        if (i % 17 == 0) or (i % 23) == 0:
            raise ValueError("Commit problem")

        run.log(f'Commit {ver.uid.hex()}', timestamp=t)


        run.complete(timestamp=t)
    except:
        run.fail("Something went wrong", timestamp=t)



def scan_fast_store(prj: client.Project, i):
    tsvs = [
        ("lifetime_value", 10),
        ("lifecycle_category", 3),
        ("lifetime_value", 10),
        ("lifecycle_category", 3),
        ("users_per_property", 10),
        ("revenue_per_property", 0.05),
        ("revenue_per_bidder", 0.05),
        ("revenue_per_device", 0.01),
        ("revenue_per_age", 0.1),
        ("revenue_per_device", 0.05),
        ("revenue_per_gender", 0.05),
        ("revenue_per_user", 20),
        ("cpm_per_age_by_property", 0.1),
        ("users_per_property_7d", 20),
    ]
    source = prj.get_dataset('fast-store').get_last_version()
    job = prj.get_or_create_job("scan_event_store")


    dag_id = "something"
    run = job.start_run(inputs=[source], title=dag_id)


    stage = prj.get_or_create_dataset("aggregated-data").get_last_version().prepare_commit(clean_slate=True)
    stage.add_input(run)

    for tsv, size_mb in tsvs:
        stage.add_file(f'{tsv}.tsv', size=int(size_mb * 1024 * 1024), records=rint(100, 200000))
    stage.commit()
    run.complete()




rendered_reports = [
"Revenue by property",
"Sessions by property",
"Ads by property",
"Revenue per bidder",
"Revenue per user",
"Weekly active users",
"User distribution by revenue segments",
"CPM distribution",
"Anonymous auctions"
]

rendered_reports = [(x, x.lower().replace(' ','-') + '-rpt') for x in rendered_reports]



realtime_reports = [
"Auction performance",
"Auction wins",
    "Auction errors",
    "Revenue by device",
    "Revenue by property",
    "Bidder performance",
    "Active users"
]
realtime_reports = [(x, x.lower().replace(' ','-')+'-rt-rpt') for x in realtime_reports]


def render_reports(prj: client.Project, i):
    input = prj.get_dataset('aggregated-data').get_last_version()
    job = prj.get_or_create_job(f'render-reports')
    run = job.start_run(inputs=[input])


    for title, name in rendered_reports:
        report = prj.systems.get_or_add_report(name)
        report.add_version(f'v{i}', inputs=[run])

    run.fail("Something went wrong")




def setup_analytics_demo(cl: client.Client):
    print("Start setting up")
    prj = cl.add_project("web_01", "Web Analytics")




    slow = prj.systems.add_db('event-store', title="Event Store")
    # system versions can have inputs and outputs
    slow.add_version('v1')

    api = prj.systems.add_service('api', title="Analytics API")
    api.add_version('v1', outputs=[slow])


    prj.create_job("ingest-fast", title="update reporting store")
    prj.add_dataset("fast-store", title="Reporting store", data_format="FB+GZ")
    prj.add_dataset('aggregated-data', title="Aggregated Data", data_format="TSV+GZ")


    prj.create_job("scan_event_store", "Run report aggregation")
    prj.create_job("render-reports", "Render reports")



    # session management system - near realtime
    sessions = prj.systems.add_db('sessions.lmdb')  # unmanaged dataset?
    influx_user_sessions = prj.systems.add_table('influx-user-sessions') # unmanaged dataset?

    svc = prj.systems.add_service('session-tracking')
    svc.add_version('v1', inputs=[api], outputs=[sessions, influx_user_sessions])

    # rendered reports

    for title, name in rendered_reports:
        prj.systems.add_report(name, title=title)


    # realtime

    rpts = []
    for title, name in realtime_reports:
        svc = prj.systems.add_report(name, title=title)
        rpts.append(svc)

    rtp = prj.systems.add_service("realtime-processing", title="Realtime processing")
    rtp.add_version('v1', inputs=[api], outputs=rpts)



    for i in range(100):
        import_into_fast_storage(prj, i)
        scan_fast_store(prj, i) # saves into datasets
        render_reports(prj, i)














    print("done")
