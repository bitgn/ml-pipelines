from env import *


def given_empty_system(e: Env):
    """throw an error"""

    e.scenario(
        when.client(lambda c: c.create_dataset(name="test",project_name='non-existent')),
        then.not_found(subject_name='non-existent')
    )

def given_a_project(e: Env):
    "Create a new dataset"
    prj = preset.project_created(e)

    e.given_events(prj)

    e.scenario(
        when.client(lambda c: c.create_dataset(name="test", project_name=prj.name)),
        then.client_ok(uid=b'12')
    )

