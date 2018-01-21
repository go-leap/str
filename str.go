package ustr

import (
	"bytes"
	"fmt"
	"strings"
)

// Buf wraps `bytes.Buffer`.
type Buf struct {
	bytes.Buffer
}

func (me *Buf) Write(s string) (int, error) {
	return me.Buffer.WriteString(s)
}

func (me *Buf) Writef(s string, args ...interface{}) (int, error) {
	return me.WriteString(fmt.Sprintf(s, args...))
}

func (me *Buf) Writeln(s string) (int, error) {
	return me.WriteString(s + "\n")
}

func (me *Buf) Writelnf(s string, args ...interface{}) (int, error) {
	return me.WriteString(fmt.Sprintf(s, args...) + "\n")
}

var (
	// Fmt aliases `fmt.Sprintf` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `fmt` import.
	Fmt = fmt.Sprintf

	// Join aliases `strings.Join` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	Join = strings.Join

	// Trim aliases `strings.TrimSpace` — merely a handy short-hand during rapid iteration in non-critical code-paths that already do import `ustr` to not have to repeatedly pull in and out the extra `strings` import.
	Trim = strings.TrimSpace
)

// Split returns an empty slice if `s` is emtpy, otherwise calls `strings.Split`.
func Split(s string, sep string) (slice []string) {
	if len(s) != 0 {
		slice = strings.Split(s, sep)
	}
	return
}
