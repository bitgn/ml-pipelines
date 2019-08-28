import argparse
import json
import time
from dataclasses import dataclass
from typing import List

import grpc

import test_api as api

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

import client.ml_pipelines as client
import subprocess

parser = argparse.ArgumentParser(description='Process some integers.')

parser.add_argument("--web", action="store", dest="web", default="localhost:8080")
parser.add_argument("--grpc", action="store", dest="grpc", default="localhost:9111")

parser.add_argument("--executable", action="store", dest="executable", required=True)
parser.add_argument("--template-path", action="store", dest="template_path", required=True)
args = parser.parse_args()


root = os.path.join(os.path.dirname(__file__), 'scenarios')
sys.path.append(root)

file_fails = 0
file_ok = 0
assertion_fails = 0

channel = grpc.insecure_channel(args.grpc)
test_service = api.TestStub(channel)
app_client = client.connect(args.grpc)





webBase = f"http://{args.web}"


def wait_for_server_to_start():
    while True:
        try:
            test_service.Ping(api.PingRequest())
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




        join = os.path.join(root, "stats.json")
        st = None
        if os.path.exists(join):
            with open(join, 'r') as o:
                st = json.load(o)
        else:
            st = { 'assert_ok': 0, 'spec_ok':0 }



        delta_assert = assert_ok - st['assert_ok']
        delta_specs = spec_ok - st['spec_ok']


        print("             FAIL   OK    Δ")


        def delta(n):
            color = ""
            if n < 0:
                color=CRED
            if n > 0:
                color = CGREEN
            return f'{color}{n:4d}{CEND}'

        def bad(n):
            if n > 0:
                return f'{CRED}{n:4d}{CEND}'
            return f'{CGREEN}{n:4d}{CEND}'

        def good(n):
            if n < 0:
                return f'{CRED}{n:4d}{CEND}'
            return f'{CGREEN}{n:4d}{CEND}'


        if spec_ok == spec_count:
            print(f'{CGREEN}✔ Specs{CEND}      ', end='')
        else:
            print(f'{CRED}✗ Specs{CEND}      ', end='')


        print(f"{bad(spec_count-spec_ok)} {good(spec_ok)} {delta(delta_specs)}")


        if assert_count == assert_ok:
            print(f"{CGREEN}✔ Assertions{CEND} ", end='')
        else:
            print(f"{CRED}x Assertions{CEND} ", end='')

        print(f"{bad(assert_count-assert_ok)} {good(assert_ok)} {delta(delta_assert)}")


        with open(join, 'w') as o:
            json.dump({'assert_ok':assert_ok, 'spec_ok':spec_ok}, o)


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

    nested = [
        args.executable,
        "--specs",
        "--grpc", args.grpc,
        "--web", args.web,
        "--template-path", args.template_path,
        # we are going to wipe things anyway
        # TODO: wipe db instead
        "--upgrade", "none"
                             ]
    print(f"Launching server: {CYELLOW}{' '.join(nested)}{CEND}")

    proc = subprocess.Popen(nested)



    wait_for_server_to_start()
    print("Server ready!")

    for l in os.listdir(root):

        if not l.endswith('.py'):
            continue
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

                req.Events.extend(api.serialize(test_env.events))
                test_service.Setup(req)

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
                                    result = a.web_action(soup)
                                    if result:
                                        res.fails.append(result)
                                    else:
                                        res.assert_ok += 1
                            else:
                                res.fails.append("no expectations provided")


                    elif s.when.client_action:
                        response = s.when.client_action(app_client)

                        if s.then and s.then:
                            for a in s.then:
                                result = a.client_action(response)

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
        test_service.Kill(api.KillRequest())
    except:
        pass

