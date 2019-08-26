import datetime as dt
import uuid
from typing import List, Callable, Optional, Any
import bs4
from dataclasses import dataclass
import requests as r
from faker import Faker

import struct
from client import ml_pipelines as client

class Env:
    def __init__(self):
        self.faker = Faker()
        self.faker.seed(1)
        self.events = []
        self.counter = 0
        self.time = dt.datetime.utcnow()
        self.scenarios:List[Scenario] = []
        pass

    def given_events(self, *e):
        self.events.extend(e)
        pass

    def scenario(self, *args):

        when = None
        then = []
        for a in args:
            if isinstance(a, When):
                when = a
            if isinstance(a, Then):
                then.append(a)

        self.scenarios.append(Scenario(when=when, then=then))

    def next_uid(self):
        id = self.counter
        arr = bytearray(16)

        struct.pack_into("<I", arr, 0, id)

        self.counter += 1
        return bytes(arr)

    def next_id(self):

        id = self.counter

        self.counter += 1
        return id


@dataclass
class When:
    web_action: Optional[Callable[[r.Session, str], r.Response]]
    client_action: Optional[Callable[[client.Client], Any]]
    text: str


@dataclass
class Then:
    web_action: Optional[Callable[[bs4.BeautifulSoup], Optional[str]]]
    client_action: Optional[Callable[[Any], Optional[str]]]


@dataclass
class Scenario:
    when: When
    then: List[Then]
