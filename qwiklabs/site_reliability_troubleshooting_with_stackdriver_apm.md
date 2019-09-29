---
title: "Site Reliability Troubleshooting with Stackdriver APM"
tags: google-cloud-platform, site-reliability-engineering, google-cloud-stackdriver, monitoring
url: https://google.qwiklabs.com/focuses/4186
---

# Goal
- How to deploy a microservices application on an existing GKE cluster
- How to select appropriate SLIs/SLOs for an application
- How to implement SLIs using Stackdriver Monitoring features
- How to use Stackdriver Trace, Profiler, and Debugger to identify software issues

# Task
- [ ] Environment Setup
- [ ] Infrastructure setup
- [ ] Create Stackdriver workspace
- [ ] Deploy application
- [ ] Develop Sample SLOs and SLIs
- [ ] Configure Latency SLI
- [ ] Configure Availability SLI
- [ ] Deploy new release
- [ ] Send some data
- [ ] Latency SLO Violation - Find the Problem
- [ ] Deploy Change to Address Latency
- [ ] Error Rate SLO Violation - Find the Problem
- [ ] Deploy Change to Address Error Rate
- [ ] Application optimization with Stackdriver APM

# Supplement
https://github.com/GoogleCloudPlatform/microservices-demo

![]()

```uml
skinparam monochrome true
skinparam backgroundColor #EEEEFF

actor User as U
actor "Load Generator" as L
control Frontend as F
agent CartSrv as S1
agent RecommendationSrv as S2
agent CheckoutSrv as S3
agent PaymentSrv as S31
agent EmailSrv as S32
agent ProductCatalogSrv as S4
agent CurrencySrv as S5
agent ShippingSrv as S6
agent AdSrv as S7
cloud "External API" as E
database "Cache (redis)" as C

U --> F
L --> F

F ---> S1
F ---> S2
F ---> S3
F ---> S4
F ---> S5
F ---> S6
F ---> S7

S3 --> S31
S3 --> S32
S3 -> S1
S3 -> S5
S3 -> S6

S2 -> S4

S1 --> C

S5 --> E
```

## Infrastructure setup
```sh
gcloud config set compute/zone us-west1-b
export PROJECT_ID=$(gcloud info --format='value(config.project)')
gcloud container clusters list
```

## Create Stackdriver workspace
```sh
gcloud container clusters list
gcloud container clusters get-credentials shop-cluster --zone us-west1-b

kubectl get nodes
```

## Deploy application
```sh
git clone -b APM-Troubleshooting-Demo-2 https://github.com/blipzimmerman/microservices-demo-1
curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/latest/skaffold-linux-amd64
chmod +x skaffold
sudo mv skaffold /usr/local/bin
cd microservices-demo-1
skaffold run

kubectl get pods

export EXTERNAL_IP=$(kubectl get service frontend-external | awk 'BEGIN { cnt=0; } { cnt+=1; if (cnt > 1) print $4; }')
curl -o /dev/null -s -w "%{http_code}\n"  http://$EXTERNAL_IP

./setup_csr.sh
```

## Develop Sample SLOs and SLIs
## Configure Latency SLI
## Configure Availability SLI
## Deploy new release
```sh
skaffold run
kubectl get pods
```

## Send some data
## Latency SLO Violation - Find the Problem
## Deploy Change to Address Latency
```yaml
      containers:
      - name: server
        image: gcr.io/accl-19-dev/frontend:rel013019fix
        imagePullPolicy: Always
```
```sh
skaffold run
```

## Error Rate SLO Violation - Find the Problem
## Deploy Change to Address Error Rate
```yaml
      containers:
      - name: server
        image: gcr.io/accl-19-dev/frontend:rel013019fix
        imagePullPolicy: Always
```
```sh
skaffold run
```

## Reference
- https://cloud.google.com/stackdriver/docs/
- https://cloud.google.com/trace/docs/
- https://cloud.google.com/logging/docs/
- https://cloud.google.com/error-reporting/docs/
- https://cloud.google.com/debugger/docs/
- https://cloud.google.com/blog/products/gcp/stackdriver-debugger-add-application-logs-on-the-fly-with-no-restarts
