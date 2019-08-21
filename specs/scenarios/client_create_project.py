from env import *
import test_api as api


def given_empty_system(e: Env):
    """display zero counts for all entities"""


    msg = api.CreateProjectRequest(ProjectId='ds   ', ProjectName='none')
    e.scenario(
        when.create_project(msg)
    )

