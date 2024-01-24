package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nbebaw/kubectl-kubepilot/lib"
)

var version = "No Version Provided"

func main() {
	lib.DefineFlags()
	// --version
	if lib.ShowVersion {
		// Define allowed Flags
		allowedFlags := map[string]bool{"version": true}
		// Check the given flags by the user
		err := lib.CheckAllowedFlags(allowedFlags)
		if err != nil {
			fmt.Println(err)
			lib.ShowHelp()
			os.Exit(1)
		}
		fmt.Printf("kubectl-kubepilot v%s\n", version)
		return
	}

	if lib.CreateUser {
		// Define allowed Flags
		allowedFlags := map[string]bool{"create-user": true, "n": true, "u": true, "ip": true}
		// Check the given flags by the user
		err := lib.CheckAllowedFlags(allowedFlags)
		if err != nil {
			fmt.Println(err)
			lib.ShowHelp()
			os.Exit(1)
		}
		lib.MainCreateUser(lib.Namespace, lib.User)
		return
	}

	// --image-check
	if lib.ImageCheck {
		// Define allowed Flags
		allowedFlags := map[string]bool{"image-check": true, "l": true}
		// Check the given flags by the user
		err := lib.CheckAllowedFlags(allowedFlags)
		if err != nil {
			fmt.Println(err)
			lib.ShowHelp()
			os.Exit(1)
		}
		lib.MainImageChecks(lib.Level)
		return
	}

	// --cm
	if lib.Cm {
		// Define allowed Flags
		allowedFlags := map[string]bool{"cm": true, "c": true, "n": true, "k": true}
		// Check the given flags by the user
		err := lib.CheckAllowedFlags(allowedFlags)
		if err != nil {
			fmt.Println(err)
			lib.ShowHelp()
			os.Exit(1)
		}
		lib.MainConfigmap(lib.Cm_name, lib.Key, lib.Namespace)
		return
	}

	// --help
	if lib.Help {
		// Define allowed Flags
		allowedFlags := map[string]bool{"help": true}
		// Check the given flags by the user
		err := lib.CheckAllowedFlags(allowedFlags)
		if err != nil {
			fmt.Println(err)
			lib.ShowHelp()
			os.Exit(1)
		}
		lib.ShowHelp()
		return
	}

	// Show help when user give only the main options
	if len(os.Args) == 1 {
		lib.ShowHelp()
		return
	}

	// Show error and help when user write an illegal arguments/options
	if flag.NArg() > 0 {
		fmt.Println("Error: Illegal option or argument.")
		lib.ShowHelp()
		return
	}
}
