package action

func statusColor(s string) string {
	return map[string]string{
		"medium-gray": "#7d7f7c",
		"green":       "#008000",
		"yellow":      "#FFFF00",
		"brown":       "#A52A2A",
		"warm-red":    "#E6534E",
		"blue-gray":   "#6699CC",
	}[s]
}
