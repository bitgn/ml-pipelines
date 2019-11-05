from typing import Optional, List, Any

from .cl_bases import *
from . import vo_pb2 as vo
import time


class Dataset:
    def __init__(self, ctx: Context, project_id:str, dataset_id:str):
        self.project_id  = project_id
        self.dataset_id = dataset_id
        self.ctx = ctx


    def add_error(self, multiline_text:str, timestamp_sec:int = 0):
        return self.add_activity(multiline_text=multiline_text, timestamp_sec=timestamp_sec, level="error")

    def add_info(self, multiline_text:str, timestamp_sec:int = 0):
        return self.add_activity(multiline_text=multiline_text, timestamp_sec=timestamp_sec, level="info")
    def add_success(self, multiline_text:str, timestamp_sec:int = 0):
        return self.add_activity(multiline_text=multiline_text, timestamp_sec=timestamp_sec, level="success")


    def add_activity(self, multiline_text:str, timestamp_sec:int = 0, level:Optional[str]=None):


        map = {
            'info': vo.ACTIVITY_INFO,
            'error': vo.ACTIVITY_ERROR,
            'success':vo.ACTIVITY_SUCCESS,
            'verbose':vo.ACTIVITY_VERBOSE,
        }



        req=api.AddDatasetActivityRequest(
            project_id=self.project_id,
            dataset_id=self.dataset_id,
            multiline_text=multiline_text,
            level=map.get(level, vo.ACTIVITY_INFO),
            timestamp=timestamp_sec,
        )

        self.ctx.add_dataset_activity(req)

