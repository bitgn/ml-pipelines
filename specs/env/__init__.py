
from .env import Env

from .env import When

from .env import Then

from .when import *
from .then import *

from . import urls
from . import preset

from api import events_pb2 as evt



def nbsp(value):
    """
    Avoid text wrapping in the middle of a phrase by adding non-breaking
    spaces where there previously were normal spaces.
    """
    return value.replace(" ", "\xa0")
