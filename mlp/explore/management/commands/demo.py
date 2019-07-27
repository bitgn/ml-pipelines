"""
Custom Django command to fill detabase with a demo data
"""

import lmdb
from django.conf import settings
from django.core.management import BaseCommand
from agg import projection
from explore import demo
import random

random.seed(1)


class Command(BaseCommand):
    help = 'Fill a demo db'

    def add_arguments(self, parser):
        parser.add_argument("--setup", default="random")
        pass

    def handle(self, *args, **options):
        with lmdb.open(settings.LMDB, subdir=False) as env:
            main = env.open_db()

            with env.begin(write=True) as tx:
                tx.drop(main)

                if options['setup'] == "empty":
                    return

                if options['setup'] == "analytics":
                    for e in demo.setup_analytics_demo():
                        projection.apply(e, tx)
                    return

                for e in demo.setup_analytics_demo():
                    projection.apply(e, tx)

                for e in demo.setup_random_demo():
                    projection.apply(e, tx)
