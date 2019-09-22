---
title: "Monitoring with Stackdriver on Kubernetes Engine"
tags: google-cloud-platform, google-kubernetes-engine, kubernetes, google-cloud-stackdriver, google-cloud-stackdriver-monitor, monitoring, terraform
url: https://www.qwiklabs.com/focuses/5157
---

# Goal
- Monitoring with Stackdriver on Kubernetes Engine

# Task
- [x] Overview
- [x] Architecture
- [x] Setup
- [x] Clone demo
- [x] Create Stackdriver workspace
- [x] Validation
- [x] Teardown

# Supplement
## Setup
```sh
gcloud config set compute/region us-central1
gcloud config set compute/zone us-central1-a
gcloud config set project qwiklabs-gcp-fece287c4002a81e
```

## Clone demo
```sh
git clone https://github.com/GoogleCloudPlatform/gke-monitoring-tutorial.git
cd gke-monitoring-tutorial
gcloud auth application-default login
```

## Create Stackdriver workspace
```sh
emacs ~/gke-monitoring-tutorial/terraform/provider.tf

make create
```

## Validation
```sh
make validate
```

## Teardown
```sh
make teardown
```
