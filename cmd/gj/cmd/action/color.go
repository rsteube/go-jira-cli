package action

func statusColor(s string) string {
	return map[string]string{
		"medium-gray": "dim white",
		"green":       "green",
		"yellow":      "yellow",
		"brown":       "magenta",
		"warm-red":    "red",
		"blue-gray":   "blue",
	}[s]
}
