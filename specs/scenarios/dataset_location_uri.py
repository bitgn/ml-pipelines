from env import *



def given_a_dataset_with_storage_location(t: Env):
    """show storage location on all related screens"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)

    ds.meta.location_uri = "url://aws"
    ds.meta.location_uri_set =True

    t.given_events(prj, ds)

    t.scenario(
        when.view_dataset(ds.dataset_id),
        then.text(f'main #ds-{ds.dataset_id} .location-uri', "url://aws"),
    )


def given_a_dataset_without_storage_location(t: Env):
    """show N/A instead"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)

    ds.meta.location_uri = ''
    ds.meta.location_uri_set = True

    t.given_events(prj, ds)

    t.scenario(
        when.view_dataset(ds.dataset_id),
        then.text(f'main #ds-{ds.dataset_id} .location-uri', "N/A"),
    )