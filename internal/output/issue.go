package output

import (
	"fmt"
	"image/color"
	"strings"
	"time"

	"github.com/StevenACoffman/j2m"
	"github.com/andygrunwald/go-jira"
	"github.com/cli/cli/pkg/iostreams"
	"github.com/cli/cli/pkg/markdown"
	"github.com/cli/cli/pkg/text"
	"github.com/cli/cli/utils"
	"github.com/muesli/gamut"
)

func colorNameFromHex(hex string) string {
	var p gamut.Palette
	p.AddColors(
		gamut.Colors{
			{"magenta", gamut.Hex("#ff00ff"), "Reference"},
			{"cyan", gamut.Hex("#00ffff"), "Reference"},
			{"red", gamut.Hex("#ff0000"), "Reference"},
			{"yellow", gamut.Hex("#ffff00"), "Reference"},
			{"blue", gamut.Hex("#0000ff"), "Reference"},
			{"green", gamut.Hex("#008000"), "Reference"},
			{"gray", gamut.Hex("#808080"), "Reference"},
		},
	)

	return p.Clamped([]color.Color{gamut.Hex(hex)})[0].Name
}

func PrintIssueList(io *iostreams.IOStreams, priorities []jira.Priority, issues []jira.Issue) error {
	printer := utils.NewTablePrinter(io)
	colorScheme := io.ColorScheme()

	// TODO cache colors
	priorityColors := make(map[string]string)
	for _, priority := range priorities {
		priorityColors[priority.Name] = colorNameFromHex(priority.StatusColor)
	}

	for _, issue := range issues {
		color := strings.Split(issue.Fields.Status.StatusCategory.ColorName, "-")[0] // ignore background

		printer.AddField(issue.Key, nil, colorScheme.ColorFromString(color))

		printer.AddField(issue.Fields.Summary, nil, nil)

		priority := issue.Fields.Priority.Name
		if len(priority) > 3 {
			priority = priority[:3]
		}
		printer.AddField(priority, nil, colorScheme.ColorFromString(strings.TrimSpace(priorityColors[issue.Fields.Priority.Name]))) // TODO handle err when color is not in map

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

func PrintIssue(io *iostreams.IOStreams, issue *jira.Issue, priorities []jira.Priority, comments bool) error {
	description, err := markdown.Render(j2m.JiraToMD(issue.Fields.Description), "dark") // TODO glamour style from config
	if err != nil {
		return err
	}

	components := make([]string, len(issue.Fields.Components))
	for index, component := range issue.Fields.Components {
		components[index] = component.Name
	}

	// TODO cache colors
	priorityColors := make(map[string]string)
	for _, priority := range priorities {
		priorityColors[priority.Name] = colorNameFromHex(priority.StatusColor)
	}

	cs := io.ColorScheme()
	fmt.Fprintf(io.Out, "%v %v [%v]\n%v %v • opened %v • %v comment(s)\nComponents: %v\nLabels: %v\n%v\n",
		cs.Bold(issue.Key),
		cs.Bold(issue.Fields.Summary),
		cs.ColorFromString(priorityColors[issue.Fields.Priority.Name])(issue.Fields.Priority.Name), // TODO handle err when color is not in map
		cs.ColorFromString(strings.SplitN(issue.Fields.Status.StatusCategory.ColorName, "-", 2)[0])(issue.Fields.Status.Name),
		issue.Fields.Type.Name,
		utils.FuzzyAgo(time.Since(time.Time(issue.Fields.Created))),
		len(issue.Fields.Comments.Comments),
		cs.Gray(strings.Join(components, ", ")),
		cs.Gray(strings.Join(issue.Fields.Labels, ", ")),
		description)

	for index, comment := range issue.Fields.Comments.Comments {
		// TODO optimize
		newest := ""
		if index == len(issue.Fields.Comments.Comments)-1 {
			newest = fmt.Sprintf(" • %v", cs.Cyan("Newest comment"))
		} else if !comments {
			continue
		}

		body, err := markdown.Render(j2m.JiraToMD(comment.Body), "dark") // TODO glamour style from config
		if err != nil {
			return err
		}

		created, err := time.Parse("2006-01-02T15:04:05Z0700", comment.Created)
		if err != nil {
			return err
		}

		edited := ""
		if comment.Created != comment.Updated {
			edited = " • Edited"
		}

		fmt.Fprintf(io.Out, "%v • %v%v%v\n%v\n",
			cs.Bold(comment.Author.DisplayName),
			utils.FuzzyAgo(time.Since(time.Time(created))),
			edited,
			newest,
			body,
		)
	}

	return nil

}
