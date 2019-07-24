import datetime


def _avoid_wrapping(value):
    """
    Avoid text wrapping in the middle of a phrase by adding non-breaking
    spaces where there previously were normal spaces.
    """
    return value.replace(" ", "\xa0")



def bytes(num):

    if num == 1:
        return "1 byte"
    for unit in ['bytes','KB','MB','GB','TB','PB','EB','ZB']:
        if abs(num) < 1024.0:
            return _avoid_wrapping(f"{num:3.1f} {unit}")
        num /= 1024.0
    return _avoid_wrapping(f"{num:.1f}Y")


def timedelta(diff: datetime.timedelta):

    s = diff.seconds
    if diff.days > (365 * 2):
        return f'{diff.days / 365} years ago'

    if diff.days > 365:
        return 'a year ago'

    if diff.days > 60:
        return f'{diff.days / 30} months ago'

    if diff.days > 30:
        return 'a month ago'

    if diff.days > 14:
        return f'{diff.days / 7} weeks ago'

    if diff.days > 7:
        return 'a week ago'

    if diff.days > 1:
        return f'{diff.days} days ago'


    if diff.days == 1:
        return 'a day ago'
    elif diff.days > 1:
        return '{} days ago'.format(diff.days)
    elif s <= 1:
        return 'just now'
    elif s < 60:
        return '{0:.0f} seconds ago'.format(s)
    elif s < 120:
        return 'a minute ago'
    elif s < 3600:
        return '{0:.0f} minutes ago'.format(s/60)
    elif s < 7200:
        return 'an hour ago'
    else:
        return '{0:.0f} hours ago'.format(s/3600)