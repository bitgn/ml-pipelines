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

scenario_ok = 0
scenario_fail = 0
scenario_fail_count = 0

for l in os.listdir(root):
    stem = pathlib.Path(l).stem

    if stem.startswith('_') or stem in ['preset', 'when']:
        continue

    module = importlib.import_module(stem)

    doc = getdoc(module)

    print(f'\n{CBOLD}{stem}{CEND}: {doc or ""}')

    count = 0
    module_fails = 0



    client = requests.sessions.Session()

    for name, factory in getmembers(module, isfunction):
        if not name.startswith('given_'):
            continue
        try:
            count += 1

            test_env = env.Env()
            factory(test_env)

            req = api.ScenarioRequest(Name=name)
            req.timestamp = int(test_env.time.timestamp())

            req.Events.extend(marshal.serialize(test_env.events))
            stub.Setup(req)

            print(f'  {CYELLOW}{name}{CEND} - {factory.__doc__}')

            for s in test_env.scenarios:
                fails = []

                response = s.when.action(client, webBase)

                soup = bs4.BeautifulSoup(response.content, 'lxml')
                if response.status_code != http.HTTPStatus.OK:
                    if soup.title:
                        fails.append(f'{response.reason}: {soup.title.text}')
                    elif soup.p.text:
                        fails.append(f'{response.reason}: {soup.p.text.strip()}')
                    else:
                        fails.append(f'{response.reason}: {soup}')
                else:
                    if s.then:

                        for a in s.then:
                            result = a.action(soup)
                            if result:
                                fails.append(result)
                    else:
                        fails.append("no expectations provided")

                if not fails:
                    print(f'    {CGREEN}✔︎ when {s.when.text}{CEND}')
                    scenario_ok+=1
                else:
                    print(f'    {CRED}✗ when {s.when.text}{CEND}:')
                    scenario_fail+=1
                    scenario_fail_count += len(fails)
                    for fail in fails:
                        print(f'      {fail}')

                module_fails += len(fails)
        except:
            print(f'    {CRED}✗ {name}{CEND}:')
            lines = traceback.format_exception(*sys.exc_info(), limit=None, chain=True)
            print(f'      {CRED}{"      ".join(lines).rstrip()}{CEND}')
            print(f'      {os.path.join(os.path.abspath(root), l) }')
    if count == 0:
        print(f'  {CRED}✗ feature has no specs{CEND}')

    if count == 0 or module_fails > 0:
        print(f'  {os.path.abspath(os.path.join(root, l))}')
        file_fails += 1
    else:
        file_ok += 1

if file_fails:
    print(f'{CRED}✗ File FAILS {file_fails}{CEND}')
if file_ok:
    print(f'{CGREEN}✔ File OK {file_ok}{CEND}')

print(f'Scenarios: {scenario_ok} OK, {scenario_fail} fail (fix {scenario_fail_count})')
