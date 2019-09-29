---
title: "Multiple VPC Networks"
tags: google-cloud-platform, virtual-private-cloud, network-engineering
url: https://google.qwiklabs.com/focuses/1230
---

# Goal
- Create custom mode VPC networks with firewall rules
- Create VM instances using Compute Engine
- Explore the connectivity for VM instances across VPC networks
- Create a VM instance with multiple network interfaces

# Task
- [x] Create custom mode VPC networks with firewall rules
- [x] Create VM instances
- [x] Explore the connectivity between VM instances
- [x] Create a VM instance with multiple network interfaces

# Supplement
## Create custom mode VPC networks with firewall rules
```sh
gcloud compute networks create managementnet --subnet-mode=custom
gcloud compute networks subnets create managementsubnet-us --network=managementnet --region=us-central1 --range=10.130.0.0/20

gcloud compute networks create privatenet --subnet-mode=custom
gcloud compute networks subnets create privatesubnet-us --network=privatenet --region=us-central1 --range=172.16.0.0/24
gcloud compute networks subnets create privatesubnet-eu --network=privatenet --region=europe-west1 --range=172.20.0.0/20

gcloud compute networks list

gcloud compute networks subnets list --sort-by=NETWORK

gcloud compute firewall-rules create managementnet-allow-icmp-ssh-rdp \
    --direction=INGRESS \
    --priority=1000 \
    --network=managementnet \
    --action=ALLOW \
    --rules=icmp,tcp:22,tcp:3389 \
    --source-ranges=0.0.0.0/0

gcloud compute firewall-rules create privatenet-allow-icmp-ssh-rdp \
    --direction=INGRESS \
    --priority=1000 \
    --network=privatenet \
    --action=ALLOW \
    --rules=icmp,tcp:22,tcp:3389 \
    --source-ranges=0.0.0.0/0

gcloud compute firewall-rules list --sort-by=NETWORK
```

## Create VM instances
```sh
gcloud compute instances create managementnet-us-vm \
    --zone=us-central1-c \
    --machine-type=n1-standard-1 \
    --subnet=managementsubnet-us

gcloud compute instances create privatenet-us-vm \
    --zone=us-central1-c \
    --machine-type=n1-standard-1 \
    --subnet=privatesubnet-us

gcloud compute instances list --sort-by=ZONE
```

## Explore the connectivity between VM instances
```sh
ping -c 3 <mynet-eu-vm external IP>
ping -c 3 <managementnet-us-vm external IP>
ping -c 3 <privatenet-us-vm external IP>

ping -c 3 <mynet-eu-vm internal IP>
ping -c 3 <managementnet-us-vm internal IP>
ping -c 3 <privatenet-us-vm internal IP>
```

## Create a VM instance with multiple network interfaces
```sh
sudo ifconfig

ping -c 3 <privatenet-us-vm internal IP>
ping -c 3 privatenet-us-vm
ping -c 3 <managementnet-us-vm internal IP>
ping -c 3 <mynet-us-vm internal IP>
ping -c 3 <mynet-eu-vm internal IP>

ip route
```
