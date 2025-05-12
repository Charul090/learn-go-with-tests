package iterations

import "strings"

func Repeated(s string, count int) string {
	var repeated strings.Builder
	for i := 0; i < count; i++ {
		repeated.WriteString(s)
	}
	return repeated.String()
}
