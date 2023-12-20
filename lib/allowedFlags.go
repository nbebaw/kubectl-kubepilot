package lib

import (
	"flag"
	"fmt"
)

// Check Allowed Flags
/*
 the function get allowedflags as map string and return the error
 if the user start the command with illagel argument
 {allowedflags} -> map[string]bool
*/
func CheckAllowedFlags(allowedflags map[string]bool) error {
	invalidFlags := false

	flag.Visit(func(f *flag.Flag) {
		if _, exists := allowedflags[f.Name]; !exists {
			invalidFlags = true
		}
	})

	if invalidFlags {
		return fmt.Errorf("error: flags are not allowed")
	}

	return nil
}
