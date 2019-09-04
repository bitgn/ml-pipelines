from typing import Optional

from .cl_bases import *
from . import vo_pb2 as vo

from .cl_job import Job
from .cl_dataset import Dataset

class Project:


    def __init__(self, ctx: Context, uid: bytes):
        self.uid = uid
        self.ctx = ctx




    def get_or_create_job(self, name: str) -> Job:

        try:

            get = api.GetJobRequest(project_uid=self.uid, name=name)
            resp = self.ctx.get_job(get)

            return Job(self.ctx, self.uid, resp.uid, resp.name)
        except errors.NotFound:
            pass

        create = api.CreateJobRequest(project_uid=self.uid, name=name)
        resp = self.ctx.create_job(create)
        return Job(self.ctx, self.uid, resp.uid, resp.name)



    def get_dataset(self, name) -> Dataset:

        get = api.GetDatasetRequest(project_uid=self.uid, name=name)
        resp = self.ctx.get_dataset(get)



        return Dataset(self.ctx, self.uid, resp.uid, resp.name,
                       location_id=resp.location_id)

    def get_or_create_dataset(self, name: str, location_id: str) -> Dataset:

        try:
            return self.get_dataset(name)

        except errors.NotFound:
            pass

        new = api.CreateDatasetRequest(
            project_uid=self.uid,
            name=name,
            meta=vo.DatasetMetadataDelta(
                location_id=location_id
            )
        )

        resp = self.ctx.create_dataset(new)
        return Dataset(
            self.ctx,
            project_uid=self.uid,
            uid=resp.uid,
            name=name,
            location_id=location_id
        )
