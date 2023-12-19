package lib

import "fmt"

var version = "vAlpha0.1" // Update this variable with your plugin's version

func PrintVersion() {
	fmt.Printf("kubectl-kubepilot %s\n", version)
}
