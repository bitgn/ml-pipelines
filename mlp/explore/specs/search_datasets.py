"""MLP-4"""

from explore.specs import when, preset
from test import *


def given_a_dataset(t: test.Env):
    """find it via project name, dataset name or sample"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)

    prj.name = "PROJECT_NAME"
    ds.name = "DATASET_NAME"

    preset.set_sample(t, ds.metadata, """{SAMPLE_FIELD: "SAMPLE_VALUE"}""")
    preset.set_description(t, ds.metadata, "readme: description_text is here")
    preset.set_data_format(t, ds.metadata, "CUSTOM_FORMAT")

    t.given_events(prj, ds)

    invalid_queries = [
        "no-match",
        "PROJECT_NAME no-match"
    ]

    for q in invalid_queries:
        t.scenario(when.search_datasets(q), then.none('main .dataset-info'))

    valid_queries = [
        "PROJECT_NAME",
        "DATASET_NAME",
        "project_name",
        "T_NAME",
        "SAMPLE_FIELD",
        "SAMPLE_VALUE",
        "DESCRIPTION_TEXT",
        "CUSTOM_FORMAT",
        "PROJECT_N DATASET_N",
        "PROJECT DATASET SAMPLE"
    ]

    for q in valid_queries:
        t.scenario(when.search_datasets(q), then.exists(f'main #ds-{ds.dataset_id}'))
