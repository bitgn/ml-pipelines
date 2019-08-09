import datetime as dt
from typing import List, Callable, Optional

import bs4
from dataclasses import dataclass
import requests as r

class Env:
    def __init__(self):
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
    action: Callable[[r.Session, str], r.Response]
    text: str


@dataclass
class Then:
    action: Callable[[bs4.BeautifulSoup], Optional[str]]


@dataclass
class Scenario:
    when: When
    then: List[Then]
