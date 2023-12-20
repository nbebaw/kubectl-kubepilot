package lib

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Main Image Checks function
/*
 the function get level as string and getting all namespaces
 and getting all images from every namespace and at the end perform image scan
 {level} -> string
*/
func MainImageChecks(level string) {
	if level == "" {
		fmt.Println("Please provide a level using -l flag")
		os.Exit(1)
	} else {
		namespaces, err := getNamespaces()
		if err != nil {
			fmt.Println("Error getting namespaces:", err)
			return
		}
		for _, ns := range namespaces {
			fmt.Printf("- Scanning in namespace %s\n", ns)
			images, err := getImagesFromNamespace(ns)
			if err != nil {
				fmt.Printf("Error getting images in namespace %s: %s\n", ns, err)
				continue
			}

			if len(images) > 1 {
				for _, img := range images {
					fmt.Printf("  Scanning %s\n", img)
					if err := performImageScan(img, level); err != nil {
						fmt.Printf("  Error scanning %s: %s\n", img, err)
						continue
					}
				}
			} else {
				fmt.Println("Empty")
			}

		}
	}
}

// Get Namespaces function
/*
 the function get all namespaces
 and returns string array and error
*/
func getNamespaces() ([]string, error) {
	output, err := exec.Command("kubectl", "get", "ns", "--no-headers", "-o", "custom-columns=:metadata.name").Output()
	if err != nil {
		return nil, err
	}

	namespaces := strings.Split(strings.TrimSpace(string(output)), "\n")
	return namespaces, nil
}

// Get Images From Namespace function
/*
 the function get namespace and find out all docker images and tags
 and returns string array and error
 {namespace} -> string
*/
func getImagesFromNamespace(namespace string) ([]string, error) {
	command := `kubectl get pods,deployments,daemonsets,statefulsets,jobs,cronjobs -n ` + namespace + ` -o jsonpath="{range .items[*]}{.spec.containers[*].image}{'\n'}"`
	output, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		return nil, err
	}

	images := strings.Split(strings.TrimSpace(string(output)), "\n")
	return images, nil
}

// Perform Image Scan function
/*
 the function get imageName and level to perform the scan of vulnerabilities
 and returns only error
 {imageName} -> string
 {level} -> string
*/
func performImageScan(imageName, level string) error {
	cmd := exec.Command("trivy", "-q", "image", "--light", "--no-progress", "--severity", level, imageName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error scanning image %s: %s", imageName, err)
	}

	vulnerabilities := string(output)
	if len(vulnerabilities) > 0 {
		fmt.Println(vulnerabilities)
	} else {
		fmt.Printf("No %s vulnerabilities found in image %s\n", level, imageName)
	}
	return nil
}
