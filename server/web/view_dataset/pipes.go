package view_dataset

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"
	"html/template"
	"io"
	"log"
	"mlp/catalog/db"
	"mlp/catalog/graph"
	"mlp/catalog/vo"
	"mlp/catalog/web/shared"
	"os/exec"
	"strings"
)

func renderDot(line string) ([]byte, error){
	cmd := exec.Command("dot", "-Tsvg")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		_, err := io.WriteString(stdin, line)
		if err != nil {
			log.Fatal(errors.Wrap(err, "Failed to write to stdin"))
		}

	}()

	return cmd.CombinedOutput()
}
/*
def build_lineage(dataset_id, tx):
    ds = db.dataset_get(tx, dataset_id)
    dot = g.Digraph(comment='lineage',
                    node_attr={'shape': 'rectangle', 'color': '#343a40',
                               'penwidth': "1.5", 'fontname': 'Arial'},
                    edge_attr={'penwidth': '1.0', 'color': '#343a40'})
    dot.attr(rankdir='LR', fontname='Arial')
    dot.node("this", label=ds.name, color="#28a745", bgcolor="#28a745")
    edge_node = {'style': 'rounded'}
    for l in ds.upstream_jobs:
        job = db.job_get(tx, l)

        if not job:
            print(f'no link for id {l}')

        dot.node(job.job_id, label=job.job_name, **edge_node)
        dot.edge(job.job_id, "this")

        for input_id in job.inputs:
            input_ds = db.dataset_get(tx, input_id)
            dot.node(input_id, label=input_ds.name, href=urls.reverse("explore:view_dataset", args=[input_id]))

            dot.edge(input_id, job.job_id, arrowhead='none')
    for l in ds.downstream_jobs:
        job = db.job_get(tx, l)
        dot.node(job.job_id, label=job.job_name, **edge_node)
        dot.edge("this", job.job_id, arrowhead='none')

        for output_id in job.outputs:
            output_ds = db.dataset_get(tx, output_id)
            dot.node(output_id, label=output_ds.name, href=urls.reverse("explore:view_dataset", args=[output_id]))
            dot.edge(job.job_id, output_id, arrowtail='none')
    return dot
 */



func hx(b []byte) string{
	return hex.EncodeToString(b)
}

func renderDatasetVersionSVG(tx *db.Tx, uid []byte, url shared.UrlResolver, f shared.Format, title string) template.HTML{

	this := db.GetDatasetVersion(tx, uid)

	if this != nil {
		title = fmt.Sprintf("<%s<BR/><I>version %d</I>>", title, this.VersionNum)
	}




	render := graph.NewRender(tx, url, f)

	render.Line("digraph{")
	render.Line("rankdir=LR;")
	render.Line("fontname=\"Arial\";")
	render.Line("node[shape=\"rectangle\" color=\"#343a40\" penwidth=\"1.5\" fontname=\"Arial\"];")
	render.Line("edge[color=\"#343a40\" penwidth=\"1.0\"];")
	render.Line("\"%s\" [label=%s color=\"#28a745\"];", hx(uid), title)





	if this != nil {



		for _, input := range this.Inputs {

			switch input.Type {
			case vo.DatasetVerInput_JOB_RUN:

				run := db.GetJobRun(tx, input.Uid)

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

				run := db.GetJobRun(tx, out.Uid)

				if run == nil {
					log.Panicln("Can't find job RUN", hx(out.Uid), base64.StdEncoding.EncodeToString(out.Uid))
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

	render.Line("}")

	//println(sb.String())

	result, err := renderDot(render.String())

	//println(string(result))

	if err != nil {
		log.Println("Problem with SVG:", string(result))
		log.Println("Original input:")
		log.Println(render.String())
		return template.HTML(err.Error())
	}

	str := string(result)

	str = strings.Replace(str, "<svg ", `<svg class="img-fluid" `, 1)
	return template.HTML(str)

}


