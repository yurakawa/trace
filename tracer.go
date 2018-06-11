package trace

import (
	"io"
	"fmt"
)

// tracer is a tracing events throughout code.
type Tracer struct {
	out io.Writer
}

// Trace writes the arguments to this Tracers io.Writer.
func (t *Tracer) Trace(a ...interface{}) {
	if t == nil  || t.out == nil {
		return
	}
	fmt.Fprintln(t.out, a...)
}

// New creates a new Tracer that will write the output to
// the specified io.Writer.
func New(w io.Writer) *Tracer {
	return &Tracer{out: w}
}

