from django import template
import datetime as dt
from utils import pretty


register = template.Library()

from explore import config


@register.filter(name="fmt_timestamp")
def timestamp_since_filter(ts):
    if not ts:
        return ''

    time = dt.datetime.fromtimestamp(ts)
    now = config.utcnow()
    delta = now - time

    return pretty.timedelta(delta)

@register.filter(name="fmt_bytes")
def bytes_filter(bytes):
    return pretty.bytes(bytes)