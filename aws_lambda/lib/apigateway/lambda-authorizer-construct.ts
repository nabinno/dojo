import cdk = require("@aws-cdk/core");
import apigw = require("@aws-cdk/aws-apigateway");
import lambda = require("@aws-cdk/aws-lambda");
import { AwsLambdaApigatewayStack } from "./apigateway-stack";

export class LambdaAuthorizerConstruct extends cdk.Construct {
  public readonly id: string;

  constructor(scope: AwsLambdaApigatewayStack, id: string) {
    super(scope, id);

    const uid = process.env.CDK_UID;

    const fn = new lambda.Function(scope, "authorizerLambda", {
      functionName: `${scope.stackName}-Authorizer`,
      runtime: lambda.Runtime.GO_1_X,
      code: lambda.Code.fromBucket(scope.lambdaBucket, "authorizer/_build/authorizer.zip"),
      handler: `_build/authorizer-${uid}`,
      memorySize: 256,
      timeout: cdk.Duration.seconds(300),
      environment: {}
    });
    fn.addPermission(`${scope.stackName}-Authorizer`, scope.lambdaPermission);

    const authorizer = new apigw.CfnAuthorizer(scope, "lambdaAuthorizer", {
      name: "lambdaAuthorizer",
      restApiId: scope.api.restApiId,
      type: "REQUEST",
      authorizerUri: `arn:aws:apigateway:${scope.region}:lambda:path/2015-03-31/functions/${fn.functionArn}/invocations`,
      authorizerResultTtlInSeconds: 300,
      identitySource: "method.request.header.authorizationToken"
    });

    this.id = authorizer.ref;
  }
}
