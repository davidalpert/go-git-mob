// Package plaintext implements a development-friendly plain textual handler.
package plaintext

import (
	"fmt"
	"github.com/apex/log"
	"io"
	"os"
	"sync"
	"time"
)

// Default handler outputting to stderr.
var Default = New(os.Stderr)

// start time.
var start = time.Now()

// Strings mapping.
var Strings = [...]string{
	log.DebugLevel: "DEBUG",
	log.InfoLevel:  "INFO",
	log.WarnLevel:  "WARN",
	log.ErrorLevel: "ERROR",
	log.FatalLevel: "FATAL",
}

// Handler implementation.
type Handler struct {
	mu     sync.Mutex
	Writer io.Writer
}

// New handler.
func New(w io.Writer) *Handler {
	return &Handler{
		Writer: w,
	}
}

// HandleLog implements log.Handler.
func (h *Handler) HandleLog(e *log.Entry) error {
	level := Strings[e.Level]
	names := e.Fields.Names()

	h.mu.Lock()
	defer h.mu.Unlock()

	ts := time.Since(start) / time.Second
	fmt.Fprintf(h.Writer, "%6s [%04d] %-25s\n", level, ts, e.Message)

	for _, name := range names {
		fmt.Fprintf(h.Writer, "       > %s = %v\n", name, e.Fields.Get(name))
	}

	fmt.Fprintln(h.Writer)

	return nil
}
