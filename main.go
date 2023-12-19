package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nbebaw/kubectl-kubepilot/lib"
)

func main() {
	var showVersion, imageCheck, cm bool
	var namespace, key, cm_name, level string

	flag.BoolVar(&imageCheck, "image-check", false, "Scan images for vulnerabilities ")
	flag.BoolVar(&cm, "cm", false, "Configmap")
	flag.BoolVar(&showVersion, "version", false, "Show version information")

	flag.StringVar(&cm_name, "c", "", "ConfigMap name")
	flag.StringVar(&namespace, "n", "default", "Namespace")
	flag.StringVar(&key, "k", "", "Key in ConfigMap")
	flag.StringVar(&level, "l", "", "Deployment name")
	flag.Parse()
	if showVersion {
		lib.PrintVersion()
		return
	}
	if imageCheck {
		lib.MainImageChecks(level)
		return
	}
	if cm {
		lib.MainConfigmap(cm_name, key, namespace)
		return
	}
	if len(os.Args) == 1 {
		lib.ShowHelp()
		return
	}
	if flag.NArg() > 0 {
		fmt.Println("Error: Illegal option or argument.")
		lib.ShowHelp()
		return
	}

}
