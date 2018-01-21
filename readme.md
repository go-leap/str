# ustr
--
    import "github.com/go-leap/str"


## Usage

```go
var (
	// Fmt aliases `fmt.Sprintf` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `fmt` import.
	Fmt = fmt.Sprintf

	// Join aliases `strings.Join` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	Join = strings.Join

	// Trim aliases `strings.TrimSpace` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	Trim = strings.TrimSpace
)
```

#### func  Split

```go
func Split(s string, sep string) (slice []string)
```
Split returns an empty slice if `s` is emtpy, otherwise calls `strings.Split`.

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
