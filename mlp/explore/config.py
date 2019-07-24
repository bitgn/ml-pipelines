import lmdb
from django.conf import settings

env = lmdb.open(settings.LMDB, subdir=False)

import datetime

utcnow = datetime.datetime.utcnow

def set_utcnow(f):
    global utcnow
    utcnow = f