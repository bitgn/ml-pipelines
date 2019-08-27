from env import *


def given_empty_system(e: Env):
    """create project if all arguments are valid"""

    e.scenario(
        when.client(lambda c: c.create_project(name="ds!!")),
        then.bad_name("ds!!")
    )

    e.scenario(
        when.client(lambda c: c.create_project(name="some")),
        then.client_ok()
    )


def given_existing_project(e: Env):
    """Can't create duplicates"""
    prj = preset.project_created(e)
    e.given_events(prj)
    e.scenario(
        when.client(lambda c: c.create_project(name=prj.name)),
        then.name_taken(prj.uid, prj.name)
    )


