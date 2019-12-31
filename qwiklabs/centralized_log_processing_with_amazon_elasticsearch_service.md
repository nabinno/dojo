---
title: Centralized Log Processing with Amazon Elasticsearch Service
tags: amazon-elasticsearch-service,elasticsearch
url: https://www.qwiklabs.com/focuses/8563?catalog_rank=%7B%22rank%22%3A2%2C%22num_filters%22%3A0%2C%22has_search%22%3Atrue%7D&parent=catalog&search_id=4126136
---

# Goal
- Deploy an Amazon ES cluster
- Create user profiles using an Amazon Cognito user pool
- Set up an AWS Lambda function to pull logs into Amazon ES
- Analyze the logs in the Kibana dashboard

# Task
- [x] Lab overview
- [x] Start Lab
- [x] Task 1: Exploring the details of the Amazon ES cluster
- [x] Task 2: Setting up a user profile for Kibana access
- [x] Task 3: Setting up a Lambda function to stream CloudWatch logs to the Amazon ES cluster
- [x] Task 4: Streaming logs to the Amazon ES cluster
- [x] Task 5: Exploring logs using the Kibana dashboard
- [x] Task 6: Log discovery and analytics
- [x] Conclusion
- [x] End Lab
- [x] Additional resources
- [x] Appendix

# Supplement
## Task 3: Setting up a Lambda function to stream CloudWatch logs to the Amazon ES cluster
**Runtime**
Node.js 10.x

**Function Code**
https://s3-us-west-2.amazonaws.com/us-west-2-aws-training/awsu-spl/spl-237/scripts/indexing-service.zip

**Environmental Variable**
DOMAIN_ENDPOINT: https://vpc-loganalysiswithes-hokjtyjj5rm6lwyuiylnzsbtci.us-west-2.es.amazonaws.com

**Basic configuration**
- Memory: 512 MB
- Timeout: 2 min

**Network**
- VPC: vpc-xxxx | Log Analytics - VPC
- Subnet: subnet-xxxx
- Security Group: sg-xxxx | Lambda-SG

## Task 4: Streaming logs to the Amazon ES cluster
**Log format**
```
# Apache HTTP Server
[host, ident, authuser, date, request, status, bytes, referrer, agent]

# SSHD
[month, day, timestamp, destIp, id, msg1, msg2, msg3, srcIp, msg5, destPort, msg7]
```

## Task 5: Exploring logs using the Kibana dashboard
**Remote Desktop Connection**
```
Public DNS: ec2-34-219-63-217.us-west-2.compute.amazonaws.com
User Name: Administrator
Password: QbSC$m6$$Rp43U@-?hdO&*nhQzC6?39G
```

**Kibana URL**
https://vpc-loganalysiswithes-hokjtyjj5rm6lwyuiylnzsbtci.us-west-2.es.amazonaws.com/_plugin/kibana/

**Kibana Dashboard Object**
```json
[
  {
    "_id": "cb40cf10-8e54-11e9-957d-17ad3d0a703a",
    "_type": "dashboard",
    "_source": {
      "title": "Advanced_Dashboard",
      "hits": 0,
      "description": "",
      "panelsJSON": "[{\"embeddableConfig\":{},\"gridData\":{\"x\":23,\"y\":0,\"w\":25,\"h\":12,\"i\":\"2\"},\"id\":\"4f403ac0-8e2e-11e9-957d-17ad3d0a703a\",\"panelIndex\":\"2\",\"type\":\"visualization\",\"version\":\"6.7.0\"},{\"embeddableConfig\":{},\"gridData\":{\"x\":0,\"y\":12,\"w\":23,\"h\":13,\"i\":\"5\"},\"id\":\"3a5fc3c0-8e3b-11e9-b78d-49a184b5bb6e\",\"panelIndex\":\"5\",\"type\":\"visualization\",\"version\":\"6.7.0\"},{\"embeddableConfig\":{},\"gridData\":{\"x\":0,\"y\":25,\"w\":23,\"h\":9,\"i\":\"6\"},\"id\":\"b87c9ea0-8e30-11e9-957d-17ad3d0a703a\",\"panelIndex\":\"6\",\"type\":\"visualization\",\"version\":\"6.7.0\"},{\"embeddableConfig\":{\"vis\":{\"legendOpen\":true}},\"gridData\":{\"x\":23,\"y\":12,\"w\":25,\"h\":13,\"i\":\"7\"},\"id\":\"cd19f350-8e3d-11e9-b78d-49a184b5bb6e\",\"panelIndex\":\"7\",\"type\":\"visualization\",\"version\":\"6.7.0\"},{\"embeddableConfig\":{},\"gridData\":{\"x\":0,\"y\":0,\"w\":23,\"h\":12,\"i\":\"10\"},\"id\":\"64895970-8e33-11e9-b78d-49a184b5bb6e\",\"panelIndex\":\"10\",\"type\":\"visualization\",\"version\":\"6.7.0\"},{\"embeddableConfig\":{},\"gridData\":{\"x\":23,\"y\":25,\"w\":25,\"h\":9,\"i\":\"11\"},\"id\":\"Top-10-Access-Keys\",\"panelIndex\":\"11\",\"type\":\"visualization\",\"version\":\"6.7.0\"},{\"gridData\":{\"x\":0,\"y\":34,\"w\":48,\"h\":12,\"i\":\"12\"},\"version\":\"6.7.0\",\"panelIndex\":\"12\",\"type\":\"search\",\"id\":\"ddac9500-8e25-11e9-957d-17ad3d0a703a\",\"embeddableConfig\":{}}]",
      "optionsJSON": "{\"darkTheme\":true,\"hidePanelTitles\":false,\"useMargins\":true}",
      "version": 1,
      "timeRestore": true,
      "timeTo": "now",
      "timeFrom": "now-24h",
      "refreshInterval": {
        "pause": false,
        "value": 30000
      },
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"query\":{\"language\":\"lucene\",\"query\":\"\"},\"filter\":[]}"
      }
    }
  },
  {
    "_id": "4fffdab0-8e01-11e9-957d-17ad3d0a703a",
    "_type": "index-pattern",
    "_source": {
      "title": "cwl-*",
      "timeFieldName": "@timestamp",
      "fields": "[{\"name\":\"@id\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"@id.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"@log_group\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"@log_group.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"@log_stream\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"@log_stream.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"@message\",\"type\":\"string\",\"count\":2,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"@message.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"@owner\",\"type\":\"string\",\"count\":1,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"@owner.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"@timestamp\",\"type\":\"date\",\"count\":2,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"_id\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":false},{\"name\":\"_index\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":false},{\"name\":\"_score\",\"type\":\"number\",\"count\":0,\"scripted\":false,\"searchable\":false,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"_source\",\"type\":\"_source\",\"count\":0,\"scripted\":false,\"searchable\":false,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"_type\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":false},{\"name\":\"account_id\",\"type\":\"number\",\"count\":2,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"action\",\"type\":\"string\",\"count\":3,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"action.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"additionalEventData.AuthenticationMethod\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"additionalEventData.AuthenticationMethod.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"additionalEventData.CipherSuite\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"additionalEventData.CipherSuite.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"additionalEventData.SignatureVersion\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"additionalEventData.SignatureVersion.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"agent\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"agent.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"apiVersion\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"apiVersion.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"authuser\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"authuser.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"awsRegion\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"awsRegion.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"bytes\",\"type\":\"number\",\"count\":2,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"date\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"date.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"day\",\"type\":\"number\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"destIp\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"destIp.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"destPort\",\"type\":\"number\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"dstaddr\",\"type\":\"string\",\"count\":3,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"dstaddr.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"dstport\",\"type\":\"number\",\"count\":5,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"end\",\"type\":\"number\",\"count\":2,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"errorCode\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"errorCode.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"errorMessage\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"errorMessage.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"eventID\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"eventID.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"eventName\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"eventName.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"eventSource\",\"type\":\"string\",\"count\":1,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"eventSource.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"eventTime\",\"type\":\"date\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"eventType\",\"type\":\"string\",\"count\":1,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"eventType.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"eventVersion\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"eventVersion.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"host\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"host.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"id\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"id.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"ident\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"ident.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"interface_id\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"interface_id.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"log_status\",\"type\":\"string\",\"count\":4,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"log_status.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"month\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"month.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"msg1\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"msg1.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"msg2\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"msg2.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"msg3\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"msg3.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"msg5\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"msg5.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"msg7\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"msg7.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"packets\",\"type\":\"number\",\"count\":2,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"protocol\",\"type\":\"number\",\"count\":1,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"readOnly\",\"type\":\"boolean\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"recipientAccountId\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"recipientAccountId.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"referrer\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"referrer.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"request\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"request.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"requestID\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"requestID.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"requestParameters\",\"type\":\"string\",\"count\":1,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"requestParameters.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"resources.ARN\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"resources.ARN.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"resources.accountId\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"resources.accountId.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"resources.type\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"resources.type.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"responseElements\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"responseElements.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"sharedEventID\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"sharedEventID.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"sourceIPAddress\",\"type\":\"string\",\"count\":1,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"sourceIPAddress.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"srcIp\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"srcIp.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"srcaddr\",\"type\":\"string\",\"count\":3,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"srcaddr.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"srcport\",\"type\":\"number\",\"count\":1,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"start\",\"type\":\"number\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"status\",\"type\":\"number\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"timestamp\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"timestamp.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"userAgent\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"userAgent.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"userIdentity.accessKeyId\",\"type\":\"string\",\"count\":1,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"userIdentity.accessKeyId.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"userIdentity.accountId\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"userIdentity.accountId.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"userIdentity.arn\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"userIdentity.arn.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"userIdentity.invokedBy\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"userIdentity.invokedBy.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"userIdentity.principalId\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"userIdentity.principalId.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"userIdentity.sessionContext.attributes.creationDate\",\"type\":\"date\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"userIdentity.sessionContext.attributes.mfaAuthenticated\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"userIdentity.sessionContext.attributes.mfaAuthenticated.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"userIdentity.sessionContext.sessionIssuer.accountId\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"userIdentity.sessionContext.sessionIssuer.accountId.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"userIdentity.sessionContext.sessionIssuer.arn\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"userIdentity.sessionContext.sessionIssuer.arn.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"userIdentity.sessionContext.sessionIssuer.principalId\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"userIdentity.sessionContext.sessionIssuer.principalId.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"userIdentity.sessionContext.sessionIssuer.type\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"userIdentity.sessionContext.sessionIssuer.type.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"userIdentity.sessionContext.sessionIssuer.userName\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"userIdentity.sessionContext.sessionIssuer.userName.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"userIdentity.type\",\"type\":\"string\",\"count\":1,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"userIdentity.type.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"userIdentity.userName\",\"type\":\"string\",\"count\":1,\"scripted\":false,\"searchable\":true,\"aggregatable\":false,\"readFromDocValues\":false},{\"name\":\"userIdentity.userName.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true},{\"name\":\"version\",\"type\":\"number\",\"count\":2,\"scripted\":false,\"searchable\":true,\"aggregatable\":true,\"readFromDocValues\":true}]"
    }
  },
  {
    "_id": "Accept",
    "_type": "search",
    "_source": {
      "title": "Accept",
      "description": "",
      "hits": 0,
      "columns": [
        "_source"
      ],
      "sort": [
        "@timestamp",
        "desc"
      ],
      "version": 1,
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"index\":\"4fffdab0-8e01-11e9-957d-17ad3d0a703a\",\"highlight\":{\"pre_tags\":[\"@kibana-highlighted-field@\"],\"post_tags\":[\"@/kibana-highlighted-field@\"],\"fields\":{\"*\":{}},\"fragment_size\":2147483647},\"filter\":[],\"query\":{\"query\":{\"query_string\":{\"query\":\"action:accept\",\"analyze_wildcard\":true}},\"language\":\"lucene\"}}"
      }
    }
  },
  {
    "_id": "HTTP-200-Code",
    "_type": "search",
    "_source": {
      "title": "HTTP 200 Code",
      "description": "",
      "hits": 0,
      "columns": [
        "_source"
      ],
      "sort": [
        "@timestamp",
        "desc"
      ],
      "version": 1,
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"index\":\"4fffdab0-8e01-11e9-957d-17ad3d0a703a\",\"query\":{\"query\":{\"query_string\":{\"query\":\"status:200\",\"analyze_wildcard\":true}},\"language\":\"lucene\"},\"highlight\":{\"pre_tags\":[\"@kibana-highlighted-field@\"],\"post_tags\":[\"@/kibana-highlighted-field@\"],\"fields\":{\"*\":{}},\"fragment_size\":2147483647},\"filter\":[]}"
      }
    }
  },
  {
    "_id": "HTTP-404-Code",
    "_type": "search",
    "_source": {
      "title": "HTTP 404 Code",
      "description": "",
      "hits": 0,
      "columns": [
        "_source"
      ],
      "sort": [
        "@timestamp",
        "desc"
      ],
      "version": 1,
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"index\":\"4fffdab0-8e01-11e9-957d-17ad3d0a703a\",\"query\":{\"query\":{\"query_string\":{\"query\":\"status:404\",\"analyze_wildcard\":true}},\"language\":\"lucene\"},\"highlight\":{\"pre_tags\":[\"@kibana-highlighted-field@\"],\"post_tags\":[\"@/kibana-highlighted-field@\"],\"fields\":{\"*\":{}},\"fragment_size\":2147483647},\"filter\":[]}"
      }
    }
  },
  {
    "_id": "All-Data",
    "_type": "search",
    "_source": {
      "title": "All Data",
      "description": "",
      "hits": 0,
      "columns": [
        "_source"
      ],
      "sort": [
        "@timestamp",
        "desc"
      ],
      "version": 1,
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"index\":\"4fffdab0-8e01-11e9-957d-17ad3d0a703a\",\"highlight\":{\"pre_tags\":[\"@kibana-highlighted-field@\"],\"post_tags\":[\"@/kibana-highlighted-field@\"],\"fields\":{\"*\":{}},\"fragment_size\":2147483647},\"filter\":[],\"query\":{\"query\":{\"query_string\":{\"query\":\"*\",\"analyze_wildcard\":true}},\"language\":\"lucene\"}}"
      }
    }
  },
  {
    "_id": "Reject",
    "_type": "search",
    "_source": {
      "title": "Reject",
      "description": "",
      "hits": 0,
      "columns": [
        "_source"
      ],
      "sort": [
        "@timestamp",
        "desc"
      ],
      "version": 1,
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"index\":\"4fffdab0-8e01-11e9-957d-17ad3d0a703a\",\"highlight\":{\"pre_tags\":[\"@kibana-highlighted-field@\"],\"post_tags\":[\"@/kibana-highlighted-field@\"],\"fields\":{\"*\":{}},\"fragment_size\":2147483647},\"filter\":[],\"query\":{\"query\":{\"query_string\":{\"query\":\"action:reject\",\"analyze_wildcard\":true}},\"language\":\"lucene\"}}"
      }
    }
  },
  {
    "_id": "ddac9500-8e25-11e9-957d-17ad3d0a703a",
    "_type": "search",
    "_source": {
      "title": "VPCFlowLogs_REJECTS_Traffic",
      "description": "",
      "hits": 0,
      "columns": [
        "srcaddr",
        "srcport",
        "dstaddr",
        "dstport",
        "action",
        "protocol"
      ],
      "sort": [
        "@timestamp",
        "desc"
      ],
      "version": 1,
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"index\":\"4fffdab0-8e01-11e9-957d-17ad3d0a703a\",\"highlightAll\":true,\"version\":true,\"query\":{\"language\":\"lucene\",\"query\":\"\"},\"filter\":[{\"$state\":{\"store\":\"appState\"},\"meta\":{\"alias\":null,\"disabled\":false,\"index\":\"4fffdab0-8e01-11e9-957d-17ad3d0a703a\",\"key\":\"@log_group.keyword\",\"negate\":false,\"params\":{\"query\":\"VPCFlowLogGroup\",\"type\":\"phrase\"},\"type\":\"phrase\",\"value\":\"VPCFlowLogGroup\"},\"query\":{\"match\":{\"@log_group.keyword\":{\"query\":\"VPCFlowLogGroup\",\"type\":\"phrase\"}}}},{\"$state\":{\"store\":\"appState\"},\"meta\":{\"alias\":null,\"disabled\":false,\"index\":\"4fffdab0-8e01-11e9-957d-17ad3d0a703a\",\"key\":\"action.keyword\",\"negate\":false,\"params\":{\"query\":\"REJECT\",\"type\":\"phrase\"},\"type\":\"phrase\",\"value\":\"REJECT\"},\"query\":{\"match\":{\"action.keyword\":{\"query\":\"REJECT\",\"type\":\"phrase\"}}}}]}"
      }
    }
  },
  {
    "_id": "f1709820-f182-11e7-8ed1-e93cb1e71cec",
    "_type": "visualization",
    "_source": {
      "title": "Top Account Ids",
      "visState": "{\"title\":\"Top Account Ids\",\"type\":\"table\",\"params\":{\"perPage\":10,\"showPartialRows\":false,\"showMeticsAtAllLevels\":false,\"sort\":{\"columnIndex\":null,\"direction\":null},\"showTotal\":false,\"totalFunc\":\"sum\"},\"aggs\":[{\"id\":\"1\",\"enabled\":true,\"type\":\"count\",\"schema\":\"metric\",\"params\":{}},{\"id\":\"2\",\"enabled\":true,\"type\":\"terms\",\"schema\":\"bucket\",\"params\":{\"field\":\"account_id\",\"size\":5,\"order\":\"desc\",\"orderBy\":\"1\"}}],\"listeners\":{}}",
      "uiStateJSON": "{\"vis\":{\"params\":{\"sort\":{\"columnIndex\":null,\"direction\":null}}}}",
      "description": "",
      "version": 1,
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"index\":\"4fffdab0-8e01-11e9-957d-17ad3d0a703a\",\"query\":{\"query\":{\"match_all\":{}},\"language\":\"lucene\"},\"filter\":[]}"
      }
    }
  },
  {
    "_id": "74a9bbe0-f183-11e7-8ed1-e93cb1e71cec",
    "_type": "visualization",
    "_source": {
      "title": "Top Regions",
      "visState": "{\"title\":\"Top Regions\",\"type\":\"table\",\"params\":{\"perPage\":10,\"showPartialRows\":false,\"showMeticsAtAllLevels\":false,\"sort\":{\"columnIndex\":null,\"direction\":null},\"showTotal\":false,\"totalFunc\":\"sum\"},\"aggs\":[{\"id\":\"1\",\"enabled\":true,\"type\":\"count\",\"schema\":\"metric\",\"params\":{}},{\"id\":\"2\",\"enabled\":true,\"type\":\"terms\",\"schema\":\"bucket\",\"params\":{\"field\":\"awsRegion.keyword\",\"size\":5,\"order\":\"desc\",\"orderBy\":\"1\",\"customLabel\":\"awsRegion\"}}],\"listeners\":{}}",
      "uiStateJSON": "{\"vis\":{\"params\":{\"sort\":{\"columnIndex\":null,\"direction\":null}}}}",
      "description": "",
      "version": 1,
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"index\":\"4fffdab0-8e01-11e9-957d-17ad3d0a703a\",\"query\":{\"query\":{\"match_all\":{}},\"language\":\"lucene\"},\"filter\":[]}"
      }
    }
  },
  {
    "_id": "Total-Packets-Src-Address",
    "_type": "visualization",
    "_source": {
      "title": "Total Packets Src Address",
      "visState": "{\"title\":\"Total Packets Src Address\",\"type\":\"pie\",\"params\":{\"shareYAxis\":true,\"addTooltip\":true,\"addLegend\":false,\"isDonut\":false,\"legendPosition\":\"right\"},\"aggs\":[{\"id\":\"1\",\"enabled\":true,\"type\":\"sum\",\"schema\":\"metric\",\"params\":{\"field\":\"packets\"}},{\"id\":\"2\",\"enabled\":true,\"type\":\"terms\",\"schema\":\"segment\",\"params\":{\"field\":\"srcaddr.keyword\",\"size\":10,\"order\":\"desc\",\"orderBy\":\"1\",\"customLabel\":\"\"}}],\"listeners\":{}}",
      "uiStateJSON": "{\"vis\":{\"legendOpen\":true}}",
      "description": "",
      "savedSearchId": "All-Data",
      "version": 1,
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"filter\":[]}"
      }
    }
  },
  {
    "_id": "Packets-Reject",
    "_type": "visualization",
    "_source": {
      "title": "Packets - Reject",
      "visState": "{\"type\":\"line\",\"params\":{\"shareYAxis\":true,\"addTooltip\":true,\"addLegend\":true,\"showCircles\":true,\"smoothLines\":false,\"interpolate\":\"linear\",\"scale\":\"linear\",\"drawLinesBetweenPoints\":true,\"radiusRatio\":9,\"times\":[],\"addTimeMarker\":false,\"defaultYExtents\":false,\"setYExtents\":false,\"yAxis\":{}},\"aggs\":[{\"id\":\"1\",\"type\":\"sum\",\"schema\":\"metric\",\"params\":{\"field\":\"packets\"}},{\"id\":\"2\",\"type\":\"date_histogram\",\"schema\":\"segment\",\"params\":{\"field\":\"@timestamp\",\"interval\":\"custom\",\"customInterval\":\"30s\",\"min_doc_count\":1,\"extended_bounds\":{}}}],\"listeners\":{},\"title\":\"Packets - Reject\"}",
      "uiStateJSON": "{}",
      "description": "",
      "savedSearchId": "Reject",
      "version": 1,
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"filter\":[]}"
      }
    }
  },
  {
    "_id": "Packets-Accept",
    "_type": "visualization",
    "_source": {
      "title": "Packets - Accept",
      "visState": "{\"type\":\"line\",\"params\":{\"shareYAxis\":true,\"addTooltip\":true,\"addLegend\":true,\"showCircles\":true,\"smoothLines\":false,\"interpolate\":\"linear\",\"scale\":\"linear\",\"drawLinesBetweenPoints\":true,\"radiusRatio\":9,\"times\":[],\"addTimeMarker\":false,\"defaultYExtents\":false,\"setYExtents\":false,\"yAxis\":{}},\"aggs\":[{\"id\":\"1\",\"type\":\"sum\",\"schema\":\"metric\",\"params\":{\"field\":\"packets\"}},{\"id\":\"2\",\"type\":\"date_histogram\",\"schema\":\"segment\",\"params\":{\"field\":\"@timestamp\",\"interval\":\"custom\",\"customInterval\":\"30s\",\"min_doc_count\":1,\"extended_bounds\":{}}}],\"listeners\":{},\"title\":\"Packets - Accept\"}",
      "uiStateJSON": "{}",
      "description": "",
      "savedSearchId": "Accept",
      "version": 1,
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"filter\":[]}"
      }
    }
  },
  {
    "_id": "Top-10-Event-Sources",
    "_type": "visualization",
    "_source": {
      "title": "Top 10 Event Sources",
      "visState": "{\"title\":\"Top 10 Event Sources\",\"type\":\"table\",\"params\":{\"perPage\":10,\"showPartialRows\":false,\"showMeticsAtAllLevels\":false,\"sort\":{\"columnIndex\":null,\"direction\":null},\"showTotal\":false,\"totalFunc\":\"sum\"},\"aggs\":[{\"id\":\"1\",\"enabled\":true,\"type\":\"count\",\"schema\":\"metric\",\"params\":{}},{\"id\":\"2\",\"enabled\":true,\"type\":\"terms\",\"schema\":\"bucket\",\"params\":{\"field\":\"eventSource.keyword\",\"size\":10,\"order\":\"desc\",\"orderBy\":\"1\",\"customLabel\":\"eventSource\"}}],\"listeners\":{}}",
      "uiStateJSON": "{\"vis\":{\"params\":{\"sort\":{\"columnIndex\":null,\"direction\":null}}}}",
      "description": "",
      "savedSearchId": "All-Data",
      "version": 1,
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"filter\":[]}"
      }
    }
  },
  {
    "_id": "Reject-Packets-Src-Address",
    "_type": "visualization",
    "_source": {
      "title": "Reject Packets Src Address",
      "visState": "{\"title\":\"Reject Packets Src Address\",\"type\":\"pie\",\"params\":{\"shareYAxis\":true,\"addTooltip\":true,\"addLegend\":false,\"isDonut\":false,\"legendPosition\":\"right\"},\"aggs\":[{\"id\":\"1\",\"enabled\":true,\"type\":\"sum\",\"schema\":\"metric\",\"params\":{\"field\":\"packets\"}},{\"id\":\"2\",\"enabled\":true,\"type\":\"terms\",\"schema\":\"segment\",\"params\":{\"field\":\"srcaddr.keyword\",\"size\":10,\"order\":\"desc\",\"orderBy\":\"1\"}}],\"listeners\":{}}",
      "uiStateJSON": "{}",
      "description": "",
      "savedSearchId": "Reject",
      "version": 1,
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"filter\":[]}"
      }
    }
  },
  {
    "_id": "Top-10-Src-Address-Packets-Total",
    "_type": "visualization",
    "_source": {
      "title": "Total Src Addr",
      "visState": "{\"title\":\"Total Src Addr\",\"type\":\"table\",\"params\":{\"perPage\":10,\"showPartialRows\":false,\"showMeticsAtAllLevels\":false,\"sort\":{\"columnIndex\":null,\"direction\":null},\"showTotal\":false,\"totalFunc\":\"sum\"},\"aggs\":[{\"id\":\"1\",\"enabled\":true,\"type\":\"sum\",\"schema\":\"metric\",\"params\":{\"field\":\"packets\"}},{\"id\":\"2\",\"enabled\":true,\"type\":\"terms\",\"schema\":\"bucket\",\"params\":{\"field\":\"srcaddr.keyword\",\"size\":10,\"order\":\"desc\",\"orderBy\":\"1\",\"customLabel\":\"srcaddr\"}}],\"listeners\":{}}",
      "uiStateJSON": "{\"vis\":{\"params\":{\"sort\":{\"columnIndex\":null,\"direction\":null}}}}",
      "description": "",
      "savedSearchId": "All-Data",
      "version": 1,
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"filter\":[]}"
      }
    }
  },
  {
    "_id": "f9f59430-8e36-11e9-957d-17ad3d0a703a",
    "_type": "visualization",
    "_source": {
      "title": "Service Use By Region",
      "visState": "{\"title\":\"Service Use By Region\",\"type\":\"table\",\"params\":{\"perPage\":10,\"showPartialRows\":false,\"showMetricsAtAllLevels\":false,\"sort\":{\"columnIndex\":null,\"direction\":null},\"showTotal\":false,\"totalFunc\":\"sum\"},\"aggs\":[{\"id\":\"1\",\"enabled\":true,\"type\":\"count\",\"schema\":\"metric\",\"params\":{}},{\"id\":\"2\",\"enabled\":true,\"type\":\"terms\",\"schema\":\"bucket\",\"params\":{\"field\":\"eventSource.keyword\",\"size\":10,\"order\":\"desc\",\"orderBy\":\"1\",\"otherBucket\":false,\"otherBucketLabel\":\"Other\",\"missingBucket\":false,\"missingBucketLabel\":\"Missing\",\"customLabel\":\"Top Services\"}},{\"id\":\"3\",\"enabled\":true,\"type\":\"terms\",\"schema\":\"bucket\",\"params\":{\"field\":\"awsRegion.keyword\",\"size\":10,\"order\":\"desc\",\"orderBy\":\"1\",\"otherBucket\":false,\"otherBucketLabel\":\"Other\",\"missingBucket\":false,\"missingBucketLabel\":\"Missing\",\"customLabel\":\"Region\"}}]}",
      "uiStateJSON": "{\"vis\":{\"params\":{\"sort\":{\"columnIndex\":null,\"direction\":null}}}}",
      "description": "",
      "version": 1,
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"index\":\"4fffdab0-8e01-11e9-957d-17ad3d0a703a\",\"query\":{\"query\":\"\",\"language\":\"lucene\"},\"filter\":[{\"meta\":{\"index\":\"4fffdab0-8e01-11e9-957d-17ad3d0a703a\",\"negate\":false,\"disabled\":false,\"alias\":null,\"type\":\"phrase\",\"key\":\"@log_group.keyword\",\"value\":\"CloudTrailLogGroup\",\"params\":{\"query\":\"CloudTrailLogGroup\",\"type\":\"phrase\"}},\"query\":{\"match\":{\"@log_group.keyword\":{\"query\":\"CloudTrailLogGroup\",\"type\":\"phrase\"}}},\"$state\":{\"store\":\"appState\"}}]}"
      }
    }
  },
  {
    "_id": "Top-10-Src-Address-Packets-Reject",
    "_type": "visualization",
    "_source": {
      "title": "Reject Src Addr",
      "visState": "{\"title\":\"Reject Src Addr\",\"type\":\"table\",\"params\":{\"perPage\":10,\"showPartialRows\":true,\"sort\":{\"columnIndex\":null,\"direction\":null},\"showTotal\":false,\"totalFunc\":\"sum\",\"showMetricsAtAllLevels\":false},\"aggs\":[{\"id\":\"1\",\"enabled\":true,\"type\":\"sum\",\"schema\":\"metric\",\"params\":{\"field\":\"packets\"}},{\"id\":\"2\",\"enabled\":true,\"type\":\"terms\",\"schema\":\"bucket\",\"params\":{\"field\":\"srcaddr.keyword\",\"size\":10,\"order\":\"desc\",\"orderBy\":\"1\",\"otherBucket\":false,\"otherBucketLabel\":\"Other\",\"missingBucket\":false,\"missingBucketLabel\":\"Missing\",\"customLabel\":\"srcaddr\"}}]}",
      "uiStateJSON": "{\"vis\":{\"params\":{\"sort\":{\"columnIndex\":null,\"direction\":null}}}}",
      "description": "",
      "savedSearchId": "Reject",
      "version": 1,
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"filter\":[],\"query\":{\"query\":\"\",\"language\":\"lucene\"}}"
      }
    }
  },
  {
    "_id": "64895970-8e33-11e9-b78d-49a184b5bb6e",
    "_type": "visualization",
    "_source": {
      "title": "HTTP Response Codes",
      "visState": "{\"title\":\"HTTP Response Codes\",\"type\":\"gauge\",\"params\":{\"type\":\"gauge\",\"addTooltip\":true,\"addLegend\":false,\"isDisplayWarning\":false,\"gauge\":{\"verticalSplit\":false,\"extendRange\":true,\"percentageMode\":false,\"gaugeType\":\"Arc\",\"gaugeStyle\":\"Full\",\"backStyle\":\"Full\",\"orientation\":\"vertical\",\"colorSchema\":\"Yellow to Red\",\"gaugeColorMode\":\"Labels\",\"colorsRange\":[{\"from\":0,\"to\":25000},{\"from\":25000,\"to\":50000},{\"from\":50000,\"to\":100000}],\"invertColors\":false,\"labels\":{\"show\":true,\"color\":\"black\"},\"scale\":{\"show\":true,\"labels\":false,\"color\":\"#333\"},\"type\":\"meter\",\"style\":{\"bgWidth\":0.9,\"width\":0.9,\"mask\":false,\"bgMask\":false,\"maskBars\":50,\"bgFill\":\"#eee\",\"bgColor\":false,\"subText\":\"\",\"fontSize\":60,\"labelColor\":true}}},\"aggs\":[{\"id\":\"1\",\"enabled\":true,\"type\":\"count\",\"schema\":\"metric\",\"params\":{}},{\"id\":\"2\",\"enabled\":true,\"type\":\"terms\",\"schema\":\"group\",\"params\":{\"field\":\"status\",\"size\":5,\"order\":\"desc\",\"orderBy\":\"1\",\"otherBucket\":false,\"otherBucketLabel\":\"Other\",\"missingBucket\":false,\"missingBucketLabel\":\"Missing\",\"customLabel\":\"Response Code\"}}]}",
      "uiStateJSON": "{\"vis\":{\"defaultColors\":{\"0 - 25000\":\"rgb(255,255,204)\",\"25000 - 50000\":\"rgb(253,141,60)\",\"50000 - 100000\":\"rgb(128,0,38)\"},\"legendOpen\":true}}",
      "description": "",
      "version": 1,
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"index\":\"4fffdab0-8e01-11e9-957d-17ad3d0a703a\",\"query\":{\"language\":\"lucene\",\"query\":\"\"},\"filter\":[{\"$state\":{\"store\":\"appState\"},\"meta\":{\"alias\":null,\"disabled\":false,\"index\":\"4fffdab0-8e01-11e9-957d-17ad3d0a703a\",\"key\":\"@log_group.keyword\",\"negate\":false,\"params\":{\"query\":\"AccessLogGroup\",\"type\":\"phrase\"},\"type\":\"phrase\",\"value\":\"AccessLogGroup\"},\"query\":{\"match\":{\"@log_group.keyword\":{\"query\":\"AccessLogGroup\",\"type\":\"phrase\"}}}}]}"
      }
    }
  },
  {
    "_id": "4f403ac0-8e2e-11e9-957d-17ad3d0a703a",
    "_type": "visualization",
    "_source": {
      "title": "Traffic in real-time",
      "visState": "{\"title\":\"Traffic in real-time\",\"type\":\"metrics\",\"params\":{\"id\":\"61ca57f0-469d-11e7-af02-69e470af7417\",\"type\":\"timeseries\",\"series\":[{\"id\":\"61ca57f1-469d-11e7-af02-69e470af7417\",\"color\":\"rgba(0,125,188,1)\",\"split_mode\":\"terms\",\"metrics\":[{\"id\":\"61ca57f2-469d-11e7-af02-69e470af7417\",\"type\":\"count\"}],\"separate_axis\":0,\"axis_position\":\"right\",\"formatter\":\"number\",\"chart_type\":\"line\",\"line_width\":1,\"point_size\":1,\"fill\":0.5,\"stacked\":\"none\",\"terms_field\":\"srcaddr.keyword\",\"terms_order_by\":\"_count\",\"split_color_mode\":\"rainbow\",\"label\":\"Traffic in real-time\"}],\"time_field\":\"@timestamp\",\"index_pattern\":\"\",\"interval\":\"auto\",\"axis_position\":\"left\",\"axis_formatter\":\"number\",\"axis_scale\":\"normal\",\"show_legend\":1,\"show_grid\":1,\"default_index_pattern\":\"cwl-*\",\"background_color\":\"rgba(43,41,41,1)\",\"annotations\":[{\"id\":\"b8835cc0-8e2d-11e9-ad88-9507a00c9c78\",\"color\":\"#F00\",\"index_pattern\":\"*\",\"time_field\":\"@timestamp\",\"icon\":\"fa-tag\",\"ignore_global_filters\":1,\"ignore_panel_filters\":1}],\"background_color_rules\":[{\"id\":\"c54385c0-8e2d-11e9-ad88-9507a00c9c78\"}],\"bar_color_rules\":[{\"id\":\"c9e0dba0-8e2d-11e9-ad88-9507a00c9c78\"}],\"gauge_color_rules\":[{\"id\":\"ccc16650-8e2d-11e9-ad88-9507a00c9c78\"}],\"gauge_width\":10,\"gauge_inner_width\":10,\"gauge_style\":\"half\"},\"aggs\":[]}",
      "uiStateJSON": "{}",
      "description": "",
      "version": 1,
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"query\":{\"query\":\"\",\"language\":\"lucene\"},\"filter\":[]}"
      }
    }
  },
  {
    "_id": "b87c9ea0-8e30-11e9-957d-17ad3d0a703a",
    "_type": "visualization",
    "_source": {
      "title": "SSH Attacks",
      "visState": "{\"title\":\"SSH Attacks\",\"type\":\"table\",\"params\":{\"perPage\":10,\"showPartialRows\":false,\"showMetricsAtAllLevels\":false,\"sort\":{\"columnIndex\":1,\"direction\":\"desc\"},\"showTotal\":false,\"totalFunc\":\"sum\"},\"aggs\":[{\"id\":\"1\",\"enabled\":true,\"type\":\"count\",\"schema\":\"metric\",\"params\":{}},{\"id\":\"2\",\"enabled\":true,\"type\":\"terms\",\"schema\":\"bucket\",\"params\":{\"field\":\"srcIp.keyword\",\"size\":4,\"order\":\"desc\",\"orderBy\":\"1\",\"otherBucket\":false,\"otherBucketLabel\":\"Other\",\"missingBucket\":false,\"missingBucketLabel\":\"Missing\",\"customLabel\":\"Source IP\"}}]}",
      "uiStateJSON": "{\"vis\":{\"params\":{\"sort\":{\"columnIndex\":1,\"direction\":\"desc\"}}}}",
      "description": "",
      "version": 1,
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"index\":\"4fffdab0-8e01-11e9-957d-17ad3d0a703a\",\"query\":{\"query\":\"\",\"language\":\"lucene\"},\"filter\":[{\"meta\":{\"index\":\"4fffdab0-8e01-11e9-957d-17ad3d0a703a\",\"negate\":false,\"disabled\":false,\"alias\":null,\"type\":\"phrase\",\"key\":\"@log_group.keyword\",\"value\":\"SSHLogGroup\",\"params\":{\"query\":\"SSHLogGroup\",\"type\":\"phrase\"}},\"query\":{\"match\":{\"@log_group.keyword\":{\"query\":\"SSHLogGroup\",\"type\":\"phrase\"}}},\"$state\":{\"store\":\"appState\"}},{\"meta\":{\"index\":\"4fffdab0-8e01-11e9-957d-17ad3d0a703a\",\"negate\":false,\"disabled\":false,\"alias\":null,\"type\":\"phrase\",\"key\":\"msg1.keyword\",\"value\":\"Connection\",\"params\":{\"query\":\"Connection\",\"type\":\"phrase\"}},\"query\":{\"match\":{\"msg1.keyword\":{\"query\":\"Connection\",\"type\":\"phrase\"}}},\"$state\":{\"store\":\"appState\"}}]}"
      }
    }
  },
  {
    "_id": "Top-10-Access-Keys",
    "_type": "visualization",
    "_source": {
      "title": "Top 10 Access Keys",
      "visState": "{\"title\":\"Top 10 Access Keys\",\"type\":\"table\",\"params\":{\"perPage\":5,\"showPartialRows\":false,\"sort\":{\"columnIndex\":null,\"direction\":null},\"showTotal\":false,\"totalFunc\":\"sum\",\"showMetricsAtAllLevels\":false},\"aggs\":[{\"id\":\"1\",\"enabled\":true,\"type\":\"count\",\"schema\":\"metric\",\"params\":{}},{\"id\":\"2\",\"enabled\":true,\"type\":\"terms\",\"schema\":\"bucket\",\"params\":{\"field\":\"userIdentity.accessKeyId.keyword\",\"size\":10,\"order\":\"desc\",\"orderBy\":\"1\",\"otherBucket\":false,\"otherBucketLabel\":\"Other\",\"missingBucket\":false,\"missingBucketLabel\":\"Missing\",\"customLabel\":\"Access Key Id\"}}]}",
      "uiStateJSON": "{\"vis\":{\"params\":{\"sort\":{\"columnIndex\":null,\"direction\":null}}}}",
      "description": "",
      "savedSearchId": "All-Data",
      "version": 1,
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"filter\":[],\"query\":{\"query\":\"\",\"language\":\"lucene\"}}"
      }
    }
  },
  {
    "_id": "HTTP-200-Code-Count",
    "_type": "visualization",
    "_source": {
      "title": "HTTP 200 Code Count",
      "visState": "{\"title\":\"HTTP 200 Code Count\",\"type\":\"metric\",\"params\":{\"fontSize\":60,\"addTooltip\":true,\"addLegend\":false,\"type\":\"metric\",\"metric\":{\"percentageMode\":false,\"useRanges\":false,\"colorSchema\":\"Green to Red\",\"metricColorMode\":\"None\",\"colorsRange\":[{\"from\":0,\"to\":10000}],\"labels\":{\"show\":false},\"invertColors\":false,\"style\":{\"bgFill\":\"#000\",\"bgColor\":false,\"labelColor\":false,\"subText\":\"\",\"fontSize\":60}}},\"aggs\":[{\"id\":\"1\",\"enabled\":true,\"type\":\"count\",\"schema\":\"metric\",\"params\":{}}]}",
      "uiStateJSON": "{}",
      "description": "",
      "savedSearchId": "HTTP-200-Code",
      "version": 1,
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"filter\":[],\"query\":{\"query\":\"\",\"language\":\"lucene\"}}"
      }
    }
  },
  {
    "_id": "HTTP-404-Code-Count",
    "_type": "visualization",
    "_source": {
      "title": "HTTP 404 Code Count",
      "visState": "{\"title\":\"HTTP 404 Code Count\",\"type\":\"metric\",\"params\":{\"fontSize\":60,\"addTooltip\":true,\"addLegend\":false,\"type\":\"metric\",\"metric\":{\"percentageMode\":false,\"useRanges\":false,\"colorSchema\":\"Green to Red\",\"metricColorMode\":\"None\",\"colorsRange\":[{\"from\":0,\"to\":10000}],\"labels\":{\"show\":false},\"invertColors\":false,\"style\":{\"bgFill\":\"#000\",\"bgColor\":false,\"labelColor\":false,\"subText\":\"\",\"fontSize\":60}}},\"aggs\":[{\"id\":\"1\",\"enabled\":true,\"type\":\"count\",\"schema\":\"metric\",\"params\":{}}]}",
      "uiStateJSON": "{}",
      "description": "",
      "savedSearchId": "HTTP-404-Code",
      "version": 1,
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"filter\":[],\"query\":{\"query\":\"\",\"language\":\"lucene\"}}"
      }
    }
  },
  {
    "_id": "3a5fc3c0-8e3b-11e9-b78d-49a184b5bb6e",
    "_type": "visualization",
    "_source": {
      "title": "Flow Logs Data",
      "visState": "{\"title\":\"Flow Logs Data\",\"type\":\"table\",\"params\":{\"perPage\":10,\"showMetricsAtAllLevels\":false,\"showPartialRows\":false,\"showTotal\":false,\"sort\":{\"columnIndex\":2,\"direction\":\"desc\"},\"totalFunc\":\"sum\"},\"aggs\":[{\"id\":\"1\",\"enabled\":true,\"type\":\"count\",\"schema\":\"metric\",\"params\":{}},{\"id\":\"3\",\"enabled\":true,\"type\":\"terms\",\"schema\":\"bucket\",\"params\":{\"field\":\"srcaddr.keyword\",\"size\":10,\"order\":\"desc\",\"orderBy\":\"1\",\"otherBucket\":true,\"otherBucketLabel\":\"Other\",\"missingBucket\":false,\"missingBucketLabel\":\"Missing\",\"customLabel\":\"Source IP\"}},{\"id\":\"2\",\"enabled\":true,\"type\":\"terms\",\"schema\":\"bucket\",\"params\":{\"field\":\"action.keyword\",\"size\":10,\"order\":\"desc\",\"orderBy\":\"1\",\"otherBucket\":false,\"otherBucketLabel\":\"Other\",\"missingBucket\":false,\"missingBucketLabel\":\"Missing\",\"customLabel\":\"Action\"}}]}",
      "uiStateJSON": "{\"vis\":{\"params\":{\"sort\":{\"columnIndex\":2,\"direction\":\"desc\"}}}}",
      "description": "",
      "version": 1,
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"index\":\"4fffdab0-8e01-11e9-957d-17ad3d0a703a\",\"query\":{\"language\":\"lucene\",\"query\":\"\"},\"filter\":[{\"$state\":{\"store\":\"appState\"},\"meta\":{\"alias\":null,\"disabled\":false,\"index\":\"4fffdab0-8e01-11e9-957d-17ad3d0a703a\",\"key\":\"@log_group.keyword\",\"negate\":false,\"params\":{\"query\":\"VPCFlowLogGroup\",\"type\":\"phrase\"},\"type\":\"phrase\",\"value\":\"VPCFlowLogGroup\"},\"query\":{\"match\":{\"@log_group.keyword\":{\"query\":\"VPCFlowLogGroup\",\"type\":\"phrase\"}}}}]}"
      }
    }
  },
  {
    "_id": "cd19f350-8e3d-11e9-b78d-49a184b5bb6e",
    "_type": "visualization",
    "_source": {
      "title": "AWS Services By Use",
      "visState": "{\"title\":\"AWS Services By Use\",\"type\":\"pie\",\"params\":{\"type\":\"pie\",\"addTooltip\":true,\"addLegend\":true,\"legendPosition\":\"right\",\"isDonut\":false,\"labels\":{\"show\":false,\"values\":true,\"last_level\":true,\"truncate\":100}},\"aggs\":[{\"id\":\"1\",\"enabled\":true,\"type\":\"count\",\"schema\":\"metric\",\"params\":{}},{\"id\":\"2\",\"enabled\":true,\"type\":\"terms\",\"schema\":\"segment\",\"params\":{\"field\":\"eventSource.keyword\",\"size\":10,\"order\":\"desc\",\"orderBy\":\"1\",\"otherBucket\":false,\"otherBucketLabel\":\"Other\",\"missingBucket\":false,\"missingBucketLabel\":\"Missing\"}},{\"id\":\"3\",\"enabled\":true,\"type\":\"terms\",\"schema\":\"segment\",\"params\":{\"field\":\"awsRegion.keyword\",\"size\":5,\"order\":\"desc\",\"orderBy\":\"1\",\"otherBucket\":false,\"otherBucketLabel\":\"Other\",\"missingBucket\":false,\"missingBucketLabel\":\"Missing\"}}]}",
      "uiStateJSON": "{}",
      "description": "",
      "version": 1,
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"index\":\"4fffdab0-8e01-11e9-957d-17ad3d0a703a\",\"query\":{\"query\":\"\",\"language\":\"lucene\"},\"filter\":[]}"
      }
    }
  }
]
```
