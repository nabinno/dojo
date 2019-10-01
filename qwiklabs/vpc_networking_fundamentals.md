---
title: "VPC Networking Fundamentals"
tags: google-cloud-platform, virtual-private-cloud, network-engineering
url: https://google.qwiklabs.com/focuses/1229
---

# Goal
- Explore the default VPC network
- Create an auto mode network with firewall rules
- Create VM instances using Compute Engine
- Explore the connectivity for VM instances

# Task
- [x] Setup and Requirements
- [x] Explore the default network
- [x] Create a VPC network and VM instances
- [x] Explore the connectivity for VM instances
- [x] Test your Understanding

# Supplement
## Create a VPC network and VM instances
```sh
export PROJECT=$(gcloud config list --format "value(core.project)")

gcloud compute networks create mynetwork --project=$PROJECT --subnet-mode=auto
gcloud compute firewall-rules create mynetwork-allow-icmp --project=$PROJECT --description="Allows ICMP connections from any source to any instance on the network." --direction=INGRESS --priority=65534 --network=mynetwork --action=ALLOW --rules=icmp --source-ranges=0.0.0.0/0
gcloud compute firewall-rules create mynetwork-allow-internal --project=$PROJECT --description="Allows connections from any source in the network IP range to any instance on the network using all protocols." --direction=INGRESS --priority=65534 --network=mynetwork --action=ALLOW --rules=all --source-ranges=10.128.0.0/9
gcloud compute firewall-rules create mynetwork-allow-rdp --project=$PROJECT --description="Allows RDP connections from any source to any instance on the network using port 3389." --direction=INGRESS --priority=65534 --network=mynetwork --action=ALLOW --rules=tcp:3389 --source-ranges=0.0.0.0/0
gcloud compute firewall-rules create mynetwork-allow-ssh --project=$PROJECT --description="Allows TCP connections from any source to any instance on the network using port 22." --direction=INGRESS --priority=65534 --network=mynetwork --action=ALLOW --rules=tcp:22 --source-ranges=0.0.0.0/0

gcloud beta compute instances create mynet-us-vm \
    --project=$PROJECT \
    --zone=us-central1-c \
    --machine-type=f1-micro \
    --subnet=mynetwork \
    --network-tier=PREMIUM \
    --maintenance-policy=MIGRATE \
    --service-account=341164873387-compute@developer.gserviceaccount.com \
    --scopes=https://www.googleapis.com/auth/devstorage.read_only,https://www.googleapis.com/auth/logging.write,https://www.googleapis.com/auth/monitoring.write,https://www.googleapis.com/auth/servicecontrol,https://www.googleapis.com/auth/service.management.readonly,https://www.googleapis.com/auth/trace.append \
    --image=debian-9-stretch-v20190916 \
    --image-project=debian-cloud \
    --boot-disk-size=10GB \
    --boot-disk-type=pd-standard \
    --boot-disk-device-name=mynet-us-vm \
    --reservation-affinity=any

gcloud beta compute instances create mynet-eu-vm \
    --project=$PROJECT \
    --zone=europe-west1-c \
    --machine-type=f1-micro \
    --subnet=mynetwork \
    --network-tier=PREMIUM \
    --maintenance-policy=MIGRATE \
    --service-account=341164873387-compute@developer.gserviceaccount.com \
    --scopes=https://www.googleapis.com/auth/devstorage.read_only,https://www.googleapis.com/auth/logging.write,https://www.googleapis.com/auth/monitoring.write,https://www.googleapis.com/auth/servicecontrol,https://www.googleapis.com/auth/service.management.readonly,https://www.googleapis.com/auth/trace.append \
    --image=debian-9-stretch-v20190916 \
    --image-project=debian-cloud \
    --boot-disk-size=10GB \
    --boot-disk-type=pd-standard \
    --boot-disk-device-name=mynet-eu-vm \
    --reservation-affinity=any
```

## Explore the connectivity for VM instances
```sh
ping -c 3 <Enter mynet-eu-vm's internal IP here>
ping -c 3 mynet-eu-vm
ping -c 3 <Enter mynet-eu-vm's external IP here>
```

## Test your Understanding
```sh
ping -c 3 <Enter mynet-eu-vm's internal IP here>
ping -c 3 <Enter mynet-eu-vm's external IP here>
ping -c 3 <Enter mynet-eu-vm's internal IP here>
exit
```

## Reference
- https://cloud.google.com/vpc/docs/vpc
