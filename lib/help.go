package lib

import (
	"fmt"
)

func ShowHelp() {
	fmt.Println("Usage:")
	fmt.Println("  kubectl kubepilot [options]")
	fmt.Println("\nOptions:")
	fmt.Println("  --cm          	: Fetch value from ConfigMap")
	fmt.Println("     -c          	: ConfigMap name")
	fmt.Println("     -n          	: Namespace")
	fmt.Println("     -k          	: Key in ConfigMap")
	fmt.Println("  --image-check 	: Scan images for vulnerabilities")
	fmt.Println("   -l			: Vulnerabilitie Level")
	fmt.Println("  --version     	: Show version information")
}
