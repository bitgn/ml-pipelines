package view_dataset

import (
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"
	"html/template"
	"io"
	"log"
	"mlp/catalog/db"
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

func renderDatasetVersionSVG(tx *db.Tx, uid []byte, url shared.UrlResolver, title string) template.HTML{

	this := db.GetDatasetVersion(tx, uid)

	hx := hex.EncodeToString


	var sb strings.Builder
	sb.WriteString("digraph{\n")
	sb.WriteString("rankdir=LR;\n")
	sb.WriteString("fontname=\"Arial\";\n")
	sb.WriteString("node[shape=\"rectangle\" color=\"#343a40\" penwidth=\"1.5\" fontname=\"Arial\"];\n")
	sb.WriteString("edge[color=\"#343a40\" penwidth=\"1.0\"]\n;")
	sb.WriteString(fmt.Sprintf("this [label=\"%s\" color=\"#28a745\"]; \n", title))

	if this != nil {

		for _, j := range this.Inputs {

			switch j.Type {
			case vo.DatasetVerInput_JOB_RUN:

				run := db.GetJobRun(tx, j.Uid)

				sb.WriteString(fmt.Sprintf("  \"%s\" [label=\"%s\" style=\"rounded\"]; \n", hx(run.Uid), run.Title))
				sb.WriteString(fmt.Sprintf("  \"%s\" -> this;\n", hx(run.Uid)))

				for _, input := range run.Inputs {

					switch input.Type {

					case vo.JobRunInput_DatasetVer:
						ver := db.GetDatasetVersion(tx, input.Uid)
						ds := db.GetDataset(tx, ver.DatasetUid)

						link := url.ViewDatasetVersion(ds.ProjectName, ds.Name, ver.Uid)
						sb.WriteString(fmt.Sprintf("  \"%s\" [label=\"%s\" href=\"%s\"];\n", hx(ds.Uid), ds.Title, link))
						sb.WriteString(fmt.Sprintf("  \"%s\" -> \"%s\" [arrowhead=\"none\"];", hx(ds.Uid), hx(run.Uid)))

					default:
						log.Panicln("Unknown job run input")
					}

				}

			default:
				log.Panicln("Unknown dataset version type")

			}

		}

		for _, j := range this.Outputs {

			switch j.Type {
			case vo.DatasetVerOutput_JOB_RUN:

				run := db.GetJobRun(tx, j.Uid)

				sb.WriteString(fmt.Sprintf("  \"%s\" [label=\"%s\" style=\"rounded\"];\n", hx(run.Uid), run.Title))
				sb.WriteString(fmt.Sprintf("  this -> \"%s\" [arrowhead=\"none\"];\n", hx(run.Uid)))

				for _, output := range run.Outputs {

					switch output.Type {
					case vo.JobRunOutput_DatasetVer:
						ver := db.GetDatasetVersion(tx, output.Uid)
						ds := db.GetDataset(tx, ver.DatasetUid)
						link := url.ViewDatasetVersion(ds.ProjectName, ds.Name, ver.Uid)
						sb.WriteString(fmt.Sprintf("  \"%s\" [label=\"%s\" href=\"%s\"];\n", hx(ds.Uid), ds.Title, link))
						sb.WriteString(fmt.Sprintf("  \"%s\" -> \"%s\" [arrowtail=\"none\"];", hx(run.Uid), hx(ds.Uid)))

					default:
						log.Panicln("Unknown job run output")
					}

				}

			default:
				log.Panicln("Unknown type")
			}

		}
	}


	sb.WriteString("}")

	//println(sb.String())


	result, err := renderDot(sb.String())



	//println(string(result))

	if err != nil {
		log.Println("Problem with SVG:", string(result))
		log.Println("Original input:")
		log.Println(sb.String())
		return template.HTML(err.Error())
	}

	str := string(result)

	str = strings.Replace(str, "<svg ", `<svg class="img-fluid" `,1)
	return template.HTML(str)


}
