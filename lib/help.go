package lib

import (
	"fmt"
)

// Show Help
/*
 the function prints help to the user
*/
func ShowHelp() {
	fmt.Println("Usage:")
	fmt.Println("  kubectl kubepilot [options]")
	fmt.Println("\nOptions:")
	fmt.Println("  --cm          	: Fetch value from ConfigMap")
	fmt.Println("     -c          	: ConfigMap name")
	fmt.Println("     -n          	: Namespace")
	fmt.Println("     -k          	: Key in ConfigMap")
	fmt.Println("  --create-user		: Create user and user rights for a specific namespace")
	fmt.Println("     -n          	: Namespace")
	fmt.Println("     -u          	: Username")
	fmt.Println("  --image-check 	: Scan images for vulnerabilities")
	fmt.Println("   -l			: Vulnerabilitie Level")
	fmt.Println("  --version     	: Show version information")
	fmt.Println("  --help     		: Show Help")
}
