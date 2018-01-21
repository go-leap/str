# ustr
--
    import "github.com/go-leap/str"


## Usage

```go
var (
	// Fmt aliases `fmt.Sprintf` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `fmt` import.
	Fmt = fmt.Sprintf

	// Has aliases `strings.Contains` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	Has = strings.Contains

	// Int aliases `strconv.Itoa` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strconv` import.
	Int = strconv.Itoa

	// Join aliases `strings.Join` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	Join = strings.Join

	// Lo aliases `strings.ToLower` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	Lo = strings.ToLower

	// Pos aliases `strings.Index` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	Pos = strings.Index

	// Pref aliases `strings.HasPrefix` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	Pref = strings.HasPrefix

	// Repl aliases `strings.NewReplacer` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	Repl = strings.NewReplacer

	// Suff aliases `strings.HasSuffix` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	Suff = strings.HasSuffix

	// Times aliases `strings.Repeat` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	Times = strings.Repeat

	// Trim aliases `strings.TrimSpace` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	Trim = strings.TrimSpace

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

#### func  BreakOnFirst

```go
func BreakOnFirst(s string, needle string) (prefix string, suffix string)
```
BreakOnFirst returns the prefix and suffix next to the first `needle`
encountered in `s`. (If no match, `prefix` is `""` and `suffix` will be `s`.)

#### func  BreakOnLast

```go
func BreakOnLast(s string, needle string) (prefix string, suffix string)
```
BreakOnLast returns the prefix and suffix next to the last `needle` encountered
in `s`. (If no match, `prefix` is `""` and `suffix` will be `s`.)

#### func  Combine

```go
func Combine(s1 string, sep string, s2 string) string
```
Combine returns `s1` or `s2` or `s1 + sep + s2`, depending on their emptyness.

#### func  EnsureCase

```go
func EnsureCase(s string, runeIndex int, upper bool) string
```
EnsureCase returns `s` with the rune at `runeIndex` (not byte index) guaranteed
to be upper-case if `upper`, or lower-case if not.

#### func  Fewest

```go
func Fewest(strs []string, substr string, otherwise func([]string) string) (s string)
```
Fewest returns the `s` in `strs` with the lowest `strings.Count` of `substr`. If
the count is identical for all, it returns `otherwise(strs)` (if supplied).

#### func  Filtered

```go
func Filtered(strs []string, check func(string) bool) (filtered []string)
```
Filtered returns all `strs` that satisfy `check`.

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

#### func  Has1Of

```go
func Has1Of(s string, subStrings ...string) bool
```
Has1Of returns whether `s` contains any of the specified `subStrings`.

#### func  In

```go
func In(s string, strs ...string) bool
```
In returns whether `strs` contains `s`.

#### func  IsLower

```go
func IsLower(s string) bool
```
IsLower returns whether all `unicode.IsLetter` runes in `s` satisfy
`unicode.IsLower`.

#### func  IsUpper

```go
func IsUpper(s string) bool
```
IsUpper returns whether all `unicode.IsLetter` runes in `s` satisfy
`unicode.IsUpper`.

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

#### func  Pref1Of

```go
func Pref1Of(s string, prefixes ...string) string
```
Pref1Of returns the first of the specified (non-empty) `prefixes` that `s`
begins with, or `""`.

#### func  Replace

```go
func Replace(s string, oldNewPairs ...string) string
```
Replace allocates a one-off throw-away `strings.NewReplacer` to perform the
specified replacements.

#### func  Shortest

```go
func Shortest(strs []string) (s string)
```
Shortest returns the shortest `s` in `strs`.

#### func  Split

```go
func Split(s string, sep string) (strs []string)
```
Split returns an empty slice if `s` is emtpy, otherwise calls `strings.Split`.

#### func  SplitByWhitespaceAndReJoinBySpace

```go
func SplitByWhitespaceAndReJoinBySpace(s string) string
```
SplitByWhitespaceAndJoin returns `s` with all occurrences of multiple subsequent
`unicode.IsSpace` runes in a row collapsed into one single white-space (`" "`)
rune.

#### func  ToBool

```go
func ToBool(s string, fallback bool) bool
```
ToBool returns either the `bool` denoted by `s`, or `fallback`.

#### func  ToFloat

```go
func ToFloat(s string, fallback float64) float64
```
ToFloat returns either the `float64` denoted by `s`, or `fallback`.

#### func  ToInt

```go
func ToInt(s string, fallback int64) int64
```
ToInt returns either the `int64` denoted by `s`, or `fallback`.

#### func  ToUint

```go
func ToUint(s string, fallback uint64) uint64
```
ToUint returns either the `uint64` denoted by `s`, or `fallback`.

#### func  Without

```go
func Without(strs []string, withoutVals ...string) []string
```
Without returns `strs` sans all specified `withoutVals`.

#### type Buf

```go
type Buf struct {
	bytes.Buffer
}
```

Buf wraps `bytes.Buffer`.

#### func (*Buf) Write

```go
func (me *Buf) Write(s string) (int, error)
```

#### func (*Buf) Writef

```go
func (me *Buf) Writef(s string, args ...interface{}) (int, error)
```

#### func (*Buf) Writeln

```go
func (me *Buf) Writeln(s string) (int, error)
```

#### func (*Buf) Writelnf

```go
func (me *Buf) Writelnf(s string, args ...interface{}) (int, error)
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
Matches returns whether the `s` matches `me`, which could:

- begin and end with an asterisk `*` wildcard: "contains" semantics

- only begin with an asterisk `*` wildcard: "endsWith" semantics

- only end with an asterisk `*` wildcard: "beginsWith" semantics

- only consist of an asterisk `*` wildcard: returns `true` always

- otherwise: returns whether `s == me`.

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
