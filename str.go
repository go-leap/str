package ustr

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var (
	// Fmt aliases `fmt.Sprintf` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `fmt` import.
	Fmt = fmt.Sprintf

	// Has aliases `strings.Contains` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	Has = strings.Contains

	// Idx aliases `strings.IndexRune` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	Idx = strings.IndexRune

	// Int aliases `strconv.Itoa` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strconv` import.
	Int = strconv.Itoa

	// Join aliases `strings.Join` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	Join = strings.Join

	// Last aliases `strings.LastIndex` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	Last = strings.LastIndex

	// Lo aliases `strings.ToLower` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	Lo = strings.ToLower

	// NumRunes aliases `unicode/utf8.RuneCountInString` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `unicode/utf8` import.
	NumRunes = utf8.RuneCountInString

	// Pos aliases `strings.Index` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	Pos = strings.Index

	// Pref aliases `strings.HasPrefix` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	Pref = strings.HasPrefix

	// Reader aliases `strings.NewReader` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	Reader = strings.NewReader

	// Repl aliases `strings.NewReplacer` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	Repl = strings.NewReplacer

	// Suff aliases `strings.HasSuffix` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	Suff = strings.HasSuffix

	// Times aliases `strings.Repeat` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	Times = strings.Repeat

	// Trim aliases `strings.TrimSpace` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	Trim = strings.TrimSpace

	// TrimL aliases `strings.TrimLeft` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	TrimL = strings.TrimLeft

	// TrimPref aliases `strings.TrimPrefix` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	TrimPref = strings.TrimPrefix

	// TrimR aliases `strings.TrimRight` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	TrimR = strings.TrimRight

	// TrimSuff aliases `strings.TrimSuffix` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	TrimSuff = strings.TrimSuffix

	// Up aliases `strings.ToUpper` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	Up = strings.ToUpper
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

// EnsureCase returns `s` with the rune at `runeIndex` (not byte index) guaranteed to be upper-case if `upper`, or lower-case if not.
func EnsureCase(s string, runeIndex int, upper bool) string {
	f := unicode.ToLower
	if upper {
		f = unicode.ToUpper
	}
	runes := []rune(s)
	runes[runeIndex] = f(runes[runeIndex])
	return string(runes)
}

// Fewest returns the `s` in `strs` with the lowest `strings.Count` of `substr`.
// If the count is identical for all, it returns `otherwise(strs)` (if supplied).
func Fewest(strs []string, substr string, otherwise func([]string) string) (s string) {
	lastnum, counts := math.MaxInt64, make(map[int]bool, len(strs))
	for _, str := range strs {
		num := strings.Count(str, substr)
		if counts[num] = true; num < lastnum {
			s, lastnum = str, num
		}
	}
	if len(counts) == 1 && otherwise != nil {
		s = otherwise(strs)
	}
	return
}

// Filt returns all `strs` that satisfy `check`.
func Filt(strs []string, check func(string) bool) (filtered []string) {
	if filtered = strs; len(strs) > 0 && check != nil {
		filtered = make([]string, 0, len(strs))
		for _, s := range strs {
			if check(s) {
				filtered = append(filtered, s)
			}
		}
	}
	return
}

// FirstIn returns the first in `subStrings` to satisfy `strings.Contains(s, substr)`, or `""`.
func FirstIn(s string, subStrings ...string) string {
	for _, substr := range subStrings {
		if strings.Contains(s, substr) {
			return substr
		}
	}
	return ""
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

func Int64(i int64) string {
	return strconv.FormatInt(i, 10)
}

// Has1Of returns whether `s` contains any of the specified `subStrings`.
func Has1Of(s string, subStrings ...string) bool {
	return FirstIn(s, subStrings...) != ""
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

// Longest returns the longest `s` in `strs`.
func Longest(strs []string) (s string) {
	for _, str := range strs {
		if len(str) > len(s) {
			s = str
		}
	}
	return
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

// Pref1Of returns the first of the specified (non-empty) `prefixes` that `s` begins with, or `""`.
func Pref1Of(s string, prefixes ...string) string {
	for _, prefix := range prefixes {
		if prefix != "" && strings.HasPrefix(s, prefix) {
			return prefix
		}
	}
	return ""
}

// Replace allocates a one-off throw-away `strings.NewReplacer` to perform the specified replacements.
func Replace(s string, oldNewPairs ...string) string {
	if len(oldNewPairs) == 0 {
		return s
	}
	return strings.NewReplacer(oldNewPairs...).Replace(s)
}

// Sans returns `strs` without the specified `excludedStrs`.
func Sans(strs []string, excludedStrs ...string) []string {
	return Filt(strs, func(s string) bool {
		return !In(s, excludedStrs...)
	})
}

// Shortest returns the shortest `s` in `strs`.
func Shortest(strs []string) (s string) {
	for _, str := range strs {
		if s == "" || len(str) < len(s) {
			s = str
		}
	}
	return
}

// Split returns an empty slice if `s` is emtpy, otherwise calls `strings.Split`.
func Split(s string, sep string) (strs []string) {
	if len(s) != 0 {
		strs = strings.Split(s, sep)
	}
	return
}

//	SplitByWhitespaceAndJoin returns `s` with all occurrences of multiple subsequent `unicode.IsSpace` runes in a row collapsed into one single white-space (`" "`) rune.
func SplitByWhitespaceAndReJoinBySpace(s string) string {
	return strings.Join(strings.Fields(s), " ")
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
