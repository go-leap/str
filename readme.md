# ustr
--
    import "github.com/go-leap/str"


## Usage

```go
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

#### func  FirstOf

```go
func FirstOf(strs ...string) (s string)
```
FirstOf returns the first non-empty `s` encountered in `strs`.

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

#### func  Map

```go
func Map(strs []string, f func(string) string) (items []string)
```
Map applies `f` to each `string` in `strs` and returns the results in `items`.

#### func  Split

```go
func Split(s string, sep string) (slice []string)
```
Split returns an empty slice if `s` is emtpy, otherwise calls `strings.Split`.

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
