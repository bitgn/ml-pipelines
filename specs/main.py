import argparse
import time
from dataclasses import dataclass
from typing import List

import grpc
from test_api import api_pb2_grpc, marshal
from test_api import api_pb2 as api
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


parser = argparse.ArgumentParser(description='Process some integers.')

parser.add_argument("--web", action="store", dest="web", default="localhost:8080")
parser.add_argument("--grpc", action="store", dest="grpc", default="localhost:9111")

args = parser.parse_args()


root = 'scenarios'
sys.path.append(root)

file_fails = 0
file_ok = 0
assertion_fails = 0

channel = grpc.insecure_channel(args.grpc)
stub = api_pb2_grpc.TestStub(channel)

catalog = api_pb2_grpc.CatalogStub(channel)

webBase = f"http://{args.web}"


def wait_for_server_to_start():
    while True:
        try:
            stub.Ping(api.PingRequest())
            break
        except:
            time.sleep(0.1)
            continue

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


@dataclass
class SuiteResult:
    modules: List[ModuleResult]

    def __init__(self):
        self.modules = []

    def ok(self):
        if not self.modules:
            return False
        for f in self.modules:
            if not f.ok():
                return False
        return True



    def print(self):

        assert_count = 0
        assert_ok = 0

        scenarios_count = 0
        scenarios_ok = 0

        spec_count = 0
        spec_ok = 0

        for m in self.modules:

            for f in m.funcs:
                spec_count +=1
                if f.ok():
                    spec_ok +=1
                for s in f.scenarios:
                    assert_count += s.assert_count
                    assert_ok += s.assert_ok

                    scenarios_count +=1
                    if s.assert_count == s.assert_ok:
                        scenarios_ok +=1


        print()
        if spec_ok == spec_count:
            print(f'{CGREEN}✔ Specs:      {spec_ok} of {spec_count} OK{CEND}')
        else:
            print(f'{CRED}✗ Specs:      {spec_ok} of {spec_count} FAIL{CEND}')
        if assert_count == assert_ok:
            print(f"{CGREEN}✔ Assertions: {assert_ok} of {assert_count} OK{CEND}")
        else:
            print(f"{CRED}✗ Assertions: {assert_ok} of {assert_count} FAIL{CEND}")


@dataclass()
class ScenarioResult:
    when: str
    fails: List[str]
    assert_count:int
    assert_ok:int


    def __init__(self):
        self.fails = []
        self.assert_ok = 0
        self.assert_count = 0

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

    suite = SuiteResult()

    wait_for_server_to_start()
    print("Server ready!")

    for l in os.listdir(root):
        stem = pathlib.Path(l).stem

        if stem.startswith('_') or stem in ['preset', 'when']:
            continue

        module = importlib.import_module(stem)

        doc = getdoc(module)
        loc = os.path.abspath(os.path.join(root, l))

        mr = ModuleResult(stem, loc, doc)

        suite.modules.append(mr)



        client = requests.sessions.Session()

        for name, factory in getmembers(module, isfunction):
            if not name.startswith('given_'):
                continue
            fr = FuncResult(name, factory.__doc__)
            mr.funcs.append(fr)
            try:

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

                    res.assert_count = len(s.then)

                    if s.when.web_action:
                        response = s.when.web_action(client, webBase)

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
                                        res.assert_ok += 1
                            else:
                                res.fails.append("no expectations provided")


                    elif s.when.client_action:
                        response = s.when.client_action(catalog)
                        #print(type(response))
                        if isinstance(response, grpc.RpcError):
                            typed = grpc.RpcError(response)

                            res.fails.append(f'{CBOLD}{typed}{CEND}')
                        else:

                            if s.then:
                                for a in s.then:
                                    result = a.action(response)

                                    if result:
                                        res.fails.append(result)
                                    else:
                                        res.assert_ok += 1

                            else:
                                res.fails.append("no expectations provided")






            except (ConnectionRefusedError, requests.exceptions.ConnectionError):
                print(f"{CBOLD}{CRED}Connection refused!{CEND}")
                # let server print its log
                time.sleep(1)
                quit(1)


            except:
                fr.exception= traceback.format_exception(*sys.exc_info(), limit=None, chain=True)





        if not mr.ok():
            mr.print()

    suite.print()
finally:
    try:
        stub.Kill(api.KillRequest())
    except:
        pass
