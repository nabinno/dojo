---
title: Launch Single Node Kubernetes Cluster
tags: kubernetes,minikube
url: https://learning.oreilly.com/scenarios/launch-a-single/9781492062066/
---

```sh
##
Your Interactive Bash Terminal.

$ minikube version
minikube version: v1.6.2
commit: 54f28ac5d3a815d1196cd5d57d707439ee4bb392
$ minikube start --wait=false
* minikube v1.6.2 on Ubuntu 18.04
* Selecting 'none' driver from user configuration (alternates: [])
* Running on localhost (CPUs=2, Memory=2461MB, Disk=47990MB) ...
* OS release is Ubuntu 18.04.3 LTS
* Preparing Kubernetes v1.17.0 on Docker '18.09.7' ...
  - kubelet.resolv-conf=/run/systemd/resolve/resolv.conf
* Pulling images ...
* Launching Kubernetes ...
* Configuring local host environment ...
* Done! kubectl is now configured to use "minikube"

##
$ kubectl cluster-info
Kubernetes master is running at https://172.17.0.31:8443
KubeDNS is running at https://172.17.0.31:8443/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy

To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.
$ kubectl get nodes
NAME       STATUS   ROLES    AGE    VERSION
minikube   Ready    master   119s   v1.17.0

##
$ kubectl create deployment first-deployment --image=katacoda/docker-http-server
deployment.apps/first-deployment created
$ kubectl get pods
NAME                               READY   STATUS    RESTARTS   AGE
first-deployment-666c48b44-2n4dl   1/1     Running   0          19s
$ kubectl get pods
NAME                               READY   STATUS    RESTARTS   AGE
first-deployment-666c48b44-2n4dl   1/1     Running   0          20s
$ kubectl expose deployment first-deployment --port=80 --type=NodePort
service/first-deployment exposed
$ export PORT=$(kubectl get svc first-deployment -o go-template='{{range.spec.ports}}{{if .nodePort}}{{.nodePort}}{{"\n"}}{{end}}{{end}}')
$ echo "Accessing host01:$PORT"
Accessing host01:32129
$ curl host01:$PORT
<h1>This request was processed by host: first-deployment-666c48b44-2n4dl</h1>

##
$ minikube addons enable dashboard
* dashboard was successfully enabled
$ kubectl apply -f /opt/kubernetes-dashboard.yaml
service/kubernetes-dashboard-katacoda created
$ kubectl get pods -n kubernetes-dashboard -w
NAME                                         READY   STATUS    RESTARTS   AGE
dashboard-metrics-scraper-7b64584c5c-dq8b8   1/1     Running   0          4m32s
kubernetes-dashboard-79d9cd965-gnlj7         1/1     Running   0          4m32s
```
