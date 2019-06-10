package ustr

import (
	"fmt"

	"github.com/go-leap/std"
)

// type Writer interface {
// 	WriteRune(rune) (int, error)
// 	WriteString(string) (int, error)
// }

// Buf used to wrap `bytes.Buffer`, now wraps `ustd.BytesWriter`.
type Buf struct {
	ustd.BytesWriter
}

// Write is equivalent to / short-hand for `me.BytesWriter.WriteString(s)`.
func (me *Buf) Write(s string) {
	me.WriteString(s)
}

// Writef uses `fmt.Sprintf`.
func (me *Buf) Writef(s string, args ...interface{}) {
	me.WriteString(fmt.Sprintf(s, args...))
}

// Writeln appends `\n` to `s` then writes.
func (me *Buf) Writeln(s string) {
	me.WriteString(s + "\n")
}

// Writelnf uses `fmt.Sprintf`.
func (me *Buf) Writelnf(s string, args ...interface{}) {
	me.WriteString(fmt.Sprintf(s, args...) + "\n")
}

// String returns `me.Data` converted to `string`.
func (me *Buf) String() string { return string(me.Data) }
