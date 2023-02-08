#!/bin/bash

set -e 

# Save the current directory
current_dir=$(pwd)

# Change the current directory to the script directory
cd "$(dirname "$0")"

# Execute the undeploy
kubectl delete --all deployment
kubectl delete configmap redis-configmap
kubectl delete configmap listener-configmap
kubectl delete configmap publisher-configmap
kubectl delete secret redis-secrets
kubectl delete secret listener-secrets
kubectl delete secret publisher-secrets

# Change back to the original directory
cd "$current_dir"
