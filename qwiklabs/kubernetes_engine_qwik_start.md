---
title: "Kubernetes Engine: Qwik Start"
tags: google-cloud-platform, kubernetes
url: https://www.qwiklabs.com/focuses/878
---

# Goal
- Google Kubernetes Engine (GKE)

# Task
- [x] Setting a default compute zone
- [x] Create a Kubernetes Engine cluster
- [x] Get authentication credentials for the cluster
- [x] Deploying an application to the cluster

# Supplement
## Setting a default compute zone
```sh
gcloud config set compute/zone us-central1-a
```

## Create a Kubernetes Engine cluster
```sh
gcloud container clusters create my-cluster
```

## Get authentication credentials for the cluster
```sh
gcloud container clusters get-credentials my-cluster
```

## Deploying an application to the cluster
```sh
kubectl run hello-server --image=gcr.io/google-samples/hello-app:1.0 --port 8080
kubectl expose deployment hello-server --type="LoadBalancer"
kubectl get service hello-server
```

## Clean up
```sh
gcloud container clusters delete my-cluster
```

## Reference
**Kubernetes**
- https://kubernetes.io/docs/concepts/services-networking/service/
- https://kubernetes.io/docs/concepts/workloads/controllers/deployment/

**Google Kubernetes Engine**
- https://cloud.google.com/kubernetes-engine/docs/concepts/node-pools
- https://cloud.google.com/kubernetes-engine/docs/how-to/node-auto-repair
- https://cloud.google.com/kubernetes-engine/docs/how-to/node-auto-upgrades
- https://cloud.google.com/kubernetes-engine/docs/concepts/cluster-autoscaler
- https://cloud.google.com/monitoring/kubernetes-engine/

**Other**
- https://cloud.google.com/container-registry/docs/
- https://cloud.google.com/compute/docs/load-balancing-and-autoscaling