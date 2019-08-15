"""MLP-4"""

from env import *

def given_a_dataset(t: Env):
    """find it via project name, dataset name or sample"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)

    prj.name = "PROJECT_NAME"
    ds.name = "DATASET_NAME"

    preset.set_sample(t, ds.meta, """{SAMPLE_FIELD: "SAMPLE_VALUE"}""")
    preset.set_description(t, ds.meta, "readme: description_text is here")
    preset.set_data_format(t, ds.meta, "CUSTOM_FORMAT")

    ds.meta.location_uri = "LOCATION_URI"
    ds.meta.location_uri_set = True

    ds.meta.location_id = "LOC_ID"
    ds.meta.location_id_set = True

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
        "LOCATION_URI",
        "LOC_ID",
        # multi-phrase search
        "PROJECT_N DATASET_N",
        "PROJECT DATASET SAMPLE",

    ]

    for q in valid_queries:
        t.scenario(when.search_datasets(q), then.exists(f'main #ds-{ds.dataset_id}'))

    for q in valid_queries:
        invalid = q + "_ZZZ"
        t.scenario(when.search_datasets(invalid), then.none(f'main #ds-{ds.dataset_id}'))
