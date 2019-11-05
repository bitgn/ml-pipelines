from typing import Optional, List, Any

from .cl_bases import *
from . import vo_pb2 as vo
import time


class Dataset:
    def __init__(self, ctx: Context, project_id:str, dataset_id:str):
        self.project_id  = project_id
        self.dataset_id = dataset_id
        self.ctx = ctx


    def add_activity(self, multiline_text:str, timestamp_sec:int = 0):





        req=api.AddDatasetActivityRequest(
            project_id=self.project_id,
            dataset_id=self.dataset_id,
            multiline_text=multiline_text,
            level=vo.ACTIVITY_INFO,
            timestamp=timestamp_sec,
        )

        self.ctx.add_dataset_activity(req)

