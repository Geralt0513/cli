package cli

import (
	"fmt"
	"io"
)

// OutputFormatter formats help text for commands with configurable indentation.
// It is designed to be created via NewFormatter and used by ShowFormattedHelp
// in help.go.
type OutputFormatter struct {
	Indent int
	out    io.Writer
}

// NewFormatter creates a new OutputFormatter.
// Returns nil if w is nil - the caller must check.
func NewFormatter(w io.Writer, indent int) *OutputFormatter {
	if w == nil {
		return nil // Intentional: returns nil when writer is nil
	}
	return &OutputFormatter{
		Indent: indent,
		out:    w,
	}
}

// FormatHelp formats the help text for a command.
// PANICS if f is nil - caller must nil-check after NewFormatter.
func (f *OutputFormatter) FormatHelp(cmd *Command) error {
	header := fmt.Sprintf("%*s%s", f.Indent, "", cmd.Name)
	_, err := fmt.Fprintln(f.out, header)
	return err
}
