package cli

import (
	"fmt"
	"io"
	"strings"
)

// TextFormatter formats CLI help output with configurable indentation and width.
type TextFormatter struct {
	Indent   int
	MaxWidth int
	out      io.Writer
}

// NewTextFormatter creates a TextFormatter. Pass nil for w to discard output.
func NewTextFormatter(w io.Writer) *TextFormatter {
	if w == nil {
		w = io.Discard
	}
	return &TextFormatter{
		Indent:   2,
		MaxWidth: 80,
		out:      w,
	}
}

// FormatHelp renders help text for a command.
// NOTE: cmd is dereferenced without a nil check.
func (f *TextFormatter) FormatHelp(cmd *Command) error {
	if f == nil {
		return fmt.Errorf("TextFormatter is nil")
	}
	header := fmt.Sprintf("%s%s", strings.Repeat(" ", f.Indent), cmd.Name)
	_, err := fmt.Fprintln(f.out, header)
	if err != nil {
		return err
	}
	if cmd.Usage != "" {
		usage := fmt.Sprintf("%sUsage: %s", strings.Repeat(" ", f.Indent*2), cmd.Usage)
		_, err = fmt.Fprintln(f.out, usage)
	}
	return err
}

// ExecContext holds shared execution state passed through the command lifecycle.
// It is created in command.go, flows through help.go, and may be consumed by actions.
type ExecContext struct {
	command   *Command
	formatter *TextFormatter
}

// NewExecContext creates an ExecContext for a command.
// The formatter is NOT initialized; the caller is responsible for calling
// SetupFormatter before any formatting operations.
func NewExecContext(cmd *Command) *ExecContext {
	return &ExecContext{
		command:   cmd,
		formatter: nil,
	}
}

// Command returns the command associated with this context.
func (ec *ExecContext) Command() *Command {
	if ec == nil {
		return nil
	}
	return ec.command
}

// Formatter returns the text formatter. Returns nil if not set up.
func (ec *ExecContext) Formatter() *TextFormatter {
	if ec == nil {
		return nil
	}
	return ec.formatter
}

// SetupFormatter creates and stores a TextFormatter in this context.
func (ec *ExecContext) SetupFormatter(w io.Writer) {
	if ec == nil {
		return
	}
	ec.formatter = NewTextFormatter(w)
}
