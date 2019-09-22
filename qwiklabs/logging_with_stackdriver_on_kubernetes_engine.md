---
title: "Logging with Stackdriver on Kubernetes Engine"
tags: google-cloud-platform, google-kubernetes-engine, kubernetes, google-cloud-stackdriver
url: https://www.qwiklabs.com/focuses/5539
---

# Goal
- Cloud Storage
- Pub/Sub
- BigQuery

# Task
- [x] Overview
- [x] Architecture
- [x] Setup
- [x] Deployment
- [x] Validation
- [x] Generating Logs
- [x] Logs in Stackdriver
- [x] Viewing Log Exports
- [x] Logs in Cloud Storage
- [x] Logs in BigQuery
- [x] Teardown
- [x] Troubleshooting for your production environment

# Supplement
## Setup
```sh
git clone https://github.com/GoogleCloudPlatform/gke-logging-sinks-demo
cd gke-logging-sinks-demo
gcloud config set compute/region us-central1
gcloud config set compute/zone us-central1-a
gcloud config set project qwiklabs-gcp-1b7fabe56f996496
```

## Deployment
```sh
emacs ~/gke-logging-sinks-demo/terraform/provider.tf

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

## Reference
- https://cloud.google.com/monitoring/kubernetes-engine/
- https://cloud.google.com/logging/docs/view/overview
- https://cloud.google.com/logging/docs/view/advanced-filters
- https://cloud.google.com/logging/docs/export/
- https://cloud.google.com/solutions/processing-logs-at-scale-using-dataflow
- https://www.terraform.io/docs/providers/google/index.html
- https://github.com/GoogleCloudPlatform/gke-logging-sinks-demo
