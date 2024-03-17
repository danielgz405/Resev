package utils

import "regexp"

func IsTimeFormat(s string) bool {
	re := regexp.MustCompile(`^\d{2}:\d{2}$`)
	return re.MatchString(s)
}
