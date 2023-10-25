package builtins

import (
	"fmt"
	"io"
	"os"
)

// var (
// 	ErrInvalidArgCount = errors.New("invalid argument count")
// 	HomeDir, _         = os.UserHomeDir()
// )

func ListDirectory(w io.Writer, args ...string) error {
	switch len(args) {
	case 0: // change to home directory if available

		output, err := os.ReadDir(".")
		if err != nil {
			return err
		}
		_, err = fmt.Fprintln(w, output)

		return err

	case 1:
		output, err := os.ReadDir(args[0])
		if err != nil {
			return err
		}
		_, err = fmt.Fprintln(w, output)

		return err
	default:
		return fmt.Errorf("%w: expected zero or one arguments (directory)", ErrInvalidArgCount)

	}
}
