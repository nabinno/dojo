---
title: "Create an Internal Load Balancer"
tags: google-cloud-platform, load-balancing
url: https://google.qwiklabs.com/focuses/1250
---

# Goal
- Create HTTP and health check firewall rules
- Configure two instance templates
- Create two managed instance groups
- Configure and test an internal load balancer

# Task
- [x] Configure HTTP and health check firewall rules
- [x] Configure instance templates and create instance groups
- [x] Configure the Internal Load Balancer
- [x] Test the Internal Load Balancer

# Supplement
![](create_an_internal_load_balancer.png)

```uml
skinparam monochrome true
skinparam backgroundColor #EEEEFF

node "Region(us-central)" {
    node "Subnet(A)" {
        actor Client as C
        collections InstanceGroup as IGA
    }

    control ForwardingRule as Fwd
    control "BackendService" as BS

    node "Subnet(B)" {
        collections InstanceGroup as IGB
    }
}

C ---> Fwd

Fwd -> BS

BS ---> IGA

BS ---> IGB
```

## Configure HTTP and health check firewall rules
```sh
export PROJECT=$(gcloud config list --format "value(core.project)")
export GCE_ACCOUNT=$(gcloud iam service-accounts list --filter "Compute Engine default service account"  --format "value(EMAIL)")

gcloud compute firewall-rules create app-allow-http \
    --project=$PROJECT \
    --direction=INGRESS \
    --priority=1000 \
    --network=my-internal-app \
    --action=ALLOW \
    --rules=tcp:80 \
    --source-ranges=0.0.0.0/0 \
    --target-tags=lb-backend

gcloud compute firewall-rules create app-allow-health-check \
    --project=$PROJECT \
    --direction=INGRESS \
    --priority=1000 \
    --network=default \
    --action=ALLOW \
    --rules=tcp \
    --source-ranges=130.211.0.0/22,35.191.0.0/16 \
    --target-tags=lb-backend
```

## Configure instance templates and create instance groups
```sh
gcloud beta compute instance-templates create instance-template-1 \
    --project=$PROJECT \
    --machine-type=n1-standard-1 \
    --subnet=projects/$PROJECT/regions/us-central1/subnetworks/subnet-a \
    --network-tier=PREMIUM \
    --metadata=startup-script-url=gs://cloud-training/gcpnet/ilb/startup.sh \
    --maintenance-policy=MIGRATE \
    --service-account=$GCE_ACCOUNT \
    --scopes=https://www.googleapis.com/auth/devstorage.read_only,https://www.googleapis.com/auth/logging.write,https://www.googleapis.com/auth/monitoring.write,https://www.googleapis.com/auth/servicecontrol,https://www.googleapis.com/auth/service.management.readonly,https://www.googleapis.com/auth/trace.append \
    --region=us-central1 \
    --tags=lb-backend \
    --image=debian-9-stretch-v20190916 \
    --image-project=debian-cloud \
    --boot-disk-size=10GB \
    --boot-disk-type=pd-standard \
    --boot-disk-device-name=instance-template-1 \
    --reservation-affinity=any

gcloud beta compute instance-templates create instance-template-1 \
    --project=$PROJECT \
    --machine-type=n1-standard-1 \
    --subnet=projects/$PROJECT/regions/us-central1/subnetworks/subnet-b \
    --network-tier=PREMIUM \
    --metadata=startup-script-url=gs://cloud-training/gcpnet/ilb/startup.sh \
    --maintenance-policy=MIGRATE \
    --service-account=$GCE_ACCOUNT \
    --scopes=https://www.googleapis.com/auth/devstorage.read_only,https://www.googleapis.com/auth/logging.write,https://www.googleapis.com/auth/monitoring.write,https://www.googleapis.com/auth/servicecontrol,https://www.googleapis.com/auth/service.management.readonly,https://www.googleapis.com/auth/trace.append \
    --region=us-central1 \
    --tags=lb-backend \
    --image=debian-9-stretch-v20190916 \
    --image-project=debian-cloud \
    --boot-disk-size=10GB \
    --boot-disk-type=pd-standard \
    --boot-disk-device-name=instance-template-1 \
    --reservation-affinity=any

gcloud compute instance-groups managed create instance-group-1 \
    --project=$PROJECT \
    --base-instance-name=instance-group-1 \
    --template=instance-template-1 \
    --size=1 \
    --zone=us-central1-a
gcloud beta compute instance-groups managed set-autoscaling "instance-group-1" \
    --project $PROJECT \
    --zone "us-central1-a" \
    --cool-down-period "45" \
    --max-num-replicas "5" \
    --min-num-replicas "1" \
    --target-cpu-utilization "0.8"

gcloud compute instance-groups managed create instance-group-2 \
    --project=$PROJECT \
    --base-instance-name=instance-group-2 \
    --template=instance-template-2 \
    --size=1 \
    --zone=us-central1-b
gcloud beta compute instance-groups managed set-autoscaling "instance-group-2" \
    --project $PROJECT \
    --zone "us-central1-b" \
    --cool-down-period "45" \
    --max-num-replicas "5" \
    --min-num-replicas "1" \
    --target-cpu-utilization "0.8"

gcloud beta compute instances create utility-vm \
    --project=$PROJECT
    --zone=us-central1-f \
    --machine-type=f1-micro \
    --subnet=subnet-a \
    --private-network-ip=10.10.20.50 \
    --network-tier=PREMIUM \
    --maintenance-policy=MIGRATE \
    --service-account=$GCE_ACCOUNT \
    --scopes=https://www.googleapis.com/auth/devstorage.read_only,https://www.googleapis.com/auth/logging.write,https://www.googleapis.com/auth/monitoring.write,https://www.googleapis.com/auth/servicecontrol,https://www.googleapis.com/auth/service.management.readonly,https://www.googleapis.com/auth/trace.append \
    --image=debian-9-stretch-v20190916 \
    --image-project=debian-cloud \
    --boot-disk-size=10GB \
    --boot-disk-type=pd-standard \
    --boot-disk-device-name=utility-vm \
    --reservation-affinity=any
```

```sh
curl 10.10.20.2
curl 10.10.30.2
exit
```

## Test the Internal Load Balancer
```sh
curl 10.10.30.5
```

## Reference
- https://cloud.google.com/compute/docs/instance-groups
