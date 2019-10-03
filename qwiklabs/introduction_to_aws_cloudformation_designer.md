---
title: "Introduction to AWS CloudFormation Designer"
tags: amazon-web-services, aws-cloudformation, aws-cloudformation-designer, provisioning, software-deployment
url: https://www.qwiklabs.com/focuses/384?parent=catalog
---

# Goal
- Create a simple single-instance stack using AWS CloudFormation Designer
- Add and edit Mappings, Resource Parameters, and Outputs
- Validate the stack and save your basic template
- Launch the stack and view the running stack and its Outputs
- Delete and clean up when the stack is no longer required

# Task
- [x] Task 1: Add an Instance Resource
- [x] Task 2: Set Parameters
- [x] Task 3: Add Mappings
- [x] Task 4:  Add Outputs
- [x] Task 5: Specify Resource Properties
- [x] Task 6: Create the Stack
- [x] Task 7: View the Running Stack
- [x] Task 8: Delete the Stack

# Supplement
## Task 1: Add an Instance Resource
```yaml
Resources:
  ServerInstance:
    Type: 'AWS::EC2::Instance'
    Properties: {}
```

## Task 2: Set Parameters
```yaml
Parameters:
  InstanceType:
    Description: Server EC2 instance type
    Type: String
    Default: t2.micro
    AllowedValues:
      - t1.micro
      - t2.micro
      - t2.small
      - t2.medium
    ConstraintDescription: must be a valid EC2 instance type.
  KeyName:
    Description: Name of an EC2 KeyPair to enable SSH access to the instance.
    Type: 'AWS::EC2::KeyPair::KeyName'
    ConstraintDescription: must be the name of an existing EC2 KeyPair.
```

## Task 3: Add Mappings
```yaml
Mappings:
  AWSInstanceType2Arch:
    t1.micro:
      Arch: PV64
    t2.micro:
      Arch: HVM64
    t2.small:
      Arch: HVM64
    t2.medium:
      Arch: HVM64
  AWSRegionArch2AMI:
    us-east-1:
      PV64: ami-1ccae774
      HVM64: ami-1ecae776
      HVMG2: ami-8c6b40e4
    us-west-2:
      PV64: ami-ff527ecf
      HVM64: ami-e7527ed7
      HVMG2: ami-abbe919b
    us-west-1:
      PV64: ami-d514f291
      HVM64: ami-d114f295
      HVMG2: ami-f31ffeb7
    eu-west-1:
      PV64: ami-bf0897c8
      HVM64: ami-a10897d6
      HVMG2: ami-d5bc24a2
    eu-central-1:
      PV64: ami-ac221fb1
      HVM64: ami-a8221fb5
      HVMG2: ami-7cd2ef61
    ap-northeast-1:
      PV64: ami-27f90e27
      HVM64: ami-cbf90ecb
      HVMG2: ami-6318e863
    ap-southeast-1:
      PV64: ami-acd9e8fe
      HVM64: ami-68d8e93a
      HVMG2: ami-3807376a
    ap-southeast-2:
      PV64: ami-ff9cecc5
      HVM64: ami-fd9cecc7
      HVMG2: ami-89790ab3
    sa-east-1:
      PV64: ami-bb2890a6
      HVM64: ami-b52890a8
      HVMG2: NOT_SUPPORTED
    cn-north-1:
      PV64: ami-fa39abc3
      HVM64: ami-f239abcb
      HVMG2: NOT_SUPPORTED
```

## Task 4:  Add Outputs
```yaml
Outputs:
  Hostname:
    Value: !Join 
      - ''
      - - 'Hostname:'
        - !GetAtt 
          - ServerInstance
          - PublicIp
    Description: Newly created server hostname
```

## Task 5: Specify Resource Properties
```yaml
Resources:
  ServerInstance:
    Type: 'AWS::EC2::Instance'
    Properties:
      InstanceType: !Ref InstanceType
      ImageId: !FindInMap 
        - AWSRegionArch2AMI
        - !Ref 'AWS::Region'
        - !FindInMap 
          - AWSInstanceType2Arch
          - !Ref InstanceType
          - Arch
      KeyName: !Ref KeyName
```
