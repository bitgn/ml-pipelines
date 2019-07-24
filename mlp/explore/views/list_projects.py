from django.shortcuts import render
from explore.config import env
from agg import db


def list_projects(request):


    with env.begin() as tx:
        projects = list(db.project_list(tx))



    return render(request, "explore/projects.html", context={
        'projects': projects,
        'MENU_ACTIVE': 'projects'
    })



