---
title: "Terraform Fundamentals"
tags: google-cloud-platform, terraform
url: https://www.qwiklabs.com/focuses/1208
---

# Goal
- Getting Started with Terraform in Google Cloud
- Install Terraform from Installation Binaries
- Create a VM instance infrastructure using Terraform

# Task
- [x] Introduction
- [x] Objectives
- [x] Setup and Requirements
- [x] What is Terraform?
- [x] Verifying Terraform Installation
- [x] Build Infrastructure
- [x] Test your Understanding

# Supplement
## Verifying Terraform Installation
```sh
export PROJECT=$(gcloud config list --format "value(core.project)")

terraform
```

## Build Infrastructure
```sh
cat <<EOF >instance.tf
resource "google_compute_instance" "default" {
  project      = "${PROJECT}"
  name         = "terraform"
  machine_type = "n1-standard-1"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network = "default"
    access_config {
    }
  }
}
EOF

ls
terraform init
terraform plan
terraform apply
terraform show
terraform plan
```
