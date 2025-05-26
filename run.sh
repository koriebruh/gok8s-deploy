#!/bin/bash

# Build the Docker image and Push it to the registry
docker build -t koriebruh/go-apps-deployment:v2 .
docker push koriebruh/go-apps-deployment:v2

# load the application into Kubernetes
kubectl apply -f k8s/deployment.yaml