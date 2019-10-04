---
title: "TypeScript Workshop"
tags: amazon-web-services, aws-cdk, aws-cloudformation, provisioning, software-deployment
url: https://cdkworkshop.com/20-typescript.html
---

# Goal
- This version of the workshop will guide you through a getting started experience in TypeScript

# Task
- [x] Prerequistites
- [x] New Project
- [ ] Hello, CDK!
- [ ] Writing constructs
- [ ] Using construct libraries
- [ ] Clean up

# Supplement
## Prerequistites
```sh
$ npm i -g aws-cdk
$ cdk --version
1.11.0 (build 4ed4d96)

$ aws configure
```

## New Project
**Initialize project**
```sh
$ mkdir cdk-workshop && cd $_
$ cdk init sample-app --language typescript

# Useful commands

 * `npm run build`   compile typescript to js
 * `npm run watch`   watch for changes and compile
 * `npm run test`    perform the jest unit tests
 * `cdk deploy`      deploy this stack to your default AWS account/region
 * `cdk diff`        compare deployed stack with current state
 * `cdk synth`       emits the synthesized CloudFormation template
```

**Show CloudFormation configuration**
```sh
$ cdk synth
Resources:
  CdkWorkshopQueue50D9D426:
    Type: AWS::SQS::Queue
    Properties:
      VisibilityTimeout: 300
...
```

**Deploy bootstrap stack**
```sh
$ npm run watch
$ cdk bootstrap
```

**Deploy current stack**
```sh
$ cdk deploy
```

## Hello, CDK!
```sh
$ cat >>EOF >cdk-workshop-stack.ts
import cdk = require('@aws-cdk/core');

export class CdkWorkshopStack extends cdk.Stack {
  constructor(scope: cdk.App, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    // nothing here!
  }
}
EOF

$ cdk diff
...
Resources
[-] AWS::SQS::Queue CdkWorkshopQueue50D9D426 destroy
[-] AWS::SQS::QueuePolicy CdkWorkshopQueuePolicyAF2494A5 destroy
[-] AWS::SNS::Subscription CdkWorkshopQueueCdkWorkshopStackCdkWorkshopTopicD7BE96438B5AD106 destroy
[-] AWS::SNS::Topic CdkWorkshopTopicD368A42F destroy

$ cdk deploy

$ cdk diff
Stack CdkWorkshopStack
There were no differences
```

## Writing constructs
```sh
$ npm i @aws-cdk/aws-lambda

$ cat >>EOF >cdk-workshop-stack.ts
import cdk = require('@aws-cdk/core');
import lambda = require('@aws-cdk/aws-lambda');

export class CdkWorkshopStack extends cdk.Stack {
  constructor(scope: cdk.App, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    // defines an AWS Lambda resource
    const hello = new lambda.Function(this, 'HelloHandler', {
      runtime: lambda.Runtime.NODEJS_8_10, // execution environment
      code: lambda.Code.asset('lambda'), // code loaded from the "lambda" directory
      handler: 'hello.handler' // file is "hello", function is "handler"
    })
  }
}
EOF

$ cdk diff
...
Parameters
[+] Parameter HelloHandler/Code/S3Bucket HelloHandlerCodeS3Bucket4359A483: {"Type":"String","Description":"S3 bucket for asset \"CdkWorkshopStack/HelloHandler/Code\""}
[+] Parameter HelloHandler/Code/S3VersionKey HelloHandlerCodeS3VersionKey07D12610: {"Type":"String","Description":"S3 key for asset version \"CdkWorkshopStack/HelloHandler/Code\""}
[+] Parameter HelloHandler/Code/ArtifactHash HelloHandlerCodeArtifactHash5DF4E4B6: {"Type":"String","Description":"Artifact hash for asset \"CdkWorkshopStack/HelloHandler/Code\""}

Resources
[+] AWS::IAM::Role HelloHandler/ServiceRole HelloHandlerServiceRole11EF7C63
[+] AWS::Lambda::Function HelloHandler HelloHandler2E4FBA4D

$ cdk deploy
```

## Using construct libraries
## Clean up

## Reference
- http://bit.ly/cdkworkshopjp
- https://github.com/aws/aws-cdkworkshop
- https://github.com/aws/aws-cdk
- https://docs.aws.amazon.com/cdk/api/latest/
- https://docs.aws.amazon.com/cdk/latest/guide/home.html
- https://docs.aws.amazon.com/cdk/latest/guide/tools.html
- https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html
