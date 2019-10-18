from typing import Optional, List

from .cl_bases import *
from . import vo_pb2 as vo

from .cl_dataset import Dataset


class Project:


    def __init__(self, ctx: Context, project_id: str):
        self.project_id = project_id
        self.ctx = ctx




    def get_dataset(self, id:str) -> Dataset:

        get = api.GetDatasetRequest(project_id=self.project_id, dataset_id=id)
        resp = self.ctx.get_dataset(get)



        return Dataset(self.ctx, project_id=self.project_id, dataset_id=resp.dataset_id)

    def list_datasets(self) -> List[Dataset]:
        get = api.ListDatasetsRequest(project_id=self.project_id)
        resp = self.ctx.list_datasets(get)

        res = []

        d: api.DatasetInfoResponse
        for d in resp.datasets:
            res.append(Dataset(self.ctx, project_id=d.project_id, dataset_id=d.dataset_id))

        return res

    def add_dataset(self, dataset_id:str,
                    description: Optional[str] = None,
                    summary: Optional[str] = None) -> Dataset:
        meta = vo.DatasetMetadataDelta(

        )
        if description:
            meta.description_set = True
            meta.description = description

        if summary:
            meta.summary_set = True
            meta.summary = summary

        new = api.CreateDatasetRequest(
            project_id=self.project_id,
            dataset_id=dataset_id,
            meta=meta
        )

        resp = self.ctx.create_dataset(new)
        return Dataset(
            self.ctx,
            project_id=self.project_id,
            dataset_id=resp.dataset_id,

        )

    def get_or_create_dataset(self, dataset_id: str) -> Dataset:
        try:
            return self.get_dataset(dataset_id)
        except errors.NotFound:
            pass

        return self.add_dataset(dataset_id)


