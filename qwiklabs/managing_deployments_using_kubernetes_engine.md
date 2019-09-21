---
title: "Managing Deployments Using Kubernetes Engine"
tags: google-cloud-platform, google-kubernetes-engine, kubernetes, docker
url: https://www.qwiklabs.com/focuses/639
---

# Goal
- Practice with kubectl tool
- Create deployment yaml files
- Launch, update, and scale deployments
- Practice with updating deployments and deployment styles

# Task
- [ ] Introduction to deployments
- [ ] Setup
- [ ] Learn about the deployment object
- [ ] Create a deployment
- [ ] Rolling update
- [ ] Canary deployments
- [ ] Blue-green deployments

# Supplement
```sh
gcloud config set compute/zone us-central1-a
git clone https://github.com/googlecodelabs/orchestrate-with-kubernetes.git
cd orchestrate-with-kubernetes/kubernetes
gcloud container clusters create bootcamp --num-nodes 5 --scopes "https://www.googleapis.com/auth/projecthosting,storage-rw"
```

## Learn about the deployment object
```sh
kubectl explain deployment
kubectl explain deployment --recursive
kubectl explain deployment.metadata.name
```

## Create a deployment
```sh
cat deployments/auth.yaml
kubectl create -f deployments/auth.yaml
kubectl get deployments
kubectl get replicasets
kubectl get pods
kubectl create -f services/auth.yaml
kubectl create -f deployments/hello.yaml
kubectl create -f services/hello.yaml
kubectl create secret generic tls-certs --from-file tls/
kubectl create configmap nginx-frontend-conf --from-file=nginx/frontend.conf
kubectl create -f deployments/frontend.yaml
kubectl create -f services/frontend.yaml
kubectl get services frontend
curl -ks https://<EXTERNAL-IP>
curl -ks https://`kubectl get svc frontend -o=jsonpath="{.status.loadBalancer.ingress[0].ip}"`

kubectl explain deployment.spec.replicas
kubectl scale deployment hello --replicas=5
kubectl get pods | grep hello- | wc -l
kubectl scale deployment hello --replicas=3
kubectl get pods | grep hello- | wc -l
```

## Rolling update
```sh
kubectl edit deployment hello
kubectl get replicaset
kubectl rollout history deployment/hello

kubectl rollout pause deployment/hello
kubectl rollout status deployment/hello
kubectl get pods -o jsonpath --template='{range .items[*]}{.metadata.name}{"\t"}{"\t"}{.spec.containers[0].image}{"\n"}{end}'

kubectl rollout resume deployment/hello
kubectl rollout status deployment/hello

kubectl rollout undo deployment/hello
kubectl rollout history deployment/hello
kubectl get pods -o jsonpath --template='{range .items[*]}{.metadata.name}{"\t"}{"\t"}{.spec.containers[0].image}{"\n"}{end}'
```

## Canary deployments
```sh
cat deployments/hello-canary.yaml
kubectl create -f deployments/hello-canary.yaml
kubectl get deployments
curl -ks https://`kubectl get svc frontend -o=jsonpath="{.status.loadBalancer.ingress[0].ip}"`/version
```

## Blue-green deployments
```sh
kubectl apply -f services/hello-blue.yaml
kubectl create -f deployments/hello-green.yaml
curl -ks https://`kubectl get svc frontend -o=jsonpath="{.status.loadBalancer.ingress[0].ip}"`/version
kubectl apply -f services/hello-green.yaml
curl -ks https://`kubectl get svc frontend -o=jsonpath="{.status.loadBalancer.ingress[0].ip}"`/version
kubectl apply -f services/hello-blue.yaml
curl -ks https://`kubectl get svc frontend -o=jsonpath="{.status.loadBalancer.ingress[0].ip}"`/version
```