package output

import (
	"strings"
	"time"

	"github.com/andygrunwald/go-jira"
	"github.com/cli/cli/pkg/iostreams"
	"github.com/cli/cli/utils"
)

func PrintIssues(io *iostreams.IOStreams, issues []jira.Issue) error {
	printer := utils.NewTablePrinter(io)
	colorScheme := io.ColorScheme()
	for _, issue := range issues {
		color := strings.Split(issue.Fields.Status.StatusCategory.ColorName, "-")[0] // ignore background

		printer.AddField(issue.Key, nil, colorScheme.ColorFromString(color))
		printer.AddField(issue.Fields.Summary, nil, nil)

		components := make([]string, len(issue.Fields.Components))
		for index, component := range issue.Fields.Components {
			components[index] = component.Name
		}
		printer.AddField(strings.Join(components, ", "), nil, colorScheme.Gray)

		ago := utils.FuzzyAgo(time.Since(time.Time(issue.Fields.Updated)))
		printer.AddField(ago, nil, colorScheme.Gray)
		printer.EndRow()
	}
	return printer.Render()
}
