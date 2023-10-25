package builtins

import (
	"fmt"
	"io"
	"os"
)

func PrintWorkingDirectory(w io.Writer, args ...string) error {
	switch len(args) {
	case 0: // change to home directory if available
		HomeDir, _ = os.UserHomeDir()
		_, err := fmt.Fprintln(w, HomeDir)

		return err
	default:
		return fmt.Errorf("%w: expected zero arguments (directory)", ErrInvalidArgCount)
	}
}
