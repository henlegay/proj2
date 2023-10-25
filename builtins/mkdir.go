package builtins

import (
	"fmt"
	"os"
)

func MakeDirectory(args ...string) error {
	switch len(args) {
	case 1: // change to home directory if available

		err := os.Mkdir(args[0], 0700)
		if err != nil {
			return err
		}

		return err
	default:
		return fmt.Errorf("%w: expected one argument (directory)", ErrInvalidArgCount)
	}
}

