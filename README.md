# go-jira-cli

Simple [Jira](https://www.atlassian.com/software/jira) terminal client based on [go-jira](https://github.com/andygrunwald/go-jira) and [github cli](https://github.com/cli/cli).

[![asciicast](https://asciinema.org/a/414802.svg)](https://asciinema.org/a/414802)

## Status

WIP

## Example

```sh
docker-compose run --rm gj
gj issue view --host <TAB>
```

## Getting Started

```sh
gj auth login [host] # e.g. 'issues.apache.org/jira'
```
- **anonymous** without credentials
- **basic** auth with `username` and `token`
- **cookie** with `username` and `password` (only cookie will be stored)

### Dependencies

gj makes use of the [`bat` tool](https://github.com/sharkdp/bat), a `cat` clone
with color support.

Please install this using your local package manager beforehand.

### Shell completion

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
```
