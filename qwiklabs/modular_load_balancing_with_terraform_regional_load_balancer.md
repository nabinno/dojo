---
title: "Modular Load Balancing with Terraform - Regional Load Balancer"
tags: google-cloud-platform, load-balancing, terraform
url: https://www.qwiklabs.com/focuses/1207
---

# Goal
- Learn about the load balancing modules for Terraform
- Create a regional TCP load balancer
- Create a regional internal TCP load balancer
- Create a global HTTP load balancer with Kubernetes Engine
- Create a global HTTPS content-based load balancer

# Task
- [x] Overview
- [x] Objectives
- [x] Setup
- [x] Terraform Modules Overview
- [x] Install Terraform
- [x] Clone the examples repository
- [x] TCP load balancer with regional forwarding rule

# Supplement
## Terraform Modules Overview
```tf
module "gce-lb-fr" {
  source       = "github.com/GoogleCloudPlatform/terraform-google-lb"
  region       = "${var.region}"
  name         = "group1-lb"
  service_port = "${module.mig1.service_port}"
  target_tags  = ["${module.mig1.target_tags}"]
}
```

```tf
module "gce-ilb" {
  source         = "github.com/GoogleCloudPlatform/terraform-google-lb-internal"
  region         = "${var.region}"
  name           = "group2-ilb"
  ports          = ["${module.mig2.service_port}"]
  health_port    = "${module.mig2.service_port}"
  source_tags    = ["${module.mig1.target_tags}"]
  target_tags    = ["${module.mig2.target_tags}","${module.mig3.target_tags}"]
  backends       = [
    { group = "${module.mig2.instance_group}" },
    { group = "${module.mig3.instance_group}" },
  ]
}
```

```tf
module "gce-lb-http" {
  source            = "github.com/GoogleCloudPlatform/terraform-google-lb-http"
  name              = "group-http-lb"
  target_tags       = ["${module.mig1.target_tags}", "${module.mig2.target_tags}"]
  backends          = {
    "0" = [
      { group = "${module.mig1.instance_group}" },
      { group = "${module.mig2.instance_group}" }
    ],
  }
  backend_params    = [
    # health check path, port name, port number, timeout seconds.
    "/,http,80,10"
  ]
}
```

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
git clone https://github.com/GoogleCloudPlatform/terraform-google-lb
cd ~/terraform-google-lb/
git checkout -b 028dd3dc09ae1bf47e107ab435310c0a57b1674c
cd ~/terraform-google-lb/examples/basic

nano main.tf
nano> provider google {
  region = "${var.region}"
  version = "1.18.0"
}
```

## TCP load balancer with regional forwarding rule
```sh
export GOOGLE_PROJECT=$(gcloud config list --format "value(core.project)")
terraform init
terraform plan
terraform apply
EXTERNAL_IP=$(terraform output load-balancer-ip)
echo "http://${EXTERNAL_IP}"
```
