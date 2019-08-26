from env import *



def given_a_dataset_with_metadata(t: Env):
    """show dataset properties"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)

    t.given_events(prj, ds)

    href = urls.view_project(ds.project_name)
    t.scenario(
        when.view_dataset(ds.project_name, ds.name),
        then.link(f'main .project-link', href=href, text=prj.name),
    )

    t.scenario(
        when.list_datasets(),
        then.link(f'main #ds-{ds.uid.hex()} .project-link', href=href, text=prj.name),

    )
