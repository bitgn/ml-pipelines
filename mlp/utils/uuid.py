
from random import getrandbits

from uuid import uuid1, getnode

_my_clock_seq = getrandbits(14)
_my_node = getnode()


def sequential(node=None):
    return uuid1(node=node, clock_seq=_my_clock_seq)