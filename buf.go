package ustr

import (
	"bytes"
	"fmt"
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
