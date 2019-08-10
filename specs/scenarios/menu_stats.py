from env import *



def given_empty_system(e: Env):
    """display zero counts for all entities"""
    e.scenario(
        when.list_projects(),

        then.text('.sidebar .project-count', "0"),
        then.text('.sidebar .dataset-count', "0"),
        then.text('.sidebar .model-count', "0"),
        then.text('.sidebar .expert-count', "0"),
        then.text('.sidebar .report-count', "0"),
        then.text('.sidebar .job-count', "0"),
    )


def given_system_with_entities(e:Env):
    """display appropriate counts"""


    prj = preset.project_created(e)





    e.given_events(prj,
                   preset.dataset_created(e, prj),
                   preset.dataset_created(e, prj),
                   preset.expert_added(e),
                   preset.expert_added(e),
                   preset.expert_added(e),

                   preset.job_added(e, prj),
                   preset.job_added(e, prj),
                   preset.job_added(e, prj),
                   preset.job_added(e, prj),
                   )



    e.scenario(
        when.list_projects(),

        then.text('.sidebar .project-count', "1"),
        then.text('.sidebar .dataset-count', "2"),
        then.text('.sidebar .model-count', "0"),
        then.text('.sidebar .expert-count', "3"),
        then.text('.sidebar .report-count', "0"),
        then.text('.sidebar .job-count', "4"),
    )