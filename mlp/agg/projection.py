from functools import singledispatch

from proto import events_pb2 as evt
from proto import dto_pb2 as dto
import lmdb

from agg import db


@singledispatch
def apply(e, tx: lmdb.Transaction):
    print(f'Unregistered event handler {type(e)}')


@apply.register
def _(e: evt.ProjectCreated, tx: lmdb.Transaction):
    db.project_add(tx, dto.ProjectData(id=e.project_id, name=e.name, dataset_count=0))
    stats = db.stats_get(tx)
    stats.project_count += 1
    db.stats_put(tx, stats)


def _apply_metadata(s: evt.DatasetMetadata, t: dto.DatasetData):
    for f in s.set_fields:

        if f == evt.FIELD_DESCRIPTION:
            t.description = s.description
            t.description_set = True
            continue

        if f == evt.FIELD_DATA_FORMAT:
            t.data_format = s.data_format
            t.data_format_set = True
            continue

        if f == evt.FIELD_SAMPLE:
            t.sample_kind = s.sample_kind
            t.sample_body = s.sample_body
            t.sample_set = True
            continue
        if f == evt.FIELD_FILE_COUNT:
            t.file_count = s.file_count
            t.file_count_set = True
            continue
        if f == evt.FIELD_RAW_BYTES:
            t.raw_bytes = s.raw_bytes
            t.raw_bytes_set = True
            continue

        if f == evt.FIELD_ZIP_BYTES:
            t.zip_bytes = s.zip_bytes
            t.zip_bytes_set = True
            continue

        if f == evt.FIELD_RECORD_COUNT:
            t.record_count = s.record_count
            t.record_count_set = True
            continue

        if f == evt.FIELD_UPDATE_TIMESTAMP:
            t.update_timestamp = s.update_timestamp
            t.update_timestamp_set = True
            continue

    for f in s.del_fields:
        print("DEL!")
        if f == evt.FIELD_DESCRIPTION:
            t.description = None
            t.description_set = False
            continue

        if f == evt.FIELD_DATA_FORMAT:
            t.data_format = None
            t.data_format_set = False
            continue

        if f == evt.FIELD_SAMPLE:
            t.sample_kind = None
            t.sample_body = None
            t.sample_set = False
            continue
        if f == evt.FIELD_FILE_COUNT:
            t.file_count = None
            t.file_count_set = False
            continue
        if f == evt.FIELD_RAW_BYTES:
            t.raw_bytes = None
            t.raw_bytes_set = False
            continue

        if f == evt.FIELD_ZIP_BYTES:
            t.zip_bytes = None
            t.zip_bytes_set = False
            continue

        if f == evt.FIELD_RECORD_COUNT:
            t.record_count = None
            t.record_count_set = False
            continue

        if f == evt.FIELD_UPDATE_TIMESTAMP:
            t.update_timestamp = None
            t.update_timestamp_set = False
            continue


@apply.register
def _(e: evt.DatasetCreated, tx: lmdb.Transaction):
    val = dto.DatasetData(
        project_id=e.project_id,
        dataset_id=e.dataset_id,
        name=e.name,
    )

    _apply_metadata(e.metadata, val)

    db.dataset_add(tx, val)
    prj = db.project_get(tx, project_id=e.project_id)
    prj.dataset_count += 1
    prj.zip_bytes += e.metadata.raw_bytes
    prj.raw_bytes += e.metadata.raw_bytes
    db.project_add(tx, prj)

    stats = db.stats_get(tx)
    stats.dataset_count += 1
    db.stats_put(tx, stats)


@apply.register
def _(e: evt.DatasetUpdated, tx: lmdb.Transaction):
    val = db.dataset_get(tx, e.dataset_id)

    _apply_metadata(e.metadata, val)

    db.dataset_add(tx, val)


@apply.register
def _(e: evt.JobAdded, tx: lmdb.Transaction):
    val = dto.Job(
        job_id=e.job_id,
        job_name=e.job_name,
        inputs=e.inputs,
        outputs=e.outputs,
        project_id=e.project_id
    )

    stats = db.stats_get(tx)
    stats.job_count += 1
    db.stats_put(tx, stats)

    db.job_put(tx, val)

    # we are doing it the slow way for now

    for input_id in e.inputs:
        ds = db.dataset_get(tx, input_id)
        ds.downstream_jobs.append(e.job_id)
        db.dataset_add(tx, ds)

    for output_id in e.outputs:
        ds = db.dataset_get(tx, output_id)
        ds.upstream_jobs.append(e.job_id)
        db.dataset_add(tx, ds)

    prj = db.project_get(tx, project_id=e.project_id)
    prj.job_count += 1
    db.project_add(tx, prj)
