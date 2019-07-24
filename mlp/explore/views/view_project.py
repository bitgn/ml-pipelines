from typing import List, Tuple, Dict

from django.shortcuts import render

from explore.config import env
from explore.domain import stale
from agg import db

from proto import dto_pb2 as dto


def view_project(request, project_id):
    with env.begin() as tx:
        prj = db.project_get(tx, project_id)
        if not prj:
            return "Project not found"

        meta : List[Tuple[dto.DatasetData,Dict]] = []

        for ds in db.dataset_list(tx, project_id):
            val = {'stale': stale.is_stale(ds)}
            meta.append((ds, val))

    return render(request, "explore/project.html", context={
        'datasets': meta,
        'name': prj.name,
    })

