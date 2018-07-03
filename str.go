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
	// Eq aliases `strings.EqualFold` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	Eq = strings.EqualFold

	// Fmt aliases `fmt.Sprintf` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `fmt` import.
	Fmt = fmt.Sprintf

	// From aliases `fmt.Sprint` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `fmt` import.
	From = fmt.Sprint

	// Has aliases `strings.Contains` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	Has = strings.Contains

	// IdxB aliases `strings.IndexByte` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	IdxB = strings.IndexByte

	// IdxR aliases `strings.IndexRune` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	IdxR = strings.IndexRune

	// Int aliases `strconv.Itoa` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strconv` import.
	Int = strconv.Itoa

	// Join aliases `strings.Join` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	Join = strings.Join

	// LastB aliases `strings.LastIndexByte` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	LastB = strings.LastIndexByte

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

	// TrimLR aliases `strings.Trim` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	TrimLR = strings.Trim

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

// BreakOnFirstOrSuff returns the prefix and suffix next to the first `needle` encountered in `s`.
// (If no match, `prefix` is `""` and `suffix` will be `s`.)
func BreakOnFirstOrSuff(s string, needle string) (prefix string, suffix string) {
	if i := strings.Index(s, needle); i < 0 {
		suffix = s
	} else {
		prefix, suffix = s[:i], s[i+len(needle):]
	}
	return
}

// BreakOnFirstOrPref returns the prefix and suffix next to the first `needle` encountered in `s`.
// (If no match, `suffix` is `""` and `prefix` will be `s`.)
func BreakOnFirstOrPref(s string, needle string) (prefix string, suffix string) {
	if i := strings.Index(s, needle); i < 0 {
		prefix = s
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

// Case returns `s` with the rune at `runeIndex` (not byte index) guaranteed to be upper-case if `upper`, or lower-case if not.
func Case(s string, runeIndex int, upper bool) string {
	f := unicode.ToLower
	if upper {
		f = unicode.ToUpper
	}
	return Rune(s, runeIndex, f)
}

// Case returns `s` with the rune at `runeIndex` (not byte index) guaranteed to be lower-case.
func CaseLo(s string, runeIndex int) string { return Rune(s, runeIndex, unicode.ToLower) }

// Case returns `s` with the rune at `runeIndex` (not byte index) guaranteed to be upper-case.
func CaseUp(s string, runeIndex int) string { return Rune(s, runeIndex, unicode.ToUpper) }

func CaseSnake(s string) string {
	var buf strings.Builder
	var start int
	needcase := true
	rewrite := func(pos int, r rune, to func(rune) rune) {
		if start < pos {
			buf.WriteString(s[start:pos])
		}
		buf.WriteRune(to(r))
		start = pos + utf8.RuneLen(r)
	}
	for pos, r := range s {
		if !unicode.IsLetter(r) {
			needcase = true
		} else if !needcase && unicode.IsUpper(r) {
			rewrite(pos, r, unicode.ToLower)
		} else {
			if needcase && unicode.IsLower(r) {
				rewrite(pos, r, unicode.ToUpper)
			}
			needcase = false
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

// Combine returns `s1` or `s2` or `s1 + sep + s2`, depending on their emptyness.
func Combine(s1 string, sep string, s2 string) string {
	if h1 := s1 != ""; h1 && s2 != "" {
		return s1 + sep + s2
	} else if h1 {
		return s1
	}
	return s2
}

func CommonPrefix(s ...string) (pref string) {
	if len(s) == 1 {
		return s[0]
	}
	if len(s) > 1 {
		var preflen int
		for cont, others, i := true, s[1:], 0; cont && i < len(s[0]); i++ {
			for j := range others {
				if i >= len(others[j]) || others[j][i] != s[0][i] {
					cont = false
					break
				}
			}
			if !cont {
				preflen = i
			}
		}
		pref = s[0][:preflen]
	}
	return
}

// Drop is a lower-level, byte-based TrimRight.
func Drop(s string, r byte) string {
	end := len(s)
	for i := end - 1; i > -1 && s[i] == r; i-- {
		end--
	}
	return s[:end]
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

// Filter returns all `strs` that satisfy `check`.
func Filter(strs []string, check func(string) bool) (filtered []string) {
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

// ForEachOccurrenceInBetween finds occurrences between two separators and
// calls `modify` for each of them, changing that occurrence in `s` to its
// return value; finally it returns `s` with all applied modifications.
//
// For example, it could be used to modify all hrefs in markdown links using simply
// the separators "](" and ")" --- `modify` would receive each inner href value.
func ForEachOccurrenceInBetween(s string, subStrStart string, subStrEnd string, modify func(string) string) string {
	var startfrom int
	for startfrom < len(s) {
		contstr := s[startfrom:]
		pos := strings.Index(contstr, subStrStart)
		if pos < 0 {
			break
		}
		pos += len(subStrStart)
		pos2 := strings.Index(contstr[pos:], subStrEnd)
		if pos2 < 0 {
			break
		}
		pos2 += pos

		culprit := contstr[pos:pos2]
		if maybe := modify(culprit); maybe != culprit {
			s = s[:startfrom] + contstr[:pos] + maybe + contstr[pos2:]
			startfrom += pos + len(maybe) + len(subStrEnd)
		} else {
			startfrom += pos2 + len(subStrEnd)
		}
	}
	return s
}

// IdxBMatching returns, for example, 3 for `("x[y]", ']', '[')` but 6 (not 5) for `("x[y[z]]", ']', '[')`.
func IdxBMatching(s string, needle byte, skipOneForEachAdditionalOccurrenceOf byte) (idx int) {
	var skipcount int
	for i := 0; i < len(s); i++ {
		if s[i] == skipOneForEachAdditionalOccurrenceOf {
			skipcount++
		} else if s[i] == needle {
			if skipcount--; skipcount == 0 {
				return i
			}
		}
	}
	return -1
}

// IdxRMatching returns, for example, 3 for `("x[y]", ']', '[')` but 6 (not 5) for `("x[y[z]]", ']', '[')`.
func IdxRMatching(s string, needle rune, skipOneForEachAdditionalOccurrenceOf rune) (idx int) {
	var skipcount int
	for i, r := range s {
		if r == skipOneForEachAdditionalOccurrenceOf {
			skipcount++
		} else if r == needle {
			if skipcount--; skipcount == 0 {
				return i
			}
		}
	}
	return -1
}

// Int64 aliases `strconv.FormatInt(i, 10)` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strconv` import.
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

// Replace allocates a one-off throw-away `strings.NewReplacer` to perform the specified replacements if `oldNewPairs` has more than 1 pair (2 elements); otherwise, calls `strings.Replace`.
func Replace(s string, oldNewPairs ...string) string {
	if l := len(oldNewPairs); l == 0 {
		return s
	} else if l == 2 {
		return strings.Replace(s, oldNewPairs[0], oldNewPairs[1], -1)
	}
	return strings.NewReplacer(oldNewPairs...).Replace(s)
}

// Rune returns `s` with the rune at `runeIndex` (not byte index) changed by `f`.
func Rune(s string, runeIndex int, f func(rune) rune) string {
	runes := []rune(s)
	runes[runeIndex] = f(runes[runeIndex])
	return string(runes)
}

// Sans returns `strs` without the specified `excludedStrs`.
func Sans(strs []string, excludedStrs ...string) []string {
	return Filter(strs, func(s string) bool {
		return !In(s, excludedStrs...)
	})
}

// Skip is a lower-level, byte-based TrimLeft.
func Skip(s string, r byte) string {
	var skip int
	for i := 0; i < len(s) && s[i] == r; i++ {
		skip++
	}
	return s[skip:]
}

// Shortest returns the `shortest` in `strs`.
func Shortest(strs []string) (shortest string) {
	for i := range strs {
		if shortest == "" || len(strs[i]) < len(shortest) {
			if shortest = strs[i]; shortest == "" {
				break
			}
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

// ToF64 returns either the `float64` denoted by `s`, or `fallback`.
func ToF64(s string, fallback float64) float64 {
	if v, err := strconv.ParseFloat(s, 64); err == nil {
		return v
	}
	return fallback
}

// ToInt returns either the `int` denoted by `s` (in base 10), or `fallback`.
func ToInt(s string, fallback int) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		return fallback
	}
	return i
}

// ToI64 returns either the `int64` denoted by `s`, or `fallback`.
func ToI64(s string, base int, fallback int64) int64 {
	if v, err := strconv.ParseInt(s, base, 64); err == nil {
		return v
	}
	return fallback
}

// ToUi64 returns either the `uint64` denoted by `s`, or `fallback`.
func ToUi64(s string, base int, fallback uint64) uint64 {
	if v, err := strconv.ParseUint(s, base, 64); err == nil {
		return v
	}
	return fallback
}

// If returns `then` if `check`, else `otherwise`.
func If(check bool, then string, otherwise string) string {
	if check {
		return then
	}
	return otherwise
}

// Until is a convenience short-hand for `BeforeFirst(s, needle, s)`.
func Until(s string, needle string) string {
	return BeforeFirst(s, needle, s)
}
