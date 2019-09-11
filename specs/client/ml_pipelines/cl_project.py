from typing import Optional

from .cl_bases import *
from . import vo_pb2 as vo

from .cl_job import Job
from .cl_dataset import Dataset
from .cl_system import System

class MultiCommit:
    def __init__(self, ctx, project_uid, clean_slate):
        """

        :type clean_slate: bool
        :type project_uid: bytes
        :type ctx: Context
        """
        self.clean_slate = clean_slate
        self.project_uid = project_uid
        self.ctx = ctx


class Systems:
    def __init__(self, ctx: Context, uid: bytes):
        self.uid = uid
        self.ctx = ctx


    def get_or_add_report(self, name: str) -> System:
        return self._get_or_add(name, kind=vo.SystemKind.Report)


    def add_report(self, name:str, title: Optional[str] = None, location_uri: Optional[str] = None) -> System:
        return self._add(name, vo.SystemKind.Report, title, location_uri=location_uri)


    def add_db(self, name:str, title: Optional[str] = None, location_uri: Optional[str] = None) -> System:
        return self._add(name, vo.SystemKind.Database, title, location_uri=location_uri)

    def add_table(self, name:str, title: Optional[str] = None, location_uri: Optional[str] = None) -> System:
        return self._add(name, vo.SystemKind.Table, title, location_uri=location_uri)

    def add_topic(self, name:str, title: Optional[str] = None, location_uri: Optional[str] = None) -> System:
        return self._add(name, vo.SystemKind.Topic, title, location_uri=location_uri)


    def add_service(self, name:str, title: Optional[str] = None, location_uri: Optional[str] = None) -> System:
        return self._add(name, vo.SystemKind.Service, title, location_uri=location_uri)



    def _add(self, name: str, kind: vo.SystemKind, title: Optional[str] = None, location_uri: Optional[str] = None) -> System:
        meta = vo.SystemMetadataDelta()

        meta.kind = kind
        meta.kind_set = True

        if title:
            meta.title = title
            meta.title_set = True

        if location_uri:
            meta.location_uri = location_uri
            meta.location_uri_set = True



        new = api.AddSystemRequest(
            project_uid=self.uid,
            name=name,
            meta=meta,
        )
        resp = self.ctx.add_system(new)
        return System(
            self.ctx,
            project_uid=self.uid,
            uid=resp.uid,
            name=name,
        )
    def get(self, name:str) -> System:

        req = api.GetSystemRequest(
            project_uid=self.uid,
            name=name
        )
        resp = self.ctx.get_system(req)

        return System(self.ctx, project_uid=self.uid, uid=resp.uid, name=name)

    def _get_or_add(self, name: str, kind: vo.SystemKind) -> System:
        try:
            return self.get(name)
        except errors.NotFound:
            pass

        return self._add(name, kind)


class Project:


    def __init__(self, ctx: Context, uid: bytes):
        self.uid = uid
        self.ctx = ctx
        self.systems = Systems(ctx, uid)


    def prepare_commit(self, clean_slate:bool = False):
        return MultiCommit(self.ctx, self.uid, clean_slate)


    def create_job(self, name:str, title:Optional[str]=None) -> Job:
        meta = vo.JobMetadataDelta()
        if title:
            meta.title = title
            meta.title_set = True

        create = api.CreateJobRequest(project_uid=self.uid, name=name,meta=meta)


        resp = self.ctx.create_job(create)
        return Job(self.ctx, self.uid, resp.uid, resp.name)

    def get_job(self, name:str) -> Job:
        get = api.GetJobRequest(project_uid=self.uid, name=name)
        resp = self.ctx.get_job(get)

        return Job(self.ctx, self.uid, resp.uid, resp.name)

    def get_or_create_job(self, name: str) -> Job:
        try:
            return self.get_job(name)
        except errors.NotFound:
            pass

        return self.create_job(name)


    def get_dataset(self, name) -> Dataset:

        get = api.GetDatasetRequest(project_uid=self.uid, name=name)
        resp = self.ctx.get_dataset(get)



        return Dataset(self.ctx, self.uid, resp.uid, resp.name,
                       location_id=resp.location_id)

    def add_dataset(self, name,
                    location_id: Optional[str] = None,
                    title: Optional[str] = None, data_format: Optional[str] = None,
                    description: Optional[str] = None,
                    location_uri: Optional[str] = None) -> Dataset:
        meta = vo.DatasetMetadataDelta(

        )
        if title:
            meta.title_set=True
            meta.title=title


        if data_format:
            meta.data_format=data_format
            meta.data_format_set=True

        if description:
            meta.description=description
            meta.description_set=True

        if location_uri:
            meta.location_uri = location_uri
            meta.location_uri_set = True





        new = api.CreateDatasetRequest(
            project_uid=self.uid,
            name=name,
            meta=meta
        )

        resp = self.ctx.create_dataset(new)
        return Dataset(
            self.ctx,
            project_uid=self.uid,
            uid=resp.uid,
            name=name,
            location_id=location_id,
        )

    def get_or_create_dataset(self, name: str, location_id: Optional[str]=None) -> Dataset:
        try:
            return self.get_dataset(name)
        except errors.NotFound:
            pass

        return self.add_dataset(name, location_id=location_id)


