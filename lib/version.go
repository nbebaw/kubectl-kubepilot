package lib

import "fmt"

var version = "v0.1.0" // Update this variable with your plugin's version

// Print plugin Version
/*
 Print plugin Version
*/
func PrintVersion() {
	fmt.Printf("kubectl-kubepilot %s\n", version)
}
