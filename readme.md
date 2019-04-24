# ustr
--
    import "github.com/go-leap/str"


## Usage

```go
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
```

#### func  AfterFirst

```go
func AfterFirst(s string, needle string, otherwise string) string
```
AfterFirst returns the suffix of `s` beginning right after the first occurrence
of `needle`, or `otherwise` if no match.

#### func  AfterLast

```go
func AfterLast(s string, needle string, otherwise string) string
```
AfterLast returns the suffix of `s` beginning right after the last occurrence of
`needle`, or `otherwise` if no match.

#### func  BeforeFirst

```go
func BeforeFirst(s string, needle string, otherwise string) string
```
BeforeFirst returns the prefix of `s` up to the first occurrence of `needle`, or
`otherwise` if no match.

#### func  BeforeFirstSpace

```go
func BeforeFirstSpace(s string, otherwise string) string
```
BeforeFirstSpace returns the prefix of `s` up to the first occurrence of a
`rune` satisfying `unicode.IsSpace`, or `otherwise` if no match.

#### func  BeforeLast

```go
func BeforeLast(s string, needle string, otherwise string) string
```
BeforeLast returns the prefix of `s` up to the last occurrence of `needle`, or
`otherwise` if no match.

#### func  BeginsLower

```go
func BeginsLower(s string) bool
```
BeginsLower returns whether the first rune in `s` satisfies both
`unicode.IsLetter` and `unicode.IsLower`.

#### func  BeginsUpper

```go
func BeginsUpper(s string) bool
```
BeginsUpper returns whether the first rune in `s` satisfies both
`unicode.IsLetter` and `unicode.IsUpper`.

#### func  BreakOnFirstOrPref

```go
func BreakOnFirstOrPref(s string, needle string) (prefix string, suffix string)
```
BreakOnFirstOrPref returns the prefix and suffix next to the first `needle`
encountered in `s`. (If no match, `suffix` is `""` and `prefix` will be `s`.)

#### func  BreakOnFirstOrSuff

```go
func BreakOnFirstOrSuff(s string, needle string) (prefix string, suffix string)
```
BreakOnFirstOrSuff returns the prefix and suffix next to the first `needle`
encountered in `s`. (If no match, `prefix` is `""` and `suffix` will be `s`.)

#### func  BreakOnLast

```go
func BreakOnLast(s string, needle string) (prefix string, suffix string)
```
BreakOnLast returns the prefix and suffix next to the last `needle` encountered
in `s`. (If no match, `prefix` is `""` and `suffix` will be `s`.)

#### func  Case

```go
func Case(s string, runeIndex int, upper bool) string
```
Case returns `s` with the rune at `runeIndex` (not byte index) guaranteed to be
upper-case if `upper`, or lower-case if not.

#### func  CaseLo

```go
func CaseLo(s string, runeIndex int) string
```
Case returns `s` with the rune at `runeIndex` (not byte index) guaranteed to be
lower-case.

#### func  CaseSnake

```go
func CaseSnake(s string) string
```

#### func  CaseUp

```go
func CaseUp(s string, runeIndex int) string
```
Case returns `s` with the rune at `runeIndex` (not byte index) guaranteed to be
upper-case.

#### func  Combine

```go
func Combine(s1 string, sep string, s2 string) string
```
Combine returns `s1` or `s2` or `s1 + sep + s2`, depending on their emptyness.

#### func  CommonPrefix

```go
func CommonPrefix(s ...string) (pref string)
```

#### func  Drop

```go
func Drop(s string, r byte) string
```
Drop is a lower-level, byte-based TrimRight.

#### func  Fewest

```go
func Fewest(strs []string, substr string, otherwise func([]string) string) (s string)
```
Fewest returns the `s` in `strs` with the lowest `strings.Count` of `substr`. If
the count is identical for all, it returns `otherwise(strs)` (if supplied).

#### func  Filter

```go
func Filter(strs []string, check func(string) bool) (filtered []string)
```
Filter returns all `strs` that satisfy `check`.

#### func  FirstIn

```go
func FirstIn(s string, subStrings ...string) string
```
FirstIn returns the first in `subStrings` to satisfy `strings.Contains(s,
substr)`, or `""`.

#### func  FirstOf

```go
func FirstOf(strs ...string) (s string)
```
FirstOf returns the first non-empty `s` encountered in `strs`.

#### func  ForEachOccurrenceInBetween

```go
func ForEachOccurrenceInBetween(s string, subStrStart string, subStrEnd string, modify func(string) string) string
```
ForEachOccurrenceInBetween finds occurrences between two separators and calls
`modify` for each of them, changing that occurrence in `s` to its return value;
finally it returns `s` with all applied modifications.

For example, it could be used to modify all hrefs in markdown links using simply
the separators "](" and ")" --- `modify` would receive each inner href value.

#### func  Has1Of

```go
func Has1Of(s string, subStrings ...string) bool
```
Has1Of returns whether `s` contains any of the specified `subStrings`.

#### func  IdxBMatching

```go
func IdxBMatching(s string, needle byte, skipOneForEachAdditionalOccurrenceOf byte) (idx int)
```
IdxBMatching returns, for example, 3 for `("x[y]", ']', '[')` but 6 (not 5) for
`("x[y[z]]", ']', '[')`.

#### func  IdxRMatching

```go
func IdxRMatching(s string, needle rune, skipOneForEachAdditionalOccurrenceOf rune) (idx int)
```
IdxRMatching returns, for example, 3 for `("x[y]", ']', '[')` but 6 (not 5) for
`("x[y[z]]", ']', '[')`.

#### func  If

```go
func If(check bool, then string, otherwise string) string
```
If returns `then` if `check`, else `otherwise`.

#### func  In

```go
func In(s string, strs ...string) bool
```
In returns whether `strs` contains `s`.

#### func  Int64

```go
func Int64(i int64) string
```
Int64 aliases `strconv.FormatInt(i, 10)` — merely a handy short-hand during
rapid iteration in non-critical code-paths that already do import `ustr` to not
have to repeatedly pull in and out the extra `strconv` import.

#### func  IsLower

```go
func IsLower(s string) bool
```
IsLower returns whether all `unicode.IsLetter` runes in `s` satisfy
`unicode.IsLower`.

#### func  IsRepeat

```go
func IsRepeat(s string) bool
```

#### func  IsUpper

```go
func IsUpper(s string) bool
```
IsUpper returns whether all `unicode.IsLetter` runes in `s` satisfy
`unicode.IsUpper`.

#### func  JoinB

```go
func JoinB(nonEmpty []string, b byte) string
```

#### func  Longest

```go
func Longest(strs []string) (s string)
```
Longest returns the longest `s` in `strs`.

#### func  Map

```go
func Map(strs []string, f func(string) string) (items []string)
```
Map applies `f` to each `string` in `strs` and returns the results in `items`.

#### func  Merge

```go
func Merge(s []string, with []string, dropIf func(string) bool) []string
```

#### func  NamedPlaceholders

```go
func NamedPlaceholders(begin byte, end byte) func(string, ...string) string
```
NamedPlaceholders is an occasionally-preferable alternative to `fmt.Sprintf` or
`strings.Replace` / `strings.Replacer` for (fully `string`ly-typed)
"micro-templating":

```go

    repl := ustr.NamedPlaceHolders('(', ']')
    hi := repl("Hello (name]!", "name", "world")

```

The delimiter `begin` and `end` chars may well be equal, if so desired. When
called, the returned `func` traverses its `string` arg exactly once, ie. it does
not re-process the replacements or its final result. It searches through its
name-value-pairs once per fully-delimited-substring. Any of those occurrences
not found in its name-value-pairs are left in-place including the delimiters.

#### func  Plu

```go
func Plu(n int, s string) (r string)
```

#### func  Pref1Of

```go
func Pref1Of(s string, prefixes ...string) string
```
Pref1Of returns the first of the specified (non-empty) `prefixes` that `s`
begins with, or `""`.

#### func  Repeat

```go
func Repeat(s string, n int) (str []byte)
```

#### func  ReplB

```go
func ReplB(s string, oldNew ...byte) string
```
ReplB replaces individual `byte`s in `s` based on the given old-new pairs:
because for some few `strings.Replace` / `strings.Replacer` scenarios, nothing
more is needed.

#### func  Replace

```go
func Replace(s string, oldNewPairs ...string) string
```
Replace allocates a one-off throw-away `strings.NewReplacer` to perform the
specified replacements if `oldNewPairs` has more than 1 pair (2 elements);
otherwise, calls `strings.Replace`.

#### func  Rune

```go
func Rune(s string, runeIndex int, f func(rune) rune) string
```
Rune returns `s` with the rune at `runeIndex` (not byte index) changed by `f`.

#### func  Sans

```go
func Sans(strs []string, excludedStrs ...string) []string
```
Sans returns `strs` without the specified `excludedStrs`.

#### func  Shortest

```go
func Shortest(strs []string) (shortest string)
```
Shortest returns the `shortest` in `strs`.

#### func  ShortestAndLongest

```go
func ShortestAndLongest(s ...string) (lenShortest int, lenLongest int)
```

#### func  Skip

```go
func Skip(s string, r byte) string
```
Skip is a lower-level, byte-based TrimLeft.

#### func  Split

```go
func Split(s string, sep string) (splits []string)
```
Split returns an empty slice if `s` is emtpy, otherwise calls `strings.Split`.

#### func  SplitB

```go
func SplitB(s string, sep byte, initialCap int) (splits []string)
```

#### func  SplitByWhitespaceAndReJoinBySpace

```go
func SplitByWhitespaceAndReJoinBySpace(s string) string
```
SplitByWhitespaceAndJoin returns `s` with all occurrences of multiple subsequent
`unicode.IsSpace` runes in a row collapsed into one single white-space (`" "`)
rune.

#### func  SplitR

```go
func SplitR(s string, sep rune, initialCap int) (splits []string)
```

#### func  Times

```go
func Times(s string, n int) string
```

#### func  ToBool

```go
func ToBool(s string, fallback bool) bool
```
ToBool returns either the `bool` denoted by `s`, or `fallback`.

#### func  ToF64

```go
func ToF64(s string, fallback float64) float64
```
ToF64 returns either the `float64` denoted by `s`, or `fallback`.

#### func  ToI64

```go
func ToI64(s string, base int, fallback int64) int64
```
ToI64 returns either the `int64` denoted by `s`, or `fallback`.

#### func  ToInt

```go
func ToInt(s string, fallback int) int
```
ToInt returns either the `int` denoted by `s` (in base 10), or `fallback`.

#### func  ToUi64

```go
func ToUi64(s string, base int, fallback uint64) uint64
```
ToUi64 returns either the `uint64` denoted by `s`, or `fallback`.

#### func  Uint64s

```go
func Uint64s(joinBy byte, values []uint64) string
```

#### func  Until

```go
func Until(s string, needle string) string
```
Until is a convenience short-hand for `BeforeFirst(s, needle, s)`.

#### type Buf

```go
type Buf struct {
	ustd.BytesWriter
}
```

Buf wraps `bytes.Buffer`.

#### func (*Buf) String

```go
func (me *Buf) String() string
```

#### func (*Buf) Write

```go
func (me *Buf) Write(s string)
```

#### func (*Buf) Writef

```go
func (me *Buf) Writef(s string, args ...interface{})
```

#### func (*Buf) Writeln

```go
func (me *Buf) Writeln(s string)
```

#### func (*Buf) Writelnf

```go
func (me *Buf) Writelnf(s string, args ...interface{})
```

#### type Pat

```go
type Pat string
```

Pat is a most-simplistic, overly-rudimentary, simplest-of-simpletons string
pattern matcher. It allows a single asterisk `*` wild-card at its beginning, its
end, or both, as described in `Pat.Match`. This covers a bafflingly substantial
amount of real-world use-cases — if more is needed, Go's `path.Match`, reg-exps
etc. will deliver instead.

#### func (Pat) AllMatch

```go
func (me Pat) AllMatch(strs ...string) bool
```
AllMatch returns whether all the specified `strs` satisfy `me.Match`.

#### func (Pat) FirstMatch

```go
func (me Pat) FirstMatch(strs ...string) string
```
FirstMatch returns the first in `strs` that `me.Match`es, or `""`.

#### func (Pat) Match

```go
func (me Pat) Match(s string) bool
```
Match returns whether `s` matches `me`, which could:

- begin and end with an asterisk `*` wildcard: "contains" semantics

- only begin with an asterisk `*` wildcard: "endsWith" semantics

- only end with an asterisk `*` wildcard: "beginsWith" semantics

- only consist of an asterisk `*` wildcard: always matches any `s`

- otherwise: matches if `s == me`.

#### type Pats

```go
type Pats []Pat
```

Pats is a slice of `Pat`s.

#### func (*Pats) Add

```go
func (me *Pats) Add(pats ...Pat)
```
Add `append`s all the specified `pats` to `me`.

#### func (Pats) FirstMatch

```go
func (me Pats) FirstMatch(s string) Pat
```
FirstMatch returns the first `Pat` in `me` to `Match(s)`, or `""`.

#### func (Pats) NoMatch

```go
func (me Pats) NoMatch(s string) bool
```
NoMatch returns whether not a single `Pat` in `me` does `Match(s)`.

#### type Writer

```go
type Writer interface {
	WriteRune(rune) (int, error)
	WriteString(string) (int, error)
}
```

Writer is the interface that wraps the basic WriteRune and WriteString methods.
