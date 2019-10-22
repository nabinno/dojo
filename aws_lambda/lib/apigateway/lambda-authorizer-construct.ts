import cdk = require("@aws-cdk/core");
import apigw = require("@aws-cdk/aws-apigateway");
import lambda = require("@aws-cdk/aws-lambda");
import { AwsLambdaApigatewayStack } from "./apigateway-stack";

export class LambdaAuthorizerConstruct extends cdk.Construct {
  public readonly id: string;

  constructor(scope: AwsLambdaApigatewayStack, id: string) {
    super(scope, id);

    const fn = new lambda.Function(scope, "authorizerLambda", {
      functionName: `${scope.stackName}-Authorizer`,
      runtime: lambda.Runtime.GO_1_X,
      code: lambda.Code.fromBucket(scope.lambdaBucket, "authorizer/authorizer.zip"),
      handler: "authorizer",
      memorySize: 256,
      timeout: cdk.Duration.seconds(300),
      environment: {}
    });

    const authorizer = new apigw.CfnAuthorizerV2(scope, "authorizer", {
      apiId: scope.api.restApiId,
      name: "lambdaAuthorizer",
      authorizerType: "REQUEST",
      authorizerUri: `arn:aws:apigateway:${scope.region}:lambda:path/2015-03-31/functions/${fn.functionArn}/invocations`,
      identitySource: ["method.request.header.Auth"]
    });

    this.id = authorizer.logicalId;
  }
}
