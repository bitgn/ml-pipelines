from env import *


def given_empty_system(e: Env):
    """create project if all arguments are valid"""

    e.scenario(
        when.client_create_project(project_id='ds!!', project_name="some"),
        then.invalid_argument(subject_id='project_id')
    )

    e.scenario(
        when.client_create_project(project_id='ds', project_name="some"),
        then.client_ok()
    )


def given_existing_project(e: Env):
    """Can't create duplicates"""
    prj = preset.project_created(e)
    e.given_events(prj)
    e.scenario(
        when.client_create_project(project_id=prj.project_id),
        then.already_exists(subject_id=prj.project_id)
    )
