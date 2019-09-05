from env import *
from client import ml_pipelines as client

def given_empty_system(e: Env):
    """create project if all arguments are valid"""

    p = preset.project_created(e)
    ds = preset.dataset_created(e, p)
    v1 = preset.dataset_version_added(e,ds)

    e.given_events(p, ds, v1)

    def checkpoint_story(c:client.Client):

        prj = c.get_project(p.name)
        job = prj.get_or_create_job('import')

        input_ds = prj.get_dataset(ds.name)

        input_ver = input_ds.get_last_version()



        #####

        run = job.start_run(title="some job title", inputs=[input_ver])

        try:

            output_ds = prj.get_or_create_dataset('some_output', location_id="gcs")

            output_ver = output_ds.get_last_version()

            for i in range(10):

                #####
                run.log(f".... Loops {i} ....")


                staging = output_ver.prepare_commit()

                staging.add_file(name=f'file_{i}', records=100, size=2000)
                staging.add_input(run)

                output_ver = staging.commit("Update from the job run")


            run.complete()
        except Exception as e:
            run.fail(e)

    def snapshot_story(c: client.Client):

        prj = c.get_project(p.name)

        job = prj.get_or_create_job('snapshot_table')

        source = prj.get_dataset(ds.name)
        #####

        run = job.start_run(title="Our next snapshot", inputs=[source.get_last_version()])

        try:
            run.log("Starting heavy import OR continue")

            # if continue previous
            # files = load_files_from_somewhere()

            output = prj.get_or_create_dataset('some_output', location_id='dl').get_last_version()

            staging = output.prepare_commit(clean_slate=True)
            staging.add_input(run)




            for i in range(10):
                run.log("Work happens here")
                staging.add_file(f'file_i', size=122, records=1233)


            staging.commit("New version from the job")
            # we don't need the job to complete in order to commit
            run.complete()
        except Exception as e:
            run.fail(e)

    e.scenario(
        when.client(checkpoint_story),
        then.client_ok()
    )
    e.scenario(
        when.client(snapshot_story),
        then.client_ok()
    )


