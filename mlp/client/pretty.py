import datetime


def sizeof_fmt(num, suffix='B'):
    for unit in ['','Ki','Mi','Gi','Ti','Pi','Ei','Zi']:
        if abs(num) < 1024.0:
            return "%3.1f%s%s" % (num, unit, suffix)
        num /= 1024.0
    return "%.1f%s%s" % (num, 'Yi', suffix)


def date_fmt(d):
    diff = datetime.datetime.utcnow() - d
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
        return '{} seconds ago'.format(s)
    elif s < 120:
        return 'a minute ago'
    elif s < 3600:
        return '{} minutes ago'.format(s/60)
    elif s < 7200:
        return 'an hour ago'
    else:
        return '{0:.0f} hours ago'.format(s/3600)