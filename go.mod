module github.com/rsteube/go-jira-cli

go 1.14

require (
	github.com/StevenACoffman/j2m v0.0.0-20190826163711-7d8d00c99217
	github.com/andygrunwald/go-jira v1.13.0
	github.com/cli/browser v1.1.0
	github.com/cli/cli v1.9.2
	github.com/cli/safeexec v1.0.0
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510
	github.com/rsteube/carapace v0.5.12
	github.com/spf13/cobra v1.1.3
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/rsteube/carapace => ../carapace
