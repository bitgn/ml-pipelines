import agg
from proto import events_pb2 as events

import lmdb
import utils

from . import projection
from . import db



class Service:
    """
    Aggregate for the domain logic. At this point we defer:

    1. Any persistent storage (to be replaced with in-mem views later)
    2. Client (agg to be invoked via python repl)

    """

    def __init__(self, env:lmdb.Environment):
        self.env = env

    def apply(self, tx: lmdb.Transaction, e):
        projection.apply(e, tx)
        db.events_append(tx, e)
        print(f'{type(e).__name__}\n{e}')

    def add_project(self, project_id: str, name: str):
        with self.env.begin(write=True) as tx:
            if db.project_get(tx, project_id):
                raise KeyError(f'Project {project_id} already exists')

            self.apply(tx, events.ProjectCreated(project_id=project_id, name=name))

    def list_projects(self):
        with self.env.begin() as tx:

            return list(db.project_list(tx))


    def add_dataset(self, project_id: str, name: str):
        with self.env.begin(write=True) as tx:
            project = db.project_get(tx, project_id)

            if not project:
                raise KeyError(f'Project {project_id} not found')

            dataset_id = str(utils.sequential_uuid())

            e = events.DatasetCreated(dataset_id=dataset_id, project_id=project_id, name=name)
            self.apply(tx, e)
            return dataset_id


