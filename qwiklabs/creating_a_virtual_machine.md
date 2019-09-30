---
title: "Creating a Virtual Machine"
tags: google-cloud-platform
url: https://www.qwiklabs.com/focuses/3563
---

# Goal
- Create a virtual machine with the GCP Console
- Create a virtual machine with gcloud command line
- Deploy a web server and connect it to a virtual machine

# Task
- [x] Create a new instance from the Cloud Console
- [x] Create a new instance with gcloud
- [x] Test your knowledge

# Supplement
## Create a new instance from the Cloud Console
- https://cloud.google.com/compute/docs/instances/connecting-to-instance
- https://cloud.google.com/compute/docs/machine-types
- https://cloud.google.com/compute/quotas
- https://cloud.google.com/compute/docs/regions-zones/

```sh
sudo su -
gcloud config set project $(gcloud config list --format "value(core.project)"
apt-get update
apt-get install nginx -y
ps auwx | grep nginx
```

## Create a new instance with gcloud
- https://cloud.google.com/compute/docs/instances/creating-instance-with-custom-machine-type
- https://cloud.google.com/compute/docs/machine-types


```sh
gcloud compute instances create gcelab2 --machine-type n1-standard-2 --zone us-central1-c

gcloud compute ssh gcelab2 --zone us-central1-c
gcelab2> exit
```

## Reference
- https://cloud.google.com/datalab/docs/how-to/machine-type
- https://cloud.google.com/vpc/docs/vpc
- https://cloud.google.com/solutions/migration-center/
- https://cloud.google.com/compute/docs/instances/
