"""Custom context processor for the Django templates"""

from django.conf import settings  # import the settings file
from explore.config import env
from agg import db


def base_template(request):
    """inject common variables for the base template to use"""
    with env.begin() as tx:
        stats = db.stats_get(tx)
    # return the value as a dictionary
    return {
        'APP_VERSION': settings.APP_VERSION,
        'PROJECT_COUNT': stats.project_count,
        'DATASET_COUNT': stats.dataset_count,
        'PIPELINE_COUNT': stats.pipeline_count,
        'MENU_ACTIVE': 'none',
        'EXPERT_COUNT': stats.expert_count,
        'DASHBOARD_COUNT': stats.dashboard_count,
        'MODEL_COUNT': stats.model_count,
    }
