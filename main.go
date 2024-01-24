package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nbebaw/kubectl-kubepilot/lib"
)

func main() {
	var showVersion, imageCheck, cm, help, createUser bool
	var namespace, key, cm_name, level, user, ip string

	// Allowed main options
	flag.BoolVar(&imageCheck, "image-check", false, "Scan images for vulnerabilities ")
	flag.BoolVar(&cm, "cm", false, "Configmap")
	flag.BoolVar(&showVersion, "version", false, "Show version information")
	flag.BoolVar(&help, "help", false, "Show help")
	flag.BoolVar(&createUser, "create-user", false, "Create a user in a specific namespace and create kubeconfig")

	// Allowed sub options
	flag.StringVar(&cm_name, "c", "", "ConfigMap name")
	flag.StringVar(&namespace, "n", "default", "Namespace")
	flag.StringVar(&key, "k", "", "Key in ConfigMap")
	flag.StringVar(&level, "l", "", "Deployment name")
	flag.StringVar(&user, "u", "", "user")
	flag.StringVar(&ip, "ip", "", "ip")
	flag.Parse()

	// --version
	if showVersion {
		// Define allowed Flags
		allowedFlags := map[string]bool{"version": true}
		// Check the given flags by the user
		err := lib.CheckAllowedFlags(allowedFlags)
		if err != nil {
			fmt.Println(err)
			lib.ShowHelp()
			os.Exit(1)
		}
		lib.PrintVersion()
		return
	}

	if createUser {
		// Define allowed Flags
		allowedFlags := map[string]bool{"create-user": true, "n": true, "u": true, "ip": true}
		// Check the given flags by the user
		err := lib.CheckAllowedFlags(allowedFlags)
		if err != nil {
			fmt.Println(err)
			lib.ShowHelp()
			os.Exit(1)
		}
		lib.MainCreateUser(namespace, user, ip)
		return
	}

	// --image-check
	if imageCheck {
		// Define allowed Flags
		allowedFlags := map[string]bool{"image-check": true, "l": true}
		// Check the given flags by the user
		err := lib.CheckAllowedFlags(allowedFlags)
		if err != nil {
			fmt.Println(err)
			lib.ShowHelp()
			os.Exit(1)
		}
		lib.MainImageChecks(level)
		return
	}

	// --cm
	if cm {
		// Define allowed Flags
		allowedFlags := map[string]bool{"cm": true, "c": true, "n": true, "k": true}
		// Check the given flags by the user
		err := lib.CheckAllowedFlags(allowedFlags)
		if err != nil {
			fmt.Println(err)
			lib.ShowHelp()
			os.Exit(1)
		}
		lib.MainConfigmap(cm_name, key, namespace)
		return
	}

	// --help
	if help {
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
