---
title: "A Tour of Qwiklabs and the Google Cloud Platform"
tags: google-cloud-platform
url: https://www.qwiklabs.com/focuses/2794
---

# Goal
- Learn about the Qwiklabs platform and identify key features of a lab environment.
- Learn about (and possibly purchase) Qwiklabs credits and launch an instance of a lab.
- Learn how to access the GCP console with specific credentials.
- Learn about GCP projects and identify common misconceptions that surround them.
- Learn how to use the GCP navigation menu to identify types of GCP services.
- Learn about primitive roles and use the Cloud IAM service to inspect actions available to specific users.
- Learn about Cloud Shell and run commands that use the gcloud toolkit.
- Learn about the API library and examine its chief features.
- Use tools that are pre-installed in Cloud Shell and run commands like touch, nano, and cat to create, edit, and output the content of files.

# Task
- [x] Accessing the GCP Console
- [x] Projects in the GCP Console
- [x] Navigation Menu and Services
- [x] Roles and Permissions
- [x] APIs and Services
- [x] Cloud Shell

# Supplement
## Roles and Permissions
- https://cloud.google.com/iam/
- https://cloud.google.com/iam/docs/understanding-roles

## APIs and Services
- https://cloud.google.com/apis/design/
- https://developers.google.com/apis-explorer/#p/

## Cloud Shell
- https://cloud.google.com/shell/docs/features

```sh
gcloud auth list
touch test.txt
ls

cat <<EOF >test.txt
GCP and Qwiklabs are the best!
EOF

cat test.txt
```
