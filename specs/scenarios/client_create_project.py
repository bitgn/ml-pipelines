from env import *


def given_empty_system(e: Env):
    """display zero counts for all entities"""

    e.scenario(
        when.create_project(project_id='ds', project_name="some")
    )

