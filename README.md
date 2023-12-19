# kubectl-kubepilot
Kubepilot is a kubectl plugin designed to simplify operations related to ConfigMaps and image vulnerability scanning within Kubernetes clusters.

### Dependencies
- <b>kubectl</b>: https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/
- <b>trivy</b>: https://aquasecurity.github.io/trivy/v0.18.3/installation/

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
  --image-check         : Scan images for vulnerabilities
   -l                   : Vulnerabilitie Level
  --version             : Show version information
```
### Example Usage
```sh
kubectl kubepilot --cm -c <ConfigMapName> -n <Namespace> -k <Key>
```
```sh
kubectl kubepilot --image-check -l <VulnerabilityLevel> (CRITICAL, HIGH, ...)
```

### Installation
```sh
wget https://github.com/nbebaw/kubectl-kubepilot/releases/download/vAlpha0.1/kubectl-kubepilot
chmod +x kubectl-kubepilot
sudo cp kubectl-kubepilot /usr/bin
```
