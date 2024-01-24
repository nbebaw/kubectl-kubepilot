package lib

import "flag"

func DefineFlags() {
	// Allowed main options
	flag.BoolVar(&ImageCheck, "image-check", false, "Scan images for vulnerabilities ")
	flag.BoolVar(&Cm, "cm", false, "Configmap")
	flag.BoolVar(&ShowVersion, "version", false, "Show version information")
	flag.BoolVar(&Help, "help", false, "Show help")
	flag.BoolVar(&CreateUser, "create-user", false, "Create a user in a specific namespace and create kubeconfig")

	// Allowed sub options
	flag.StringVar(&Cm_name, "c", "", "ConfigMap name")
	flag.StringVar(&Namespace, "n", "default", "Namespace")
	flag.StringVar(&Key, "k", "", "Key in ConfigMap")
	flag.StringVar(&Level, "l", "", "Deployment name")
	flag.StringVar(&User, "u", "", "user")
	flag.Usage = ShowHelp
	flag.Parse()
}
