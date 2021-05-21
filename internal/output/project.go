package output

import (
	"fmt"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/andygrunwald/go-jira"
	"github.com/cli/cli/pkg/iostreams"
	"github.com/cli/cli/pkg/markdown"
	"github.com/cli/cli/utils"
	"github.com/rsteube/go-jira-cli/internal/api"
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

func PrintProject(io *iostreams.IOStreams, project *jira.Project, activities *api.ActivityStream) error {
	cs := io.ColorScheme()

	fmt.Fprintln(io.Out, project.Key)
	fmt.Fprintln(io.Out, project.Name)
	fmt.Fprintln(io.Out, cs.Gray(project.Description))
	fmt.Fprintln(io.Out, project.Lead.Name)
	fmt.Fprintln(io.Out, project.ProjectCategory.Name)

	converter := md.NewConverter("", true, &md.Options{
		LinkStyle:          "referenced",
		LinkReferenceStyle: "shortcut",
	})
	for _, activity := range activities.Entry {
		mdTitle, err := converter.ConvertString(activity.Title)
		if err != nil {
			return err
		}
		fmt.Fprintln(io.Out, cs.Bold(strings.SplitN(mdTitle, "\n", 2)[0]))

		mdContent, err := converter.ConvertString(activity.Content)
		if err != nil {
			return err
		}

		renderedContent, err := markdown.Render(mdContent, "dark")
		if err != nil {
			return err
		}
		fmt.Fprintln(io.Out, renderedContent)
	}

	return nil
}
