package output

import (
	"fmt"

	"github.com/andygrunwald/go-jira"
	"github.com/cli/cli/pkg/iostreams"
	"github.com/cli/cli/utils"
)

func PrintProjectList(io *iostreams.IOStreams, projects []jira.ProjectInfo) error {
	printer := utils.NewTablePrinter(io)
	cs := io.ColorScheme()
	for _, project := range projects {
		printer.AddField(project.Key, nil, nil)
		printer.AddField(project.Name, nil, cs.Gray)
		printer.AddField(project.ProjectCategory.Name, nil, cs.Gray)
		printer.EndRow()
	}
	return printer.Render()
}

func PrintProject(io *iostreams.IOStreams, project *jira.Project) error {
	cs := io.ColorScheme()

	fmt.Fprintln(io.Out, project.Key)
	fmt.Fprintln(io.Out, project.Name)
	fmt.Fprintln(io.Out, cs.Gray(project.Description))
	fmt.Fprintln(io.Out, project.Lead.Name)
	fmt.Fprintln(io.Out, project.ProjectCategory.Name)

	return nil
}
