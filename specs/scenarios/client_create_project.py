from env import *
from test_api import api_pb2 as api


def given_empty_system(e: Env):
    """display zero counts for all entities"""


    msg = api.CreateProjectRequest(ProjectId='ds   ', ProjectName='none')
    e.scenario(
        when.create_project(msg)
    )

