# go-jira-cli

Simple [Jira](https://www.atlassian.com/software/jira) terminal client based on [go-jira](https://github.com/andygrunwald/go-jira) and [github cli](https://github.com/cli/cli).

[![asciicast](https://asciinema.org/a/414802.svg)](https://asciinema.org/a/414802)

## Status

WIP

## Getting Started

Host config with optional cookie value or user/token for basic auth (anonymous if none is set):

```sh
#~/.config/gj/hosts.yaml
issues.apache.org/jira:
  cookie:
  user:
  token:
```

Source completion:
```sh
#bash
source <(gj _carapace)

# elvish
eval (gj _carapace|slurp)

# fish
gj _carapace | source

# oil
source <(gj _carapace)

# powershell
Set-PSReadlineKeyHandler -Key Tab -Function MenuComplete
gj _carapace | Out-String | Invoke-Expression

# xonsh
COMPLETIONS_CONFIRM=True
exec($(gj _carapace))

# zsh
source <(gj _carapace)
``
