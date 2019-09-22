---
title: "GKE Migrating to Containers"
tags: google-cloud-platform, google-kubernetes-engine, kubernetes
url: https://www.qwiklabs.com/focuses/5155
---

# Goal
- Isolated - Applications have their own libraries; no conflicts will arise from different libraries in other applications
- Limited (limits on CPU/memory) - Applications may not hog resources from other applications
- Portable - The container contains everything it needs and is not tied to an OS or Cloud provider
- Lightweight - The kernel is shared, making it much smaller and faster than a full OS image

# Task
- [x] Overview
- [x] What you'll learn
- [x] Architecture
- [x] Setup
- [x] Deployment
- [x] Exploring Prime-Flask Environments
- [x] Validation
- [x] Load Testing
- [x] Tear Down

# Supplement
## Setup
```sh
sudo apt-get install apache2-utils
ab -V
git clone https://github.com/GoogleCloudPlatform/gke-migration-to-containers.git
cd gke-migration-to-containers
gcloud config set compute/region us-central1
gcloud config set compute/zone us-central1-a
```

## Deployment
```sh
make create
```

## Exploring Prime-Flask Environments
```sh
gcloud compute ssh vm-webserver --zone us-central1-a
vm-webserver> ps auxf
vm-webserver> exit

gcloud compute ssh cos-vm --zone us-central1-a
cos-vm> ps auxf
cos-vm> ls /usr/local/bin/python
cos-vm> sudo docker ps
cos-vm> sudo docker exec -it $(sudo docker ps | awk '/prime-flask/ {print $1}') ps auxf
cos-vm> exit

gcloud container clusters get-credentials prime-server-cluster
kubectl get pods
kubectl exec $(kubectl get pods -lapp=prime-server -ojsonpath='{.items[].metadata.name}')  -- ps aux
```

## Validation
```sh
make validate
```

## Load Testing
```sh
ab -c 120 -t 60  http://<IP_ADDRESS>/prime/10000
kubectl scale --replicas 3 deployment/prime-server
```

## Tear Down
```sh
make teardown
```