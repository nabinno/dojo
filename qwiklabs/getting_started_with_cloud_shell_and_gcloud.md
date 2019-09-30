---
title: "Getting Started with Cloud Shell & gcloud"
tags: google-cloud-platform
url: https://www.qwiklabs.com/focuses/563
---

# Goal
- Practice using gcloud commands
- Connect to compute services hosted on the Google Cloud Platform

# Task
- [ ] Start Cloud Shell
- [ ] Understanding Regions and Zones
- [ ] Initializing Cloud SDK
- [ ] Setting environment variables
- [ ] Create a virtual machine with gcloud
- [ ] Using gcloud commands
- [ ] Auto-completion
- [ ] SSH into your vm instance
- [ ] Use the Home directory
- [ ] Test your knowledge

# Supplement
## Start Cloud Shell
## Understanding Regions and Zones
```sh
gcloud compute project-info describe --project $(gcloud config list --format "value(core.project)")
```

## Initializing Cloud SDK
## Setting environment variables
```sh
export PROJECT_ID=$(gcloud config list --format "value(core.project)")
export ZONE=<your_zone>
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
gcloud compute instances describe <your_vm>
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
