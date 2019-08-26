package view_dataset

import (
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"
	"html/template"
	"io"
	"log"
	"mlp/catalog/db"
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

func renderSVG(tx *db.Tx, dataset_id []byte, url shared.UrlResolver) template.HTML{

	this := db.GetDataset(tx, dataset_id)

	hx := hex.EncodeToString


	var sb strings.Builder
	sb.WriteString("digraph{\n")
	sb.WriteString("rankdir=LR;\n")
	sb.WriteString("fontname=\"Arial\";\n")
	sb.WriteString("node[shape=\"rectangle\" color=\"#343a40\" penwidth=\"1.5\" fontname=\"Arial\"];\n")
	sb.WriteString("edge[color=\"#343a40\" penwidth=\"1.0\"]\n;")
	sb.WriteString(fmt.Sprintf("this [label=\"%s\" color=\"#28a745\"]; \n", this.Title))


	for _, j := range this.UpstreamJobs{
		job := db.GetJob(tx, j)

		sb.WriteString(fmt.Sprintf("  \"%s\" [label=\"%s\" style=\"rounded\"]; \n",  hx(job.Uid), job.Title))
		sb.WriteString(fmt.Sprintf("  \"%s\" -> this;\n", hx(job.Uid)))

		for _, input := range job.Inputs{
			ds := db.GetDataset(tx, input.SourceId)
			du := url.ViewDataset(ds.ProjectName, ds.Name)
			sb.WriteString(fmt.Sprintf("  \"%s\" [label=\"%s\" href=\"%s\"];\n", hx(ds.Uid), ds.Title, du))
			sb.WriteString(fmt.Sprintf("  \"%s\" -> \"%s\" [arrowhead=\"none\"];", hx(ds.Uid), hx(job.Uid)))
		}

	}

	for _, j := range this.DownstreamJobs{
		job := db.GetJob(tx, j)

		sb.WriteString(fmt.Sprintf("  \"%s\" [label=\"%s\" style=\"rounded\"];\n", hx(job.Uid), job.Title))
		sb.WriteString(fmt.Sprintf("  this -> \"%s\" [arrowhead=\"none\"];\n", hx(job.Uid)))

		for _, output := range job.Outputs{
			ds := db.GetDataset(tx, output.TargetId)
			dataset := url.ViewDataset(ds.ProjectName, ds.Name)
			sb.WriteString(fmt.Sprintf("  \"%s\" [label=\"%s\" href=\"%s\"];\n", hx(ds.Uid), ds.Title, dataset))
			sb.WriteString(fmt.Sprintf("  \"%s\" -> \"%s\" [arrowtail=\"none\"];", hx(job.Uid), hx(ds.Uid)))

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
