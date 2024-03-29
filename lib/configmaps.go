package lib

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

// Main configmap function
/*
 the function get cm_name, key, namespace as string and execute
 kubectl with the arguments
 {cm_name} -> string
 {key} -> string
 {namespace} -> string
*/
func MainConfigmap(cm_name, key, namespace string) {
	if cm_name == "" {
		fmt.Println("Please provide a configmap using -c flag")
		os.Exit(1)
	}
	if key == "" {
		fmt.Println("Please provide a key using -k flag")
		os.Exit(1)
	}

	cmd := exec.Command("kubectl", "get", "configmap", "-n", namespace, cm_name, "-o", "json")
	output, err := cmd.Output()

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	value, err := extractValueFromConfigMap(string(output), key, cm_name)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(value)
}

// Extract Value From ConfigMap function
/*
 the function get the output from the main configmap function
 and returns the value and error if the function runs in error
 {output} -> string
 {key} -> string
 {cm_name} -> string
*/
func extractValueFromConfigMap(output string, key string, cm_name string) (string, error) {
	var data map[string]interface{}

	if err := json.Unmarshal([]byte(output), &data); err != nil {
		return "", fmt.Errorf("error unmarshalling JSON: %s", err)
	}

	configMapData, ok := data["data"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("unable to extract data from ConfigMap")
	}

	value, exists := configMapData[key]
	if !exists {
		return "", fmt.Errorf("key %s not found in ConfigMap %s", key, cm_name)
	}

	return fmt.Sprintf("%v", value), nil
}
