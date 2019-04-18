package ustr

import (
	"fmt"

	"github.com/go-leap/std"
)

// Writer is the interface that wraps the basic WriteRune and WriteString methods.
type Writer interface {
	WriteRune(rune) (int, error)
	WriteString(string) (int, error)
}

// Buf wraps `bytes.Buffer`.
type Buf struct {
	ustd.BytesWriter
}

func (me *Buf) Write(s string) {
	me.WriteString(s)
}

func (me *Buf) Writef(s string, args ...interface{}) {
	me.WriteString(fmt.Sprintf(s, args...))
}

func (me *Buf) Writeln(s string) {
	me.WriteString(s + "\n")
}

func (me *Buf) Writelnf(s string, args ...interface{}) {
	me.WriteString(fmt.Sprintf(s, args...) + "\n")
}

func (me *Buf) String() string { return string(me.Data) }
