import cdk = require("@aws-cdk/core");
import apigw = require("@aws-cdk/aws-apigateway");
import { AwsLambdaApigatewayStack } from "./apigateway-stack";

/**
 * @desc Purpose: create API endpoints and integrate with Amazon Cognito for JWT validation
 * @see https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-integrate-with-cognito.html
 */
export class CognitoAuthorizerConstruct extends cdk.Construct {
  public readonly id: string;

  constructor(scope: AwsLambdaApigatewayStack, id: string) {
    super(scope, id);

    const authorizer = new apigw.CfnAuthorizer(scope, "cognitoAuthorizer", {
      name: "cognitoAuthorizer",
      type: apigw.AuthorizationType.COGNITO,
      identitySource: `method.request.header.${scope.envAuthorizationHeaderName}`,
      restApiId: scope.api.restApiId,
      providerArns: [scope.userPool.arn]
    });

    this.id = authorizer.ref;
  }
}
