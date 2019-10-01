---
title: "VPC Networks - Controlling Access"
tags: google-cloud-platform, virtual-private-cloud, network-engineering
url: https://google.qwiklabs.com/focuses/1231
---

# Goal
- Create an nginx web server
- Create tagged firewall rules
- Create a service account with IAM roles
- Explore permissions for the Network Admin and Security Admin roles

# Task
- [x] Create the web servers
- [x] Create the firewall rule
- [x] Explore the Network and Security Admin roles

# Supplement
## Create the web servers
```sh
export PROJECT=$(gcloud config list --format "value(core.project)")

gcloud beta compute instances create blue \
    --project=$PROJECT
    --zone=us-central1-a \
    --machine-type=n1-standard-1 \
    --subnet=default \
    --network-tier=PREMIUM \
    --maintenance-policy=MIGRATE \
    --service-account=532003404587-compute@developer.gserviceaccount.com \
    --scopes=https://www.googleapis.com/auth/devstorage.read_only,https://www.googleapis.com/auth/logging.write,https://www.googleapis.com/auth/monitoring.write,https://www.googleapis.com/auth/servicecontrol,https://www.googleapis.com/auth/service.management.readonly,https://www.googleapis.com/auth/trace.append \
    --tags=web-server \
    --image=debian-9-stretch-v20190916 \
    --image-project=debian-cloud \
    --boot-disk-size=10GB \
    --boot-disk-type=pd-standard \
    --boot-disk-device-name=blue \
    --reservation-affinity=any

gcloud beta compute instances create green \
    --project=$PROJECT
    --zone=us-central1-a \
    --machine-type=n1-standard-1 \
    --subnet=default \
    --network-tier=PREMIUM \
    --maintenance-policy=MIGRATE \
    --service-account=532003404587-compute@developer.gserviceaccount.com \
    --scopes=https://www.googleapis.com/auth/devstorage.read_only,https://www.googleapis.com/auth/logging.write,https://www.googleapis.com/auth/monitoring.write,https://www.googleapis.com/auth/servicecontrol,https://www.googleapis.com/auth/service.management.readonly,https://www.googleapis.com/auth/trace.append \
    --image=debian-9-stretch-v20190916 \
    --image-project=debian-cloud \
    --boot-disk-size=10GB \
    --boot-disk-type=pd-standard \
    --boot-disk-device-name=green \
    --reservation-affinity=any
```

**Blue Server**
```sh
sudo apt-get install nginx-light -y
sudo nano /var/www/html/index.nginx-debian.html
nano> <h1>Welcome to the blue server!</h1>

cat /var/www/html/index.nginx-debian.html
exit
```

**Green Server**
```sh
sudo apt-get install nginx-light -y
sudo nano /var/www/html/index.nginx-debian.html
nano> <h1>Welcome to the blue server!</h1>

cat /var/www/html/index.nginx-debian.html
exit
```

## Create the firewall rule
```sh
gcloud compute firewall-rules create allow-http-web-server \
    --project=$PROJECT \
    --direction=INGRESS \
    --priority=1000 \
    --network=default \
    --action=ALLOW \
    --rules=tcp:80,icmp \
    --source-ranges=0.0.0.0/0 \
    --target-tags=werb-server

gcloud compute instances create test-vm --machine-type=f1-micro --subnet=default --zone=us-central1-a

curl <Enter blue's internal IP here>
curl -c 3 <Enter green's internal IP here>
curl <Enter blue's external IP here>
curl -c 3 <Enter green's external IP here>
```

## Explore the Network and Security Admin roles
```sh
gcloud compute firewall-rules list
gcloud compute firewall-rules delete allow-http-web-server

gcloud auth activate-service-account --key-file credentials.json
gcloud compute firewall-rules list

gcloud compute firewall-rules delete allow-http-web-server
gcloud compute firewall-rules list

gcloud compute firewall-rules delete allow-http-web-server
curl -c 3 <Enter blue's external IP here>
```

## Reference
- https://cloud.google.com/compute/docs/access/service-accounts
