---
title: "Serverless Web Apps using Amazon DynamoDB - Part 3"
tags: aws, amazon-dynamodb, aws-lambda, amazon-api-gateway
url: https://amazon.qwiklabs.com/focuses/397?parent=catalog
---

# Goal
- Create an API for Lambda functions to access a DynamoDB table using Amazon API Gateway
- Generate the SDK for your API
- Configure and publish content with Amazon S3

# Task
- [x] Verify Your Region
- [x] Verify Resources
- [x] Create and Deploy an API
- [x] Generate the SDK For Your API

# Supplement
## CloudFormation
![](serverless_web_apps_using_amazon_dynamodb_part3.png)

```yaml
AWSTemplateFormatVersion: '2010-09-09'
Description: 'AWS CloudFormation Template for SuperMission: Builds DynamoDB tables,
  an S3 bucket, and Lambda functions for use in a real-time voting application. **
  This template creates multiple AWS resources. You will be billed for the AWS resources
  used if you create a stack from this template.'
Parameters:
  SuperMissionTableRead:
    Description: Read capacity units for DynamoDB table
    Type: String
    MinLength: '1'
    MaxLength: '6'
    AllowedPattern: '[0-9]*'
    Default: '1'
  SuperMissionTableWrite:
    Description: Write capacity units for DynamoDB table
    Type: String
    MinLength: '1'
    MaxLength: '6'
    AllowedPattern: '[0-9]*'
    Default: '1'
Resources:
  DynamoDBTable:
    Type: AWS::DynamoDB::Table
    Properties:
      AttributeDefinitions:
        - AttributeName: SuperHero
          AttributeType: S
        - AttributeName: MissionStatus
          AttributeType: S
        - AttributeName: Villain1
          AttributeType: S
        - AttributeName: Villain2
          AttributeType: S
        - AttributeName: Villain3
          AttributeType: S
        - AttributeName: SecretIdentity
          AttributeType: S
      KeySchema:
        - AttributeName: SuperHero
          KeyType: HASH
        - AttributeName: MissionStatus
          KeyType: RANGE
      ProvisionedThroughput:
        ReadCapacityUnits: !Ref 'SuperMissionTableRead'
        WriteCapacityUnits: !Ref 'SuperMissionTableWrite'
      StreamSpecification:
        StreamViewType: NEW_AND_OLD_IMAGES
      TableName: SuperMission
      GlobalSecondaryIndexes:
        - IndexName: myGSI
          KeySchema:
            - AttributeName: MissionStatus
              KeyType: HASH
            - AttributeName: Villain1
              KeyType: RANGE
          Projection:
            NonKeyAttributes:
              - SuperHero
              - SecretIdentity
            ProjectionType: INCLUDE
          ProvisionedThroughput:
            ReadCapacityUnits: '6'
            WriteCapacityUnits: '6'
        - IndexName: myGSI2
          KeySchema:
            - AttributeName: SecretIdentity
              KeyType: HASH
            - AttributeName: Villain3
              KeyType: RANGE
          Projection:
            NonKeyAttributes:
              - SuperHero
              - Villain1
            ProjectionType: INCLUDE
          ProvisionedThroughput:
            ReadCapacityUnits: '6'
            WriteCapacityUnits: '6'
      LocalSecondaryIndexes:
        - IndexName: myLSI
          KeySchema:
            - AttributeName: SuperHero
              KeyType: HASH
            - AttributeName: Villain2
              KeyType: RANGE
          Projection:
            NonKeyAttributes:
              - Villain1
              - MissionStatus
            ProjectionType: INCLUDE
  AppendItemToList:
    Type: AWS::Lambda::Function
    Properties:
      FunctionName: AppendItemToList
      Handler: lambdafunc.handler
      Role: !GetAtt 'SuperDynamoDBScanRole.Arn'
      Code:
        S3Bucket: us-east-1-aws-training
        S3Key: awsu-spl/spl133-dynamodb-webapp-part2/static/lambdafunc.zip
      Runtime: nodejs8.10
  getheroeslist:
    Type: AWS::Lambda::Function
    Properties:
      FunctionName: getheroeslist
      Handler: getheroeslistfunc.handler
      Role: !GetAtt 'SuperDynamoDBScanRole.Arn'
      Code:
        S3Bucket: us-east-1-aws-training
        S3Key: awsu-spl/spl134-dynamodb-webapp-part3/static/getheroeslistfunc.zip
      Runtime: nodejs8.10
  getmissiondetails:
    Type: AWS::Lambda::Function
    Properties:
      FunctionName: getmissiondetails
      Handler: getmissiondetailsfunc.handler
      Role: !GetAtt 'SuperDynamoDBQueryRole.Arn'
      Code:
        S3Bucket: us-east-1-aws-training
        S3Key: awsu-spl/spl134-dynamodb-webapp-part3/static/getmissiondetailsfunc.zip
      Runtime: nodejs8.10
  SuperDynamoDBScanRole:
    Type: AWS::IAM::Role
    Properties:
      RoleName: SuperDynamoDBScanRole
      AssumeRolePolicyDocument:
        Version: '2012-10-17'
        Statement:
          - Effect: Allow
            Principal:
              Service: lambda.amazonaws.com
            Action: sts:AssumeRole
  DynamoDBFullRolePolicies:
    Type: AWS::IAM::Policy
    Properties:
      PolicyName: SuperDynamoDBScanPolicy
      PolicyDocument:
        Version: '2012-10-17'
        Statement:
          - Effect: Allow
            Action:
              - dynamodb:Query
              - dynamodb:Scan
              - s3:GetObject
              - s3:PutObject
              - dynamodb:BatchWriteItem
              - dynamodb:*
            Resource:
              - '*'
      Roles:
        - !Ref 'SuperDynamoDBScanRole'
  SuperDynamoDBQueryRole:
    Type: AWS::IAM::Role
    Properties:
      RoleName: SuperDynamoDBQueryRole
      AssumeRolePolicyDocument:
        Version: '2012-10-17'
        Statement:
          - Effect: Allow
            Principal:
              Service: lambda.amazonaws.com
            Action: sts:AssumeRole
  DynamoDBSelectedRolePolicies:
    Type: AWS::IAM::Policy
    Properties:
      PolicyName: SuperDynamoDBQueryPolicy
      PolicyDocument:
        Version: '2012-10-17'
        Statement:
          - Effect: Allow
            Action:
              - dynamodb:Query
            Resource: '*'
      Roles:
        - !Ref 'SuperDynamoDBQueryRole'
  ScheduledRule:
    Type: AWS::Events::Rule
    Properties:
      Description: ScheduledRule
      ScheduleExpression: rate(1 minute)
      State: ENABLED
      Targets:
        - Arn: !GetAtt 'AppendItemToList.Arn'
          Id: TargetFunctionV1
  PermissionForEventsToInvokeLambda:
    Type: AWS::Lambda::Permission
    Properties:
      FunctionName: !Ref 'AppendItemToList'
      Action: lambda:InvokeFunction
      Principal: events.amazonaws.com
      SourceArn: !GetAtt 'ScheduledRule.Arn'
```

## S3
**Bucket Policy**
```json
{
    "Version": "2012-10-17",
    "Statement": [{
        "Sid": "PublicReadForGetBucketObjects",
        "Effect": "Allow",
        "Principal": "*",
        "Action": "s3:GetObject",
        "Resource": "arn:aws:s3:::mybucket-2019-09/*"
    }]
}
```
