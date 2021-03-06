package ustr

import (
	"fmt"
	"github.com/go-leap/std"
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

// BeforeFirstSpace returns the prefix of `s` up to the first occurrence of a `rune` satisfying `unicode.IsSpace`, or `otherwise` if no match.
func BeforeFirstSpace(s string, otherwise string) string {
	if i := strings.IndexFunc(s, unicode.IsSpace); i >= 0 {
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

// Begins returns whether the first `rune` in `s` satisfies `ok`.
func Begins(s string, ok func(rune) bool) bool {
	for _, r := range s {
		return ok(r)
	}
	return false
}

// HasAnyOf returns whether `s` contains any of the `byte`s in `anyOneOf`.
func HasAnyOf(s string, anyOneOf ...byte) bool {
	if len(s) > 0 {
		for _, b := range anyOneOf {
			if strings.IndexByte(s, b) >= 0 {
				return true
			}
		}
	}
	return false
}

var replDitchVowels = Repl("a", "", "e", "", "i", "", "o", "", "u", "", "ä", "", "ö", "", "ü", "", "ï", "", "ø", "", "æ", "")

// Similes for when Levenshtein seems overkill.. cheap & naive but handy
func Similes(s string, candidates ...string) (sim []string) {
	if len(candidates) == 0 {
		return
	}
	sl := replDitchVowels.Replace(Lo(s))
	for _, c := range candidates {
		if len(c) >= len(s)-2 && len(c) <= len(s)+2 {
			if cl := replDitchVowels.Replace(Lo(c)); Has(sl, cl) || Has(cl, sl) {
				sim = append(sim, c)
			}
		}
	}
	return
}

// IsLen1And returns whether `s` is equal to any of the `byte`s in `anyOneOf`.
func IsLen1And(s string, anyOneOf ...byte) bool {
	if len(s) == 1 {
		for _, b := range anyOneOf {
			if s[0] == b {
				return true
			}
		}
	}
	return false
}

// HasAny returns whether any `rune` in `s` satisfies `ok`.
func HasAny(s string, ok func(rune) bool) bool {
	for _, r := range s {
		if ok(r) {
			return true
		}
	}
	return false
}

// BeginsAndContainsOnly returns whether the first `rune` in `s` satisfies `begins` and all `rune`s in `s` satisfy all predicates in `containsOnly`.
func BeginsAndContainsOnly(s string, begins func(rune) bool, containsOnly ...func(rune) bool) bool {
	for i, r := range s {
		if i == 0 && !begins(r) {
			return false
		}
		for _, ok := range containsOnly {
			if !ok(r) {
				return false
			}
		}
	}
	return len(s) > 0
}

// BeginsLetter returns whether the first rune in `s` satisfies `unicode.IsLetter`.
func BeginsLetter(s string) bool {
	for _, r := range s {
		return unicode.IsLetter(r)
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

// CaseLo returns `s` with the rune at `runeIndex` (not byte index) guaranteed to be lower-case.
func CaseLo(s string, runeIndex int) string { return Rune(s, runeIndex, unicode.ToLower) }

// CaseUp returns `s` with the rune at `runeIndex` (not byte index) guaranteed to be upper-case.
func CaseUp(s string, runeIndex int) string { return Rune(s, runeIndex, unicode.ToUpper) }

// CaseSnake returns a snake-cased attempt at `s` (based on `unicode.IsLetter`).
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

// Combine returns `s1` or `s2` or `s1 + sep + s2`, depending on their emptiness.
func Combine(s1 string, sep string, s2 string) string {
	if h1 := s1 != ""; h1 && s2 != "" {
		return s1 + sep + s2
	} else if h1 {
		return s1
	}
	return s2
}

// CountPrefixRunes returns how many occurrences of `prefix` are leading in `s`.
func CountPrefixRunes(s string, prefix rune) (n int) {
	for _, r := range s {
		if r != prefix {
			break
		} else {
			n++
		}
	}
	return
}

// CommonPrefix finds the prefix `pref` that all values in `s` share, if any.
func CommonPrefix(s ...string) (pref string) {
	if len(s) == 1 {
		return s[0]
	} else if len(s) > 1 {
		var preflen int
		for cont, others, i := true, s[1:], 0; cont && i < len(s[0]); i++ {
			for j := range others {
				if len(others[j]) < i || others[j][i] != s[0][i] {
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

func Index(strs []string, check func(string) bool) int {
	for i, s := range strs {
		if check(s) {
			return i
		}
	}
	return -1
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

// JoinB is like `strings.Join` but with a byte-length char as separator.
func JoinB(s []string, b byte) string {
	if len(s) == 0 {
		return ""
	} else if len(s) == 1 {
		return s[0]
	}
	buf := Buf{BytesWriter: ustd.BytesWriter{Data: make([]byte, 0, len(s)*len(s[0]))}}
	for i := range s {
		buf.WriteByte(b)
		buf.WriteString(s[i])
	}
	buf.Data = buf.Data[1:]
	return buf.String()
}

// Uint64s joins together the hex-formatted `values`.
func Uint64s(joinBy byte, values []uint64) string {
	var b ustd.BytesWriter
	for i := range values {
		b.WriteByte(joinBy)
		b.WriteString(strconv.FormatUint(values[i], 16))
	}
	return string(b.Data[1:])
}

// HasOneOf returns whether `s` contains any of the specified `subStrings`.
func HasOneOf(s string, subStrings ...string) bool {
	return FirstIn(s, subStrings...) != ""
}

// In returns whether `strs` contains `s`.
func In(s string, strs ...string) bool {
	if len(strs) == 1 { // occurs often enough indeed
		return s == strs[0]
	}
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

// IsRepeat returns whether `s` contains nothing but one-or-more occurrences of `first`. If `first` is `0`, it is initialized from the first `rune` in `s`.
func IsRepeat(s string, first rune) bool {
	if len(s) > 0 {
		for i, r := range s {
			if i == 0 && first == 0 {
				first = r
			} else if r != first {
				return false
			}
		}
		return true
	}
	return false
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

// Plu returns (if `s` is, say, `"foo"`) "1 foo" or "0 foos" or "2 foos", so appends "s" to `s` unless `n` is 1. The pluralizer is English-language-oriented and covers no corner-cases such as "bus" and the likes, but for simple command-line programs it's cheap.
func Plu(n int, s string) (r string) {
	r = Int(n) + " " + s
	if n != 1 {
		r += "s"
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

// Merge returns a slice with the items of `s` and `with`, no duplicates, no guaranteed ordering. If `dropIf` is given, it is called to prevent values from being included in the return slice.
func Merge(s []string, with []string, dropIf func(string) bool) []string {
	if dropIf != nil {
		for i := 0; i < len(with); i++ {
			if dropIf(with[i]) {
				with = append(with[:i], with[i+1:]...)
				i--
			}
		}
	}
	for i := range s {
		if dropIf == nil || !dropIf(s[i]) {
			var have bool
			for j := range with {
				if have = with[j] == s[i]; have {
					break
				}
			}
			if !have {
				with = append(with, s[i])
			}
		}
	}
	return with
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

// ShortestAndLongest returns the length of the shortest item in `s`, as well as the length of the longest. Both will return as `-1` if `s` is empty.
func ShortestAndLongest(s ...string) (lenShortest int, lenLongest int) {
	lenShortest, lenLongest = -1, -1
	isfirst := true
	for i := range s {
		if l := len(s[i]); isfirst {
			lenShortest, lenLongest, isfirst = l, l, false
		} else {
			if l > lenLongest {
				lenLongest = l
			}
			if l < lenShortest {
				lenShortest = l
			}
		}
	}
	return
}

// Split returns an empty slice if `s` is emtpy, otherwise calls `strings.Split`.
func Split(s string, sep string) (splits []string) {
	if len(s) != 0 {
		splits = strings.Split(s, sep)
	}
	return
}

// SplitB returns an empty slice if `s` is emtpy, otherwise it's like `Split` but with a `byte` separator.
func SplitB(s string, sep byte, initialCap int) (splits []string) {
	if initialCap > 0 {
		splits = make([]string, 0, initialCap)
	}
	var startfrom int
	for i := 0; i < len(s); i++ {
		if s[i] == sep {
			splits = append(splits, s[startfrom:i])
			startfrom = i + 1
		}
	}
	if startfrom > 0 {
		splits = append(splits, s[startfrom:])
	} else if len(s) > 0 {
		splits = []string{s}
	}
	return
}

// SplitR returns an empty slice if `s` is emtpy, otherwise it's like `Split` but with a `rune` separator.
func SplitR(s string, sep rune, initialCap int) (splits []string) {
	if initialCap > 0 {
		splits = make([]string, 0, initialCap)
	}
	var startfrom int
	var lastwasmatch bool
	for i, r := range s {
		if lastwasmatch {
			startfrom = i
		}
		if lastwasmatch = (r == sep); lastwasmatch {
			splits = append(splits, s[startfrom:i])
		}
	}
	if startfrom > 0 {
		splits = append(splits, s[startfrom:])
	} else if len(s) > 0 {
		splits = append(splits, s)
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

// Times converts `Repeat` to `string`.
func Times(s string, n int) string {
	return string(Repeat(s, n))
}

// Repeat is a somewhat leaner version of `strings.Repeat`.
func Repeat(s string, n int) (str []byte) {
	if len(s) == 1 {
		str = make([]byte, n)
		for i := range str {
			str[i] = s[0]
		}
	} else if len(s) != 0 {
		str = make([]byte, 0, len(s)*n)
		for i := 0; i < n; i++ {
			str = append(str, s...)
		}
	}
	return
}

// RepeatB is like `Repeat` but a single byte-length char.
func RepeatB(b byte, n int) (s []byte) {
	s = make([]byte, n)
	for i := range s {
		s[i] = b
	}
	return
}
