package lib

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func MainCreateUser(ns, user string) {
	if ns == "" {
		fmt.Println("Please provide a namespace using -ns flag")
		os.Exit(1)
	}
	if user == "" {
		fmt.Println("Please provide a username using -u flag")
		os.Exit(1)
	}

	// execute kubectl apply -f with text
	execWithString(generateServiceAccount(user, ns))
	execWithString(generateRole(user, ns))
	execWithString(generateRoleBinding(user, ns))

	cert := configView("cluster.certificate-authority-data")
	serverName := configView("name")
	token := generateToken(user, ns)
	fmt.Println("======================= Kubeconfig =======================")
	fmt.Println(generateKubeConfig(cert, serverName, user, ns, token))
	fmt.Println("======================= End Kubeconfig =======================")
}

func execWithString(deployment string) {
	// Set the kubectl command and arguments
	cmd := exec.Command("kubectl", "apply", "-f", "-")

	// Set the input for the kubectl command
	cmd.Stdin = strings.NewReader(deployment)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the kubectl command with the formatted YAML content as input
	if err := cmd.Run(); err != nil {
		fmt.Println("Error running kubectl command:", err)
		os.Exit(1)
	}
}

func configView(data string) string {
	cmd := exec.Command("kubectl", "config", "view", "--minify", "--flatten", "-o", "jsonpath='{.clusters[]."+data+"}'")
	outputCert, err := cmd.Output()
	if err != nil {
		fmt.Println("Error config view:", err)
		os.Exit(1)
	}
	return string(outputCert)
}

func generateToken(user, namespace string) string {
	cmd := exec.Command("kubectl", "create", "token", user, "--duration", "315360000s", "-n", namespace)
	token, err := cmd.Output()
	if err != nil {
		fmt.Println("Error generate Token:", err)
		os.Exit(1)
	}
	return string(token)
}

func generateServiceAccount(user, namespace string) string {
	serviceAccountTemplate := `
apiVersion: v1
kind: ServiceAccount
metadata:
 name: %s
 namespace: %s
`
	serviceAccount := fmt.Sprintf(serviceAccountTemplate, user, namespace)
	return serviceAccount
}

func generateRole(user, namespace string) string {
	roleTemplate := `
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: %s-role
  namespace: %s
rules:
- apiGroups: ["*"]
  resources: ["*"]
  verbs: ["*"]	
`
	role := fmt.Sprintf(roleTemplate, user, namespace)
	return role
}

func generateRoleBinding(user, namespace string) string {
	roleBindingTemplate := `
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: %s-role-binding
  namespace: %s
subjects:
- kind: ServiceAccount
  name: %s
  namespace: %s
roleRef:
  kind: Role
  name: %s-role
  apiGroup: rbac.authorization.k8s.io
`
	roleBinding := fmt.Sprintf(roleBindingTemplate, user, namespace, user, namespace, user)
	return roleBinding
}
func getHostIp() string {
	cmd := exec.Command("hostname", "-i")
	ip, err := cmd.Output()
	if err != nil {
		fmt.Println("Error getting Host IP:", err)
		os.Exit(1)
	}
	return strings.TrimSpace(string(ip))
}
func generateKubeConfig(cert, serverName, user, namespace, token string) string {
	ip := getHostIp()
	kubeConfigTemplate := `
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: %s
    server: https://%s:6443
  name: %s
contexts:
- context:
    cluster: %s
    user: %s
    namespace: %s
  name: %s
current-context: %s
kind: Config
preferences: {}
users:
- name: %s
  user:
    token: %s
`
	kubeConfig := fmt.Sprintf(kubeConfigTemplate, cert, ip, serverName, serverName, user, namespace, namespace, namespace, user, token)
	return kubeConfig
}
