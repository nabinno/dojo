---
title: "Hello Node Kubernetes"
tags: google-cloud-platform, google-kubernetes-engine, docker
url: https://google.qwiklabs.com/focuses/564
---

# Goal
- Create a Node.js server
- Create a Docker container image
- Create a container cluster
- Create a Kubernetes pod
- Scale up your services

# Task
- [x] Create your Node.js application
- [x] Create a Docker container image
- [x] Create your cluster
- [x] Create your pod
- [x] Allow external traffic
- [x] Scale up your service
- [x] Roll out an upgrade to your service
- [x] Kubernetes graphical dashboard (optional)

# Supplement
## Create your Node.js application
```js
var http = require('http');
var handleRequest = function(request, response) {
  response.writeHead(200);
  response.end("Hello World!");
}
var www = http.createServer(handleRequest);
www.listen(8080);
```

## Create a Docker container image
```dockerfile
FROM node:6.9.2
EXPOSE 8080
COPY server.js .
CMD node server.js
```
```sh
docker build -t gcr.io/qwiklabs-gcp-e86ded7157c97b3d/hello-node:v1 .
docker run -d -p 8080:8080 gcr.io/qwiklabs-gcp-e86ded7157c97b3d/hello-node:v1
curl http://localhost:8080
docker ps
docker stop $(docker ps -q)
gcloud auth configure-docker
docker push gcr.io/qwiklabs-gcp-e86ded7157c97b3d/hello-node:v1
```

## Create your cluster
```sh
gcloud config set project qwiklabs-gcp-e86ded7157c97b3d
gcloud container clusters create hello-world --num-nodes 2 --machine-type n1-standard-1 --zone us-central1-a
```

## Create your pod
```sh
kubectl run hello-node --image=gcr.io/qwiklabs-gcp-e86ded7157c97b3d/hello-node:v1 --port=8080
kubectl get deployments
kubectl get pods
kubectl cluster-info
kubectl config view
kubectl get events
kubectl logs hello-node-6ddf964b4b-6fkwz
```

## Allow external traffic
```sh
kubectl expose deployment hello-node --type="LoadBalancer"
kubectl get services
curl http://34.70.211.109:8080
```

## Scale up your service
```sh
kubectl scale deployment hello-node --replicas=4
kubectl get deployment
kubectl get pods
```

## Roll out an upgrade to your service
```sh
docker build -t gcr.io/qwiklabs-gcp-e86ded7157c97b3d/hello-node:v2 .
docker push gcr.io/qwiklabs-gcp-e86ded7157c97b3d/hello-node:v2
kubectl edit deployment hello-node

kubectl get deployments
```

## Kubernetes graphical dashboard (optional)
```sh
kubectl create clusterrolebinding cluster-admin-binding --clusterrole=cluster-admin --user=$(gcloud config get-value account)
kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v1.10.1/src/deploy/recommended/kubernetes-dashboard.yaml
kubectl -n kube-system edit service kubernetes-dashboard
kubectl -n kube-system describe $(kubectl -n kube-system get secret -n kube-system -o name | grep namespace) | grep token:
```

## Reference
- https://kubernetes.io/docs/tutorials/kubernetes-basics/update/update-intro/
