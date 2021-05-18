package output

import (
	"fmt"
	"strings"
	"time"

	"github.com/andygrunwald/go-jira"
	"github.com/cli/cli/pkg/iostreams"
	"github.com/cli/cli/pkg/text"
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
		componentsText := strings.Join(components, ", ")
		if len(components) > 0 {
			componentsText = fmt.Sprintf("(%v)", componentsText)
		}
		printer.AddField(componentsText, func(i int, s string) string {
			if len(s) < 2 {
				return s
			}
			return fmt.Sprintf("(%v)", text.Truncate(i-2, s[1:len(s)-1]))
		}, colorScheme.Gray)

		ago := utils.FuzzyAgo(time.Since(time.Time(issue.Fields.Updated)))
		printer.AddField(ago, nil, colorScheme.Gray)
		printer.EndRow()
	}
	return printer.Render()
}
