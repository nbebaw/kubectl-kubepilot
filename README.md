# kubectl-kubepilot
Kubepilot is a kubectl plugin designed to simplify operations related to ConfigMaps and image vulnerability scanning within Kubernetes clusters.

### Dependencies
- <b>kubectl</b>: https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/
- <b>trivy</b>: https://aquasecurity.github.io/trivy

### Usage
```sh
Kubectl plugin called kubepilot
Usage:
  kubectl kubepilot [options]

Options:
  --cm                  : Fetch value from ConfigMap
     -c                 : ConfigMap name
     -n                 : Namespace
     -k                 : Key in ConfigMap
  --create-user         : Create user and user rights for a specific namespace
     -n                 : Namespace
     -u                 : Username
  --image-check         : Scan images for vulnerabilities
   -l                   : Vulnerabilitie Level
  --version             : Show version information
  --help                : Show Help
```
### Example Usage
```sh
kubectl kubepilot --cm -c <ConfigMapName> -n <Namespace> -k <Key>
```
```sh
kubectl kubepilot --image-check -l <VulnerabilityLevel> (CRITICAL, HIGH, ...)
```
```sh
kubectl kubepilot --create-user  -n <NAMESPACE> -u <USER>
```

### Installation
```sh
wget https://github.com/nbebaw/kubectl-kubepilot/releases/download/v0.1.4/kubectl-kubepilot
chmod +x kubectl-kubepilot
sudo cp kubectl-kubepilot /usr/local/bin
```
