"""MLP-?????"""
from env import *


def given_no_projects(e: Env):
    """show an empty placeholder so that the user could create projects"""
    e.scenario(
        when.list_projects(),
        then.exists('main #empty-catalog-message')
    )


def given_a_project(t: Env):
    """show a project overview with a link to a project"""
    e = preset.project_created(t)

    t.given_events(e)

    href = urls.view_project(e.project_id)
    t.scenario(
        when.list_projects(),
        then.none('main #empty-catalog-message'),
        then.count('main .project-info', 1),
        then.text(f'main #prj-{e.project_id} .project-name', e.name),
        then.text(f'main #prj-{e.project_id} .dataset-count', '0'),
        then.link(f'main #prj-{e.project_id} .project-link', href=href, text=e.name),
    )


def given_a_few_projects(t: Env):
    """display all projects on the page"""
    count = 3

    for i in range(count):
        prj = preset.project_created(t)
        t.given_events(prj)

    t.scenario(
        when.list_projects(),
        then.count('main .project-info', count)
    )
