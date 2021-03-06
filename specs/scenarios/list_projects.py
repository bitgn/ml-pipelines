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
    prj = preset.project_created(t)

    t.given_events(prj)

    href = urls.view_project(prj.name)
    t.scenario(
        when.list_projects(),
        then.none('main #empty-catalog-message'),
        then.count('main .project-info', 1),

        then.link(f'main #prj-{prj.uid.hex()} .project-link', href=href, text=prj.meta.title),
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
