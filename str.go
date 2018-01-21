package ustr

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var (
	// Fmt aliases `fmt.Sprintf` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `fmt` import.
	Fmt = fmt.Sprintf

	// Pos aliases `strings.Index` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	Pos = strings.Index

	// Pref aliases `strings.HasPrefix` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	Pref = strings.HasPrefix

	// Join aliases `strings.Join` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	Join = strings.Join

	// Int aliases `strconv.Itoa` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strconv` import.
	Int = strconv.Itoa

	// Suff aliases `strings.HasSuffix` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	Suff = strings.HasSuffix

	// Trim aliases `strings.TrimSpace` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	Trim = strings.TrimSpace
)

// AfterFirst returns the suffix of `s` beginning right after the first occurrence of `needle`, or `otherwise` if no match.
func AfterFirst(s string, needle string, otherwise string) string {
	if i := strings.Index(s, needle); i >= 0 {
		return s[i+len(needle):]
	}
	return otherwise
}

// AfterLast returns the suffix of `s` beginning right after the last occurrence of `needle`, or `otherwise` if no match.
func AfterLast(s string, needle string, otherwise string) string {
	if i := strings.LastIndex(s, needle); i >= 0 {
		return s[i+len(needle):]
	}
	return otherwise
}

// BeforeFirst returns the prefix of `s` up to the first occurrence of `needle`, or `otherwise` if no match.
func BeforeFirst(s string, needle string, otherwise string) string {
	if i := strings.Index(s, needle); i >= 0 {
		return s[:i]
	}
	return otherwise
}

// BeforeLast returns the prefix of `s` up to the last occurrence of `needle`, or `otherwise` if no match.
func BeforeLast(s string, needle string, otherwise string) string {
	if i := strings.LastIndex(s, needle); i >= 0 {
		return s[:i]
	}
	return otherwise
}

// BeginsLower returns whether the first rune in `s` satisfies both `unicode.IsLetter` and `unicode.IsLower`.
func BeginsLower(s string) bool {
	for _, r := range s {
		return unicode.IsLetter(r) && unicode.IsLower(r)
	}
	return false
}

// BeginsUpper returns whether the first rune in `s` satisfies both `unicode.IsLetter` and `unicode.IsUpper`.
func BeginsUpper(s string) bool {
	for _, r := range s {
		return unicode.IsLetter(r) && unicode.IsUpper(r)
	}
	return false
}

// BreakOnFirst returns the prefix and suffix next to the first `needle` encountered in `s`.
// (If no match, `prefix` is `""` and `suffix` will be `s`.)
func BreakOnFirst(s string, needle string) (prefix string, suffix string) {
	if i := strings.Index(s, needle); i < 0 {
		suffix = s
	} else {
		prefix, suffix = s[:i], s[i+len(needle):]
	}
	return
}

// BreakOnLast returns the prefix and suffix next to the last `needle` encountered in `s`.
// (If no match, `prefix` is `""` and `suffix` will be `s`.)
func BreakOnLast(s string, needle string) (prefix string, suffix string) {
	if i := strings.LastIndex(s, needle); i < 0 {
		suffix = s
	} else {
		prefix, suffix = s[:i], s[i+len(needle):]
	}
	return
}

// Combine returns `s1` or `s2` or `s1 + sep + s2`, depending on their emptyness.
func Combine(s1 string, sep string, s2 string) string {
	if h1 := s1 != ""; h1 && s2 != "" {
		return s1 + sep + s2
	} else if h1 {
		return s1
	}
	return s2
}

// FirstOf returns the first non-empty `s` encountered in `strs`.
func FirstOf(strs ...string) (s string) {
	for _, s = range strs {
		if s != "" {
			break
		}
	}
	return
}

// In returns whether `strs` contains `s`.
func In(s string, strs ...string) bool {
	for _, str := range strs {
		if str == s {
			return true
		}
	}
	return false
}

// IsLower returns whether all `unicode.IsLetter` runes in `s` satisfy `unicode.IsLower`.
func IsLower(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.IsLower(r) {
			return false
		}
	}
	return true
}

// IsUpper returns whether all `unicode.IsLetter` runes in `s` satisfy `unicode.IsUpper`.
func IsUpper(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

// Map applies `f` to each `string` in `strs` and returns the results in `items`.
func Map(strs []string, f func(string) string) (items []string) {
	if f == nil {
		return strs
	}
	items = make([]string, len(strs))
	for i, s := range strs {
		items[i] = f(s)
	}
	return
}

// Split returns an empty slice if `s` is emtpy, otherwise calls `strings.Split`.
func Split(s string, sep string) (slice []string) {
	if len(s) != 0 {
		slice = strings.Split(s, sep)
	}
	return
}

// ToBool returns either the `bool` denoted by `s`, or `fallback`.
func ToBool(s string, fallback bool) bool {
	if v, err := strconv.ParseBool(s); err == nil {
		return v
	}
	return fallback
}

// ToFloat returns either the `float64` denoted by `s`, or `fallback`.
func ToFloat(s string, fallback float64) float64 {
	if v, err := strconv.ParseFloat(s, 64); err == nil {
		return v
	}
	return fallback
}

// ToInt returns either the `int64` denoted by `s`, or `fallback`.
func ToInt(s string, fallback int64) int64 {
	if v, err := strconv.ParseInt(s, 0, 64); err == nil {
		return v
	}
	return fallback
}

// ToUint returns either the `uint64` denoted by `s`, or `fallback`.
func ToUint(s string, fallback uint64) uint64 {
	if v, err := strconv.ParseUint(s, 0, 64); err == nil {
		return v
	}
	return fallback
}
