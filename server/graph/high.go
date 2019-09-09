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

func RenderServiceGraph(tx *db.Tx, s *shared.Site, uid []byte) template.HTML{
	render := NewRender(tx, s.Url, s.Fmt)

	render.Service(uid)
	render.ColorGreen(uid)

	renderServiceLinks(render, uid)

	return render.ToHtml()
}

func renderServiceLinks(render *SvgRender, uid []byte) {
	refs := db.ListServiceLinks(render.tx, uid)
	for _, ref := range refs {
		switch ref.Type {
		case db.ServiceLink_Input_ServiceVer:
			render.ServiceVer(ref.InstanceUid)
			render.Arrow(ref.InstanceUid, uid)
		case db.ServiceLink_Input_JobRun:

			// service -- job.run -- Outputs
			render.JobRun(ref.InstanceUid)
			render.Dash(ref.InstanceUid, uid)

			renderJobRunInputs(render, ref.InstanceUid)
		case db.ServiceLink_Output_ServiceVer:
			// Service THIS -> Service XXX
			render.ServiceVer(ref.InstanceUid)
			render.Arrow(uid, ref.InstanceUid)

		case db.ServiceLink_Output_JobRun:
			// THIS -> job run
			render.JobRun(ref.InstanceUid)
			render.Arrow(uid, ref.InstanceUid)

			renderJobRunOutputs(render, ref.InstanceUid)
		default:
			log.Panicf("Unexpected type %s", ref.Type)

		}
	}
}


func RenderServiceVersionGraph(tx *db.Tx, s *shared.Site, uid []byte) template.HTML{
	render := NewRender(tx, s.Url, s.Fmt)

	render.ServiceVer(uid)
	render.ColorGreen(uid)




	this := db.GetServiceVersion(tx, uid)

	//renderServiceLinks(render, this.ServiceUid)

	for _, input := range this.Inputs{
		switch input.Type {
		case vo.ServiceVersionInput_Service:
			render.Service(input.Uid)
			render.Arrow(input.Uid, uid)
		case vo.ServiceVersionInput_JobRun:
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
		case vo.ServiceVersionOutput_Service:
			render.Service(output.Uid)
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
		case db.JobRunOutput_ServiceVer:
			render.ServiceVer(output.Uid)
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
		case vo.JobRunInput_Service:
			render.Service(input.Uid)
			render.Dash(input.Uid, run.Uid)
		default:
			log.Panicf("Unknown job run input %s\n", input.Type)
		}

	}
}
