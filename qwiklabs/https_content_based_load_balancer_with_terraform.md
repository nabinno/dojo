---
title: "HTTPS Content-Based Load Balancer with Terraform"
tags: google-cloud-platform, load-balancing, terraform
url: https://www.qwiklabs.com/focuses/1206
---

# Goal
- Learn about the load balancing modules for Terraform
- Configure Terraform in GCP environment
- Create a global HTTPS Content-Based Load Balancer

# Task
- [x] Overview
- [x] Objectives
- [x] Setup
- [x] Install Terraform
- [x] Clone the examples repository
- [x] Content Based Load Balancer

# Supplement
## Install Terraform
```sh
terraform version
wget https://github.com/warrensbox/terraform-switcher/releases/download/0.7.737/terraform-switcher_0.7.737_linux_amd64.tar.gz
mkdir -p ${HOME}/bin
tar -xvf terraform-switcher_0.7.737_linux_amd64.tar.gz -C ${HOME}/bin
export PATH=$PATH:${HOME}/bin
tfswitch -b ${HOME}/bin/terraform 0.11.14
echo "0.11.14" >> .tfswitchrc
exec -l $SHELL
terraform version
```

## Clone the examples repository
```sh
git clone https://github.com/GoogleCloudPlatform/terraform-google-lb-http.git
git check -b https-content d3ddcdb1983834bd09e5f8f7af12900ea5535314
cd ~/terraform-google-lb-http/examples/https-content
# cd ~/terraform-google-lb-http/examples/multi-mig-http-lb
```

## Content Based Load Balancer
```sh
export GOOGLE_PROJECT=$(gcloud config list --format "value(core.project)")

nano main.tf
nano> provider "google" {
  region = "${var.region}"
  version = "1.18.0"
}

terraform init
terraform plan -out=tfplan
ls
terraform apply tfplan
./test.sh

EXTERNAL_IP=$(terraform output load-balancer-ip)
echo https://${EXTERNAL_IP}
curl https://${EXTERNAL_IP}/group1
curl https://${EXTERNAL_IP}/group2
curl https://${EXTERNAL_IP}/group3
```