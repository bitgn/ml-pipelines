import time
from dataclasses import dataclass
from typing import List

import grpc
from api import api_pb2_grpc, marshal
from api import api_pb2 as api
from evil.colors import *
import importlib
import os
import pathlib
import sys
import traceback
from inspect import getmembers, isfunction, getdoc
import env
import http

import requests
import bs4

root = 'scenarios'
sys.path.append(root)

file_fails = 0
file_ok = 0
assertion_fails = 0

channel = grpc.insecure_channel('localhost:50051')
stub = api_pb2_grpc.TestStub(channel)

webBase = "http://localhost:8000"


def wait_for_server_to_start():
    while True:
        try:
            stub.Ping(api.PingRequest())
            break
        except:
            time.sleep(0.1)
            continue


scenario_ok = 0
scenario_fail = 0
scenario_fail_count = 0

@dataclass
class ModuleResult:
    name:str
    loc:str
    doc:str
    funcs: List['FuncResult']

    def __init__(self, name, loc, doc):
        self.name = name
        self.loc = loc
        self.doc = doc or ""
        self.funcs = []

    def ok(self):
        if not self.funcs:
            return False
        for f in self.funcs:
            if not f.ok():
                return False
        return True



    def print(self):
        print(f'\n{CBOLD}{self.name}{CEND}: {self.doc or ""}')
        print(f'  {CBEIGE}{self.loc}{CEND}')

        if not self.funcs:
            print(f'  {CRED}No scenarios{CEND}')


        for fr in self.funcs:
            if not fr.ok():
                fr.print_fails()


@dataclass()
class ScenarioResult:
    when: str
    fails: List[str]

    def __init__(self):
        self.fails = []

    def ok(self):
        return not self.fails

@dataclass
class FuncResult:
    name:str
    doc:str
    scenarios:List[ScenarioResult]

    exception:List[str]




    def __init__(self, name, doc):
        self.name = name
        self.doc = doc
        self.scenarios = []
        self.exception = []

    def print_fails(self):
        print(f'  {CYELLOW}{self.name}{CEND} - {self.doc}')


        if self.exception:
            print(f'    {CRED}✗ {self.name}{CEND}:')
            print(f'      {CRED}{"      ".join(self.exception).rstrip()}{CEND}')
            return

        if not self.scenarios:
            print(f'  {CRED}✗ feature has no specs{CEND}')

        for s in self.scenarios:
            if s.ok():
                continue
            print(f'    {CBOLD}when {s.when}{CEND}')

            for fail in s.fails:
                print(f'      {CRED}{fail}{CEND}')

    def ok(self):
        if self.exception:
            return False
        if not self.scenarios:
            return False
        for r in self.scenarios:
            if not r.ok():
                return False

        return True






try:

    for l in os.listdir(root):
        stem = pathlib.Path(l).stem

        if stem.startswith('_') or stem in ['preset', 'when']:
            continue

        module = importlib.import_module(stem)

        doc = getdoc(module)
        loc = os.path.abspath(os.path.join(root, l))

        mr = ModuleResult(stem, loc, doc)


        count = 0
        module_fails = 0

        client = requests.sessions.Session()

        wait_for_server_to_start()

        for name, factory in getmembers(module, isfunction):
            if not name.startswith('given_'):
                continue
            fr = FuncResult(name, factory.__doc__)
            mr.funcs.append(fr)
            try:
                count += 1

                test_env = env.Env()
                factory(test_env)

                req = api.ScenarioRequest(Name=name)
                req.timestamp = int(test_env.time.timestamp())

                req.Events.extend(marshal.serialize(test_env.events))
                stub.Setup(req)

                for s in test_env.scenarios:

                    res = ScenarioResult()
                    fr.scenarios.append(res)
                    res.when = s.when.text

                    #

                    response = s.when.action(client, webBase)

                    soup = bs4.BeautifulSoup(response.content, 'lxml')
                    if response.status_code != http.HTTPStatus.OK:
                        if soup.title:
                            res.fails.append(f'{CBOLD}{response.reason}: {soup.title.text}')
                        elif soup.p.text:
                            res.fails.append(f'{CBOLD}{response.reason}: {soup.p.text.strip()}')
                        else:
                            res.fails.append(f'{CBOLD}{response.reason}: {soup}')
                    else:
                        if s.then:

                            for a in s.then:
                                result = a.action(soup)
                                if result:
                                    res.fails.append(result)
                        else:
                            res.fails.append("no expectations provided")

                    if res.ok():
                        scenario_ok += 1
                    else:
                        scenario_fail += 1
                        scenario_fail_count += len(res.fails)

                    module_fails += len(res.fails)


            except:
                fr.exception= traceback.format_exception(*sys.exc_info(), limit=None, chain=True)




        if count == 0 or module_fails > 0:
            file_fails += 1
        else:
            file_ok += 1

        if not mr.ok():
            mr.print()

    if file_fails:
        print(f'{CRED}✗ File FAILS {file_fails}{CEND}')
    if file_ok:
        print(f'{CGREEN}✔ File OK {file_ok}{CEND}')

    print(f'Scenarios: {scenario_ok} OK, {scenario_fail} fail (fix {scenario_fail_count})')
finally:
    try:
        stub.Kill(api.KillRequest())
    except:
        pass
