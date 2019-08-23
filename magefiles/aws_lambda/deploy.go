package awslambda

import (
	"fmt"
	"os"
)

const envName = "dev"
const awsDefaultRegion = "ap-northeast-1"

var stackName = fmt.Sprintf("lambda-stack-%v", envName)
var bucketName = fmt.Sprintf("resources.%v.host", envName)

// Deploy ....
func Deploy() {
	if out, err := doDeploy(); err == nil {
		fmt.Println("deployment completed.")
		fmt.Println(string(out))
	} else {
		fmt.Println("deployment failure.")
		fmt.Println(err)
	}
}

func doDeploy() (out []byte, err error) {
	if err := os.Chdir("aws_lambda"); err != nil {
		if out, err = createDeployYaml(); err != nil {
			return out, err
		}

		if err = os.Mkdir("build", os.ModeDir); err != nil {
			return out, err
		}

		if out, err = cfPackage(); err != nil {
			return out, err
		}

		if out, err = cfDeploy(); err != nil {
			return out, err
		}

		switch envName {
		case "dev":
			out, err = deleteOldRevisionsFromS3()
		}

		err = os.RemoveAll("build")
	}
	return
}

func createDeployYaml() ([]byte, error) {
	return outcmd(`
cat <<EOF >template.yml
AWSTemplateFormatVersion: '2010-09-09'
Transform: AWS::Serverless-2016-10-31
Description: Exercise Lambda Gin

Resources:
  SampleFunction:
	Type: AWS::Serverless::Function
	Properties:
	  Handler: main
	  CodeUri: main.zip
	  Runtime: go1.x
	  MemorySize: 128
	  Policies: AWSLambdaBasicExecutionRole
	  Timeout: 3
	  Events:
		GetResource:
		  Type: Api
		  Properties:
			Path: /{proxy+}
			Method: any

Outputs:
  SampleGinApi:
	Description: URL for application
	Value: !Sub 'https://${ServerlessRestApi}.execute-api.${AWS::Region}.amazonaws.com/Prod/pets'
	Export:
	  Name: SampleGinApi
EOF
`)
}

// @todo 2019-08-24
func cfPackage() ([]byte, error) {
	return outcmd(fmt.Sprintf(`
aws cloudformation package \
  --profile foo \
  --template-file ./template.yml \
  --s3-bucket %v \
  --s3-prefix foo \
  --output-template-file ./build/packaged-template.yml
`, bucketName))
}

func cfDeploy() ([]byte, error) {
	return outcmd(fmt.Sprintf(`
aws cloudformation deploy \
  --profile foo \
  --template-file ./build/packaged-template.yml \
  --stack-name "%v" \
  --role-arn arn:aws:iam::767815591449:role/cloudformation-role \
  --parameter-overrides EnvName=%v \
  --tags "APP-NAME=foo" "ROLE=lambda" "ENVIRONMENT=%v"
`, stackName, envName, envName))
}

func deleteOldRevisionsFromS3() ([]byte, error) {
	return outcmd(`
aws s3 ls s3://resources.host/ --profile foo |
	sort -nr |
	tail -n +6 |
	awk '{print $4}' |
	xargs -I{} -t aws s3 rm s3://resources.host/lambda/{} --profile foo
`)
}
