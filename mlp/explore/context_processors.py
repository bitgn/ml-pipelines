from django.conf import settings  # import the settings file
from explore.config import env
from agg import db


def base_template(request):
    with env.begin() as tx:
        stats = db.stats_get(tx)
    # return the value you want as a dictionary. you may add multiple values in there.
    return {
        'APP_VERSION': settings.APP_VERSION,
        'PROJECT_COUNT': stats.project_count,
        'DATASET_COUNT': stats.dataset_count,
        'PIPELINE_COUNT': stats.pipeline_count,
        'MENU_ACTIVE':'none'
    }
