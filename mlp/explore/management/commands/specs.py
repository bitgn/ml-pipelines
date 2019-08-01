import importlib
import os
import pathlib
import sys
from inspect import getmembers, isfunction, getdoc

import bs4
from django.core.management import BaseCommand
from django.test import Client
import test
import lmdb
from django.conf import settings

from agg import projection
from explore import config

import http

CBLACK = '\33[30m'
CRED = '\33[31m'
CGREEN = '\33[32m'
CYELLOW = '\33[33m'
CBLUE = '\33[34m'
CVIOLET = '\33[35m'
CBEIGE = '\33[36m'
CWHITE = '\33[37m'

CEND = '\33[0m'
CBOLD = '\33[1m'
CITALIC = '\33[3m'
CURL = '\33[4m'
CBLINK = '\33[5m'
CBLINK2 = '\33[6m'
CSELECTED = '\33[7m'

import logging


class Command(BaseCommand):
    help = 'Executes event-driven specs'

    def add_arguments(self, parser):
        parser.add_argument("--story", action="store_true")

    def handle(self, *args, **options):

        logging.disable(logging.CRITICAL)

        root = 'explore/specs'
        sys.path.append(root)

        client = Client()

        with lmdb.open(settings.LMDB, subdir=False) as env:

            main = env.open_db()

            file_fails = 0;
            file_ok = 0;
            assertion_fails = 0;

            for l in os.listdir(root):
                stem = pathlib.Path(l).stem

                if stem.startswith('_') or stem in ['preset', 'when']:
                    continue


                module = importlib.import_module(stem)

                doc = getdoc(module)

                print(f'\n{CBOLD}{stem}{CEND}: {doc or ""}')

                count = 0
                module_fails = 0

                for name, factory in getmembers(module, isfunction):
                    if not name.startswith('given_'):
                        continue

                    count += 1

                    test_env = test.Env()
                    factory(test_env)

                    def utcnow():
                        return test_env.time

                    config.set_utcnow(utcnow)

                    with env.begin(write=True) as tx:
                        tx.drop(main)

                        for e in test_env.events:
                            projection.apply(e, tx)

                    print(f'  {CYELLOW}{name}{CEND} - {factory.__doc__}')

                    for s in test_env.scenarios:
                        fails = []

                        response = s.when.action(client)

                        soup = bs4.BeautifulSoup(response.content, 'lxml')
                        if response.status_code != http.HTTPStatus.OK:
                            if soup.title:
                                fails.append(f'{response.reason_phrase}: {soup.title.text}')
                            else:
                                fails.append(f'{response.reason_phrase}: {soup}')
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
                        else:
                            print(f'    {CRED}✗ when {s.when.text}{CEND}:')
                            for fail in fails:
                                print(f'      {fail}')


                        module_fails += len(fails)

                if count == 0:
                    print(f'  {CRED}✗ feature has no specs{CEND}')

                if count == 0 or module_fails > 0:
                    print(f'  {os.path.abspath(os.path.join(root, l))}')
                    file_fails+=1
                else:
                    file_ok+=1
        print()
        if file_fails:
            print(f'{CRED}✗ File FAILS {file_fails}{CEND}')
        if file_ok:
            print(f'{CGREEN}✔ File OK {file_ok}{CEND}')