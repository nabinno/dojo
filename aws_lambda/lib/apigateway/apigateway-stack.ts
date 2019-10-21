import cdk = require("@aws-cdk/core");
import apigw = require("@aws-cdk/aws-apigateway");
import { PetsLambdaConstruct } from "./pets-lambda-construct";

export class AwsLambdaApigatewayStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    let stageName: string = "prod";
    if (this.node.tryGetContext("stage") !== undefined) {
      stageName = this.node.tryGetContext("stage");
    }
    const stackName = this.stackName;
    const account = this.account;
    const region = this.region;

    const api = new apigw.RestApi(this, "gateway", { restApiName: stackName });
    const requestValidator = new apigw.RequestValidator(this, "requestvalidator", {
      restApi: api,
      requestValidatorName: `validateRequestBodyAndParameters`,
      validateRequestBody: true,
      validateRequestParameters: false
    });

    new PetsLambdaConstruct(this, "petsLambdaConstruct", {
      stageName: stageName,
      stackName: stackName,
      account: account,
      region: region,
      api: api,
      requestValidator: requestValidator
    });
  }
}
