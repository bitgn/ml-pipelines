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

	return render.ToHtml()
}


func RenderServiceVersionGraph(tx *db.Tx, s *shared.Site, uid []byte) template.HTML{
	render := NewRender(tx, s.Url, s.Fmt)

	render.ServiceVer(uid)
	render.ColorGreen(uid)


	this := db.GetServiceVersion(tx, uid)

	for _, input := range this.Inputs{
		switch input.Type {
		case vo.ServiceVersionInput_Service:
			render.Service(input.Uid)
			render.Dash(input.Uid, uid)
		default:
			log.Panicf("Unknown service input %s", input.Type)

		}
	}

	for _, output := range this.Outputs{
		switch output.Type {
		case vo.ServiceVersionOutput_Service:
			render.Service(output.Uid)
			render.Dash(output.Uid, uid)

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

				run := db.GetJobRun(render.tx, input.Uid)

				render.JobRun(input.Uid)
				render.Arrow(input.Uid, uid)

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

			default:
				log.Panicln("Unknown dataset version type")

			}

		}

		for _, out := range this.Outputs {

			switch out.Type {
			case vo.DatasetVerOutput_JOB_RUN:

				run := db.GetJobRun(render.tx, out.Uid)

				if run == nil {
					log.Panicln("Can't find job RUN", hex.EncodeToString(out.Uid), base64.StdEncoding.EncodeToString(out.Uid))
				}

				render.JobRun(out.Uid)
				render.Dash(uid, out.Uid)

				for _, output := range run.Outputs {

					switch output.Type {
					case vo.JobRunOutput_DatasetVer:
						render.DatasetVer(output.Uid)
						render.Arrow(run.Uid, output.Uid)
					default:
						log.Panicln("Unknown job run output")
					}

				}

			default:
				log.Panicln("Unknown type")
			}

		}
	}
	return render.ToHtml()


}
