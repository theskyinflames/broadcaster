#!/bin/bash

set -e 

# Save the current directory
current_dir=$(pwd)

# Change the current directory to the script directory
cd "$(dirname "$0")"

# Execute the deploy
kubectl create configmap redis-configmap --from-file=../redis/redis.conf
kubectl apply -f ../../k8s/redis-secrets.yml
kubectl apply -f ../../k8s/redis-pod.yml
kubectl apply -f ../../k8s/listener-secrets.yml
kubectl apply -f ../../k8s/listener-configmap.yml
kubectl apply -f ../../k8s/listener-pod.yml
kubectl apply -f ../../k8s/publisher-secrets.yml
kubectl apply -f ../../k8s/publisher-configmap.yml
kubectl apply -f ../../k8s/publisher-pod.yml

# Change back to the original directory
cd "$current_dir"
