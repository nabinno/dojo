---
title: "Getting Started with Cloud Shell & gcloud"
tags: google-cloud-platform
url: https://www.qwiklabs.com/focuses/563
---

# Goal
- Practice using gcloud commands
- Connect to compute services hosted on the Google Cloud Platform

# Task
- [x] Start Cloud Shell
- [x] Understanding Regions and Zones
- [x] Initializing Cloud SDK
- [x] Setting environment variables
- [x] Create a virtual machine with gcloud
- [x] Using gcloud commands
- [x] Auto-completion
- [x] SSH into your vm instance
- [x] Use the Home directory
- [x] Test your knowledge

# Supplement
## Understanding Regions and Zones
```sh
gcloud compute project-info describe --project $(gcloud config list --format "value(core.project)")
```

## Initializing Cloud SDK
## Setting environment variables
```sh
export PROJECT_ID=$(gcloud config list --format "value(core.project)")
export ZONE=$(gcloud compute project-info describe --format "value(commonInstanceMetadata.items['google-compute-default-zone'])" --project $(gcloud config list --format "value(core.project)"))
echo $PROJECT_ID
echo $ZONE
```

## Create a virtual machine with gcloud
```sh
gcloud compute instances create gcelab2 --machine-type n1-standard-2 --zone $ZONE
gcloud compute instances create --help
```

## Using gcloud commands
```sh
gcloud -h
gcloud config --help
gcloud help config
gcloud config list
gcloud config list --all
gcloud components list
```

## Auto-completion
```sh
gcloud components install beta
gcloud beta interactive
interactive> gcloud compute instances describe <your_vm>
```

## SSH into your vm instance
```sh
gcloud compute ssh gcelab2 --zone $ZONE
gcelab2> exit
```

## Use the Home directory
```sh
cd $HOME
vi ./.bashrc
```
