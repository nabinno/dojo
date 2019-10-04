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
**Cleanup sample**
```sh
$ cat >>EOF >lib/cdk-workshop-stack.ts
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

**Hello Lambda**
```sh
$ npm i @aws-cdk/aws-lambda

$ cat >>EOF >lib/cdk-workshop-stack.ts
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

**API Gateway**
```sh
$ npm i @aws-cdk/aws-apigateway

$ cat <<EOF >lib/cdk-workshop-stack.ts
import cdk = require('@aws-cdk/core')
import lambda = require('@aws-cdk/aws-lambda')
import apigw = require('@aws-cdk/aws-apigateway')

export class CdkWorkshopStack extends cdk.Stack {
  constructor(scope: cdk.App, id: string, props?: cdk.StackProps) {
    super(scope, id, props)

    // defines an AWS Lambda resource
    const hello = new lambda.Function(this, 'HelloHandler', {
      runtime: lambda.Runtime.NODEJS_8_10, // execution environment
      code: lambda.Code.asset('lambda'), // code loaded from the "lambda" directory
      handler: 'hello.handler' // file is "hello", function is "handler"
    })

    // defines an API Gateway REST API resource backed by our "hello" function.
    new apigw.LambdaRestApi(this, 'Endpoint', {
      handler: hello
    })
  }
}
EOF

$ cdk diff
...
Resources
[+] AWS::ApiGateway::RestApi Endpoint EndpointEEF1FD8F
[+] AWS::ApiGateway::Deployment Endpoint/Deployment EndpointDeployment318525DAa881978a07aac8a834ebf84a9e1ae8b7
[+] AWS::ApiGateway::Stage Endpoint/DeploymentStage.prod EndpointDeploymentStageprodB78BEEA0
[+] AWS::IAM::Role Endpoint/CloudWatchRole EndpointCloudWatchRoleC3C64E0F
[+] AWS::ApiGateway::Account Endpoint/Account EndpointAccountB8304247
[+] AWS::ApiGateway::Resource Endpoint/Default/{proxy+} Endpointproxy39E2174E
[+] AWS::Lambda::Permission Endpoint/Default/{proxy+}/ANY/ApiPermission.CdkWorkshopStackEndpoint018E8349.ANY..{proxy+} EndpointproxyANYApiPermissionCdkWorkshopStackEndpoint018E8349ANYproxy747DCA52
[+] AWS::Lambda::Permission Endpoint/Default/{proxy+}/ANY/ApiPermission.Test.CdkWorkshopStackEndpoint018E8349.ANY..{proxy+} EndpointproxyANYApiPermissionTestCdkWorkshopStackEndpoint018E8349ANYproxy41939001
[+] AWS::ApiGateway::Method Endpoint/Default/{proxy+}/ANY EndpointproxyANYC09721C5
[+] AWS::Lambda::Permission Endpoint/Default/ANY/ApiPermission.CdkWorkshopStackEndpoint018E8349.ANY.. EndpointANYApiPermissionCdkWorkshopStackEndpoint018E8349ANYE84BEB04
[+] AWS::Lambda::Permission Endpoint/Default/ANY/ApiPermission.Test.CdkWorkshopStackEndpoint018E8349.ANY.. EndpointANYApiPermissionTestCdkWorkshopStackEndpoint018E8349ANYB6CC1B64
[+] AWS::ApiGateway::Method Endpoint/Default/ANY EndpointANY485C938B

Outputs
[+] Output Endpoint/Endpoint Endpoint8024A810: {"Value":{"Fn::Join":["",["https://",{"Ref":"EndpointEEF1FD8F"},".execute-api.",{"Ref": "AWS::Region"},".",{"Ref":"AWS::URLSuffix"},"/",{"Ref":"EndpointDeploymentStageprodB78BEEA0"},"/"]]}}

$ cdk deploy
...
Outputs:
CdkWorkshopStack.Endpoint8024A810 = https://d9mq4rw3cd.execute-api.ap-northeast-1.amazonaws.com/prod/
```


## Writing constructs

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
