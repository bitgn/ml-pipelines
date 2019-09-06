from env import *


def given_an_existing_job(e: Env):
    """Can't create entity with the same name"""
    prj = preset.project_created(e)
    job = preset.job_added(e, prj)

    e.given_events(prj, job)

    e.scenario(
        when.client(lambda c:c[prj.name].create_dataset(job.name)),
        then.already_exists(subject_uid=job.uid, subject_name=job.name)
    )

    e.scenario(
        when.client(lambda c:c[prj.name].create_job(job.name)),
        then.already_exists(subject_uid=job.uid, subject_name=job.name)
    )
