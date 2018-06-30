package ustr

import (
	"strings"
)

// NamedPlaceholders is a possible alternative to `fmt.Sprintf` and
// `strings.Replace` / `Replacer` for some (not all) "micro-templating" use-cases.
func NamedPlaceholders(begin byte, end byte) func(string, ...string) string {
	return func(s string, namesAndValues ...string) string {
		if len(s) >= 3 && len(namesAndValues) > 0 {
			var sb strings.Builder
			start, pos := 0, -1
			for i := 0; i < len(s); i++ {
				if s[i] == begin {
					pos = i
				} else if s[i] == end && pos >= 0 {
					nidx, name := -1, s[pos+1:i]
					for j := 0; j < len(namesAndValues); j += 2 {
						if namesAndValues[j] == name {
							nidx = j
							break
						}
					}
					if nidx >= 0 && nidx < len(namesAndValues)-1 {
						sb.WriteString(s[start:pos])
						sb.WriteString(namesAndValues[nidx+1])
						start = i + 1
					} else {
						pos = -1
					}
				}
			}
			if start > 0 {
				sb.WriteString(s[start:])
			}
			if sb.Len() > 0 {
				return sb.String()
			}
		}
		return s
	}
}
