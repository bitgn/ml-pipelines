"""
Custom Django command to wipe the detabase
"""

import lmdb
from django.conf import settings
from django.core.management import BaseCommand
import random

random.seed(1)


class Command(BaseCommand):
    help = 'Drop the database'

    def handle(self, *args, **options):
        with lmdb.open(settings.LMDB, subdir=False) as env:
            main = env.open_db()

            with env.begin(write=True) as tx:
                tx.drop(main)

