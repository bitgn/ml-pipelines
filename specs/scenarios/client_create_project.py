from env import *


def given_empty_system(e: Env):
    """create project if all arguments are valid"""

    e.scenario(
        when.client(lambda c: c.add_project(name="ds!!")),
        then.bad_name("ds!!")
    )

    e.scenario(
        when.client(lambda c: c.add_project(name="some")),
        then.client_ok()
    )


def given_existing_project(e: Env):
    """Can't create duplicates"""
    prj = preset.project_created(e)
    e.given_events(prj)
    e.scenario(
        when.client(lambda c: c.add_project(name=prj.name)),
        then.already_exists(prj.uid, prj.name)
    )


