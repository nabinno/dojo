---
title: Video Intelligence: Qwik Start
tags: google-cloud-video-intelligence
url: https://www.qwiklabs.com/focuses/582?parent=catalog
---

# Goal
- Google Cloud Video Intelligence

# Task
- [x] Overview
- [x] Setup and Requirements
- [x] Enable the Video Intelligence API
- [x] Set up authorization
- [x] Make an annotate video request
- [x] Congratulations!

# Supplement
## Set up authorization
```sh
gcloud iam service-accounts create quickstart
gcloud iam service-accounts keys create key.json --iam-account quickstart@qwiklabs-gcp-02-20c90a4b5d8a.iam.gserviceaccount.com
gcloud auth activate-service-account --key-file key.json
```

## Make an annotate video request
```sh
cat <<EOF >request.json
{
  "inputUri":"gs://spls/gsp154/video/chicago.mp4",
  "features": [
      "LABEL_DETECTION"
   ]
}
EOF

ACCESS_TOKEN=$(gcloud auth print-access-token)
OPERATION_NAME=$(
  curl -s -H 'Content-Type: application/json' \
    -H "Authorization: Bearer ${ACCESS_TOKEN}" \
    'https://videointelligence.googleapis.com/v1/videos:annotate' \
    -d @request.json \
  | jq -r .name
)

curl -s -H 'Content-Type: application/json' \
  -H "Authorization: Bearer ${ACCESS_TOKEN}" \
  "https://videointelligence.googleapis.com/v1/operations/${OPERATION_NAME}"
  
# Execute again after a minute
curl -s -H 'Content-Type: application/json' \
  -H "Authorization: Bearer ${ACCESS_TOKEN}" \
  "https://videointelligence.googleapis.com/v1/operations/${OPERATION_NAME}"
```

# References
- https://cloud.google.com/video-intelligence/
