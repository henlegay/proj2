package builtins

import (
	"fmt"
	"io"
)

func Echo(w io.Writer, args ...string) error {
	_, err := fmt.Fprintln(w, args)
	return err
}
