package graph

import (
	"encoding/base64"
	"encoding/hex"
	"html/template"
	"log"
	"mlp/catalog/db"
	"mlp/catalog/vo"
	"mlp/catalog/web/shared"
)

func RenderSystemGraph(tx *db.Tx, s *shared.Site, uid []byte) template.HTML{
	render := NewRender(tx, s.Url, s.Fmt)

	render.System(uid)
	render.ColorGreen(uid)

	renderSystemLinks(render, uid)

	return render.ToHtml()
}

func renderSystemLinks(render *SvgRender, uid []byte) {
	refs := db.ListSystemLinks(render.tx, uid)
	for _, ref := range refs {
		switch ref.Type {
		case db.SystemLink_Input_SystemVer:
			render.SystemVer(ref.InstanceUid)
			render.Arrow(ref.InstanceUid, uid)
		case db.SystemLink_Input_JobRun:

			// service -- job.run -- Outputs
			render.JobRun(ref.InstanceUid)
			render.Dash(ref.InstanceUid, uid)

			renderJobRunInputs(render, ref.InstanceUid)
		case db.SystemLink_Output_SystemVer:
			// System THIS -> System XXX
			render.SystemVer(ref.InstanceUid)
			render.Arrow(uid, ref.InstanceUid)

		case db.SystemLink_Output_JobRun:
			// THIS -> job run
			render.JobRun(ref.InstanceUid)
			render.Arrow(uid, ref.InstanceUid)

			renderJobRunOutputs(render, ref.InstanceUid)
		default:
			log.Panicf("Unexpected type %s", ref.Type)

		}
	}
}


func RenderSystemVersionGraph(tx *db.Tx, s *shared.Site, uid []byte) template.HTML{
	render := NewRender(tx, s.Url, s.Fmt)

	render.SystemVer(uid)
	render.ColorGreen(uid)




	this := db.GetSystemVersion(tx, uid)

	//renderSystemLinks(render, this.SystemUid)

	for _, input := range this.Inputs{
		switch input.Type {
		case vo.SystemVersionInput_System:
			render.System(input.Uid)
			render.Arrow(input.Uid, uid)
		case vo.SystemVersionInput_JobRun:
			// e.g. report is deployed from a job
			render.JobRun(input.Uid)
			render.Arrow(input.Uid, uid)

			renderJobRunInputs(render, input.Uid)
		default:
			log.Panicf("Unknown service input %s", input.Type)

		}
	}

	for _, output := range this.Outputs{
		switch output.Type {
		case vo.SystemVersionOutput_System:
			render.System(output.Uid)
			render.Arrow(uid, output.Uid)

		}
	}


	return render.ToHtml()
}


func  RenderDatasetVerGraph(tx *db.Tx, s *shared.Site, uid []byte) template.HTML{

	render := NewRender(tx, s.Url, s.Fmt)
	this := db.GetDatasetVersion(render.tx, uid)


	render.DatasetVer(uid)
	render.ColorGreen(uid)

	if this != nil {



		for _, input := range this.Inputs {

			switch input.Type {
			case vo.DatasetVerInput_JOB_RUN:
				render.JobRun(input.Uid)
				render.Arrow(input.Uid, uid)

				renderJobRunInputs(render, input.Uid)

			default:
				log.Panicln("Unknown dataset version type")

			}

		}

		for _, out := range this.Outputs {

			switch out.Type {
			case vo.DatasetVerOutput_JOB_RUN:


				render.JobRun(out.Uid)
				render.Dash(uid, out.Uid)

				renderJobRunOutputs(render, out.Uid)
			default:
				log.Panicln("Unknown type")
			}

		}
	}
	return render.ToHtml()

}



func renderJobRunOutputs(render *SvgRender, uid []byte){
	run := db.GetJobRun(render.tx, uid)

	if run == nil {
		log.Panicln("Can't find job RUN", hex.EncodeToString(uid), base64.StdEncoding.EncodeToString(uid))
	}

	for _, output := range run.Outputs {

		switch output.Type {
		case db.JobRunOutput_DatasetVer:
			render.DatasetVer(output.Uid)
			render.Arrow(run.Uid, output.Uid)
		case db.JobRunOutput_SystemVer:
			render.SystemVer(output.Uid)
			render.Arrow(run.Uid, output.Uid)
		default:
			log.Panicln("Unknown job run output")
		}

	}

}

func renderJobRunInputs(render *SvgRender, run_uid []byte) {
	run := db.GetJobRun(render.tx, run_uid)
	for _, input := range run.Inputs {

		switch input.Type {

		case vo.JobRunInput_DatasetVer:
			render.DatasetVer(input.Uid)
			render.Dash(input.Uid, run.Uid)
		case vo.JobRunInput_System:
			render.System(input.Uid)
			render.Dash(input.Uid, run.Uid)
		default:
			log.Panicf("Unknown job run input %s\n", input.Type)
		}

	}
}
