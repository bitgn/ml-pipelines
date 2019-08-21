import datetime as dt
from typing import List, Callable, Optional
from test_api import api_pb2_grpc as api
import bs4
from dataclasses import dataclass
import requests as r
from google.protobuf.message import Message
from faker import Faker


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

    def next_id(self):
        self.counter += 1
        return self.counter - 1


@dataclass
class When:
    web_action: Optional[Callable[[r.Session, str], r.Response]]
    client_action: Optional[Callable[[api.CatalogStub], Message]]
    text: str


@dataclass
class Then:
    action: Callable[[bs4.BeautifulSoup], Optional[str]]


@dataclass
class Scenario:
    when: When
    then: List[Then]
