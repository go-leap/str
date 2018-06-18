package ustr

import (
	"bytes"
	"fmt"
)

// Writer is the interface that wraps the basic WriteRune and WriteString methods.
type Writer interface {
	WriteRune(rune) (int, error)
	WriteString(string) (int, error)
}

// Buf wraps `bytes.Buffer`.
type Buf struct {
	bytes.Buffer
}

func (this *Buf) Write(s string) (int, error) {
	return this.Buffer.WriteString(s)
}

func (this *Buf) Writef(s string, args ...interface{}) (int, error) {
	return this.WriteString(fmt.Sprintf(s, args...))
}

func (this *Buf) Writeln(s string) (int, error) {
	return this.WriteString(s + "\n")
}

func (this *Buf) Writelnf(s string, args ...interface{}) (int, error) {
	return this.WriteString(fmt.Sprintf(s, args...) + "\n")
}
