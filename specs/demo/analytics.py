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



rint = randint

import time


def setup_analytics_demo(cl: client.Client):
    print("Start setting up")
    prj = cl.add_project("web_01")

    tsvs = [
        "lifetime_value",
        "lifecycle_category",
        "users_per_property",
        "revenue_per_property",
        "revenue_per_bidder",
        "revenue_per_device",
        "revenue_per_age",
        "revenue_per_gender",
        "revenue_per_user",
        "cpm_per_age_by_property",
        "users_per_property_7d",
    ]

    now = int(time.time())

    for tsv in tsvs:
        ds = prj.add_dataset(tsv, summary="Sample dataset summary")



        age = rint(3600 * 2, 3600 * 24 * 30)
        begins = now - age
        done = begins + rint(int(age/2),age)



        ds.add_info(multiline_text="initial import started", timestamp_sec=begins)
        ds.add_success(multiline_text="initial import complete", timestamp_sec=done)


        for i in range(50):
            ds.add_info(multiline_text="Starting another run", timestamp_sec=done + i*5)
            ds.add_success(multiline_text="Completed another run", timestamp_sec=done+i*5+4)

    dataset = prj.get_or_create_dataset("revenue_per_device")
    dataset.add_info("Starting incremental import\nFound offsets [[0, 309], [1, 249], [2, 262], [3, 285], [4, 278], [5, 288], [6, 298], [7, 264], [8, 273], [9, 284], [10, 287], [11, 310]]", now-3600,)
    dataset.add_error("Failed to import", now -2000)













    print("done")
