package agent

import "strings"

func ExtractCity(
	message string,
) string {
	idx := strings.LastIndex(
		strings.ToLower(message),
		"in ",
	)

	if idx == -1 {
		return ""
	}

	return strings.Trim(
		message[idx+3:],
		" ?.",
	)
}
