package ustr

import (
	"strings"
)

// ReplB replaces individual `byte`s in `s` based on the given old-new pairs: because
// for some few `strings.Replace` / `strings.Replacer` scenarios, nothing more is needed.
func ReplB(s string, oldNew ...byte) string {
	var buf strings.Builder
	start := 0
	for i, l, n := 0, len(s), len(oldNew)-1; i < l; i++ {
		for j := 0; j < n; j += 2 {
			if s[i] == oldNew[j] {
				buf.WriteString(s[start:i])
				buf.WriteByte(oldNew[j+1])
				start = i + 1
				break
			}
		}
	}
	if start > 0 {
		buf.WriteString(s[start:])
	}
	if buf.Len() > 0 {
		return buf.String()
	}
	return s
}

// NamedPlaceholders is an occasionally-preferable alternative to `fmt.Sprintf` or
// `strings.Replace` / `strings.Replacer` for (fully `string`ly-typed) "micro-templating":
//
// ```go
//  repl := ustr.NamedPlaceHolders('(', ']')
//  hi := repl("Hello (name]!", "name", "world")
// ```
//
//
// The delimiter `begin` and `end` chars may well be equal, if so desired. When
// called, the returned `func` traverses its `string` arg exactly  once, ie. it
// does not re-process the replacements or its final result. It searches through
// its name-value-pairs once per fully-delimited-substring. Any of those occurrences
// not found in its name-value-pairs are left in-place including the delimiters.
func NamedPlaceholders(begin byte, end byte) func(string, ...string) string {
	return func(s string, namesAndValues ...string) string {
		if len(s) >= 3 && len(namesAndValues) > 1 {
			var buf strings.Builder
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
						buf.WriteString(s[start:pos])
						buf.WriteString(namesAndValues[nidx+1])
						start = i + 1
					} else {
						pos = -1
					}
				}
			}
			if start > 0 {
				buf.WriteString(s[start:])
			}
			if buf.Len() > 0 {
				return buf.String()
			}
		}
		return s
	}
}
