"""
As a user I want to be able to list and filter datasets
so that I could discover and reuse them.
"""

from django.shortcuts import render

from explore.config import env
from agg import db
from explore.domain import stale


def explore_datasets(request):
    query = request.GET.get('query', '')

    query = query.strip(' \t\n\r')

    projects = {}
    values = []

    with env.begin() as tx:
        datasets = list(db.dataset_list_all(tx))

        for ds in datasets:

            if ds.project_id not in projects:
                projects[ds.project_id] = db.project_get(tx, ds.project_id)

            project = projects[ds.project_id]
            is_stale = stale.is_stale(ds)
            meta = {
                'project_name': project.name,
                'stale': is_stale
            }

            if not query:
                values.append((ds, meta))
                continue


            fts = f'{ds.name}|{project.name}|{ds.data_format}|{ds.location_id}|{ds.location_uri}'

            if ds.sample_set:
                fts += "|" + ds.sample_body.decode()

            if ds.description_set:
                fts += "|" + ds.description

            if is_stale:
                fts += "|stale"

            fts = fts.lower()

            matches = True
            for part in query.lower().split(' '):
                part = part.strip(' ')
                if not part:
                    continue
                if not part in fts:
                    matches = False
                    break


            if matches:
                values.append((ds, meta))

    catalog_is_empty = not query and not datasets

    context = {
        'query': query,
        'd': {},
        'datasets': values,
        'catalog_is_empty': catalog_is_empty,
        'MENU_ACTIVE':'datasets'
    }
    return render(request, "explore/explore-datasets.html", context)
