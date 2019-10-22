import cdk = require("@aws-cdk/core");
import apigw = require("@aws-cdk/aws-apigateway");
import lambda = require("@aws-cdk/aws-lambda");
import s3 = require("@aws-cdk/aws-s3");
import iam = require("@aws-cdk/aws-iam");
import { LambdaAuthorizerConstruct } from "./lambda-authorizer-construct";
import { PetsLambdaConstruct } from "./pets-lambda-construct";

export class AwsLambdaApigatewayStack extends cdk.Stack {
  public readonly stageName: string;
  public readonly api: apigw.RestApi;
  public readonly apiRequestValidator: apigw.RequestValidator;
  public readonly apiLambdaAuthorizer: LambdaAuthorizerConstruct;
  public readonly lambdaBucket: s3.IBucket;
  public readonly lambdaPermission: lambda.Permission;

  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    this.stageName = "prod";
    if (this.node.tryGetContext("stage") !== undefined) {
      this.stageName = this.node.tryGetContext("stage");
    }

    this.api = new apigw.RestApi(this, "gateway", { restApiName: this.stackName });
    this.apiRequestValidator = new apigw.RequestValidator(this, "requestvalidator", {
      restApi: this.api,
      requestValidatorName: "validateRequestBodyAndParameters",
      validateRequestBody: true,
      validateRequestParameters: false
    });
    this.lambdaPermission = {
      principal: new iam.ServicePrincipal("apigateway.amazonaws.com"),
      sourceArn: `arn:aws:execute-api:${this.region}:${this.account}:${this.api.restApiId}/*`
    };
    this.lambdaBucket = s3.Bucket.fromBucketName(
      this,
      "s3bucket",
      `${this.stageName}-awslambdas3deploystack`
    );
    // this.apiLambdaAuthorizer = new LambdaAuthorizerConstruct(this, "lambdaAuthorizerConstruct");

    new PetsLambdaConstruct(this, "petsLambdaConstruct");
  }
}
