package builtins

import (
	"fmt"
	"os"
)

func Remove(args ...string) error {
	switch len(args) {
	case 0: // change to home directory if available
		return fmt.Errorf("%w: No argument found. Need at least one argument.", ErrInvalidArgCount)
	default:
		for _, v := range args {
			err := os.Remove(v)

			if err != nil {
				return err
			}
		}
		return nil
	}
}
