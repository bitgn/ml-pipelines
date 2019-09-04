from typing import Optional, List, Any

from .cl_bases import *
from . import vo_pb2 as vo



class DatasetStaging:
    files: List[vo.DatasetItem]

    def __init__(self, ctx: Context, project_uid: bytes, dataset_uid: bytes, version_uid: bytes, clean_slate: bool, dataset_name: str):
        self.dataset_name = dataset_name
        self.version_uid = version_uid
        self.dataset_uid = dataset_uid
        self.project_uid = project_uid
        self.clean_slate = clean_slate
        self.ctx = ctx

        self.files = []
        self.inputs = []

    def add_file(self, name: str, records: int, size: int):
        self.files.append(vo.DatasetItem(name=name, storage_bytes=size, records=records))

    def add_input(self, run: PipelineInput):

        kind = run.__class__.__name__

        if kind == 'JobRun':
            i = vo.DatasetVerInput(uid=run.uid, type=vo.DatasetVerInput.TYPE.JOB_RUN)
            self.inputs.append(i)
            return
        raise ValueError(f'Unexpected dependency: {kind}')

    def commit(self, title: str) -> 'DatasetVersion':

        change_set = api.DatasetChangeSet(
            dataset_uid=self.dataset_uid,
            parent_version_uid=self.version_uid,
            clean_slate=self.clean_slate,
        )

        change_set.inputs.extend(self.inputs)
        change_set.items.extend(self.files)

        r = api.CommitRequest(
            title=title,
        )

        r.datasets.append(change_set)




        resp = self.ctx.commit(r)
        return DatasetVersion(
            self.ctx,
            project_uid=self.project_uid,
            dataset_uid=self.dataset_uid,
            dataset_name=self.dataset_name,
            uid = resp.dataset_versions[0]
        )


class DatasetVersion:
    def __init__(self, ctx: Context, project_uid, dataset_uid, dataset_name, uid):
        self.uid = uid
        self.dataset_name = dataset_name
        self.dataset_uid = dataset_uid
        self.project_uid = project_uid

        self.ctx = ctx

    def headless(self) -> bool:
        return self.uid is None

    def prepare_commit(self, clean_slate: bool = False) -> DatasetStaging:
        return DatasetStaging(
            self.ctx,
            self.project_uid,
            self.dataset_uid,
            self.uid,
            clean_slate,
            self.dataset_name
        )


class Dataset:
    def __init__(self, ctx: Context, project_uid, uid, name, location_id: str):
        self.location_id = location_id
        self.name = name
        self.uid = uid
        self.project_uid = project_uid
        self.ctx = ctx

    def get_last_version(self) -> DatasetVersion:

        try:

            req = api.GetLastDatasetVersionRequest(project_uid=self.project_uid, dataset_uid=self.uid)
            resp = self.ctx.get_last_version(req)

            return DatasetVersion(self.ctx, self.project_uid, self.uid, self.name, resp.uid)
        except errors.NotFound:

            return DatasetVersion(self.ctx, self.project_uid, self.uid, self.name, None)
