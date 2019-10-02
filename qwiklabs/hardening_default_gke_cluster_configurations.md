---
title: "Hardening Default GKE Cluster Configurations"
tags: google-cloud-platform, google-kubernetes-engine
url: https://www.qwiklabs.com/focuses/5158
---

# Goal
- Create a small GKE cluster using the default settings
- Validate the most common paths of pod escape and cluster privilege escalation from the perspective of a malicious internal user
- Harden the GKE cluster for these issues
- Validate the cluster no longer allows for each of those actions to occur

# Task
- [ ] Create a simple GKE cluster
- [ ] Run a Google Cloud-SDK pod
- [ ] Explore the Legacy Compute Metadata Endpoint
- [ ] Explore the GKE node bootstrapping credentials
- [ ] Leverage the Permissions Assigned to this Node Pool's Service Account
- [ ] Deploy a pod that mounts the host filesystem
- [ ] Explore and compromise the underlying host
- [ ] Deploy a second node pool
- [ ] Run a Google Cloud-SDK pod
- [ ] Explore various blocked endpoints
- [ ] Deploy PodSecurityPolicy objects
- [ ] Enable PodSecurity policy
- [ ] Deploy a blocked pod that mounts the host filesystem

# Supplement
## Create a simple GKE cluster
## Run a Google Cloud-SDK pod
## Explore the Legacy Compute Metadata Endpoint
## Explore the GKE node bootstrapping credentials
## Leverage the Permissions Assigned to this Node Pool's Service Account
## Deploy a pod that mounts the host filesystem
## Explore and compromise the underlying host
## Deploy a second node pool
## Run a Google Cloud-SDK pod
## Explore various blocked endpoints
## Deploy PodSecurityPolicy objects
## Enable PodSecurity policy
## Deploy a blocked pod that mounts the host filesystem
