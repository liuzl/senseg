package senseg

import (
	"regexp"
)

var (
	ReNewLine = regexp.MustCompile(`(?is)\n+|[-#=_+*]{4,}`)
)

func Cut(text string) []string {
	return nil
}
