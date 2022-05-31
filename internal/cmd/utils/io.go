package utils

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// IOStreams provides the standard names for iostreams.  This is useful for embedding and for unit testing.
// Inconsistent and different names make it hard to read and review code
type IOStreams struct {
	// In think, os.Stdin
	In io.Reader
	// Out think, os.Stdout
	Out io.Writer
	// ErrOut think, os.Stderr
	ErrOut io.Writer
}

// DefaultOSStreams returns a valid IOStreams pointing at os.Stdin, os.Stdout, os.Stderr
func DefaultOSStreams() IOStreams {

	return IOStreams{
		In:     os.Stdin,
		Out:    os.Stdout,
		ErrOut: os.Stderr,
	}
}

// NewTestIOStreams returns a valid IOStreams and in, out, errout buffers for unit tests
func NewTestIOStreams() (IOStreams, *bytes.Buffer, *bytes.Buffer, *bytes.Buffer) {
	in := &bytes.Buffer{}
	out := &bytes.Buffer{}
	errOut := &bytes.Buffer{}

	return IOStreams{
		In:     in,
		Out:    out,
		ErrOut: errOut,
	}, in, out, errOut
}

// WriteNewline writes a newline to o.Out
func (ios IOStreams) WriteNewline() error {
	return ios.WriteString("\n")
}

// WriteString writes a string to o.Out
func (ios IOStreams) WriteString(s string) error {
	_, err := ios.WriteBytes([]byte(s))
	return err
}

// WriteStringf writes a string to o.Out
func (ios IOStreams) WriteStringf(format string, a ...interface{}) error {
	return ios.WriteString(fmt.Sprintf(format, a...))
}

// WriteStringln writes a string to o.Out with a newline
func (ios IOStreams) WriteStringln(a ...interface{}) error {
	return ios.WriteString(fmt.Sprintln(a...))
}

// WriteBytes writes a byte array to o.Out
func (ios IOStreams) WriteBytes(b []byte) (int, error) {
	return ios.Out.Write(b)
}

// WriteBytesln writes a byte array to o.Out
func (ios IOStreams) WriteBytesln(b []byte) {
	ios.Out.Write(b)
	ios.WriteNewline()
}

// WriteErrorNewline writes a newline to o.ErrOut
func (ios IOStreams) WriteErrorNewline() error {
	return ios.WriteErrorString("\n")
}

// WriteErrorString writes a string to o.ErrOut
func (ios IOStreams) WriteErrorString(s string) error {
	_, err := ios.WriteErrorBytes([]byte(s))
	return err
}

// WriteErrorBytes writes a string to o.ErrOut
func (ios IOStreams) WriteErrorBytes(b []byte) (int, error) {
	return ios.ErrOut.Write(b)
}
