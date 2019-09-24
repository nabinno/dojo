---
title: "Awwvision: Cloud Vision API from a Kubernetes Cluster"
tags: google-cloud-platform, google-cloud-vision, kubernetes
url: https://www.qwiklabs.com/focuses/1241
---

# Goal
- A simple Redis instance
- A web app that displays the labels and associated images
- A worker that handles scraping Reddit for images and classifying them using the Vision API. Cloud Pub/Sub is used to coordinate tasks between multiple worker instances

# Task
- [ ] Create a Kubernetes Engine cluster
- [ ] Get the Sample
- [ ] Deploy the sample
- [ ] Check the Kubernetes resources on the cluster
- [ ] Visit your new web app and start its crawler

# Supplement
## Create a Kubernetes Engine cluster
```sh
gcloud config set compute/zone us-central1-a
gcloud config set project xxxxxxxxxxxxxxxxxxxx
gcloud container clusters create awwvision --num-nodes 2 --scopes cloud-platform
gcloud container clusters get-credentials awwvision
kubectl cluster-info
```

## Get the Sample
```sh
cd cloud-vision/python/awwvision
make all
```

## Deploy the sample
```sh
kubectl get pods
kubectl get deployments -o wide
kubectl get svc awwvision-webapp
```

## Reference
- https://github.com/GoogleCloudPlatform/cloud-vision/tree/master/python/awwvision
