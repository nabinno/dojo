import cdk = require("@aws-cdk/core");
import lambda = require("@aws-cdk/aws-lambda");
import apigw = require("@aws-cdk/aws-apigateway");
import s3 = require("@aws-cdk/aws-s3");
import iam = require("@aws-cdk/aws-iam");

export class AwsLambdaApigatewayStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    let stageName: string = "Prod";
    if (this.node.tryGetContext("stage") !== undefined) {
      stageName = this.node.tryGetContext("stage");
    }
    const stackName = this.stackName;
    const account = this.account;
    const region = this.region;

    const api = new apigw.RestApi(this, "gateway", { restApiName: stackName });
    const petsFn = new lambda.Function(this, "lambda", {
      functionName: `${stackName}-Pets`,
      runtime: lambda.Runtime.GO_1_X,
      code: lambda.Code.fromBucket(
        s3.Bucket.fromBucketName(
          this,
          "s3bucket",
          `${stageName.toLowerCase()}-awslambdas3deploystack`
        ),
        "pets/pets.zip"
      ),
      handler: "pets",
      memorySize: 256,
      timeout: cdk.Duration.seconds(300),
      environment: {}
    });
    petsFn.addPermission(`${stackName}-Pets`, {
      principal: new iam.ServicePrincipal("apigateway.amazonaws.com"),
      sourceArn: `arn:aws:execute-api:${region}:${account}:${api.restApiId}/*`
    });
    const petsFnUri = `arn:aws:apigateway:${region}:lambda:path/2015-03-31/functions/${petsFn.functionArn}/invocations`;
    const pets = api.root.addResource("pets");
    const pet = pets.addResource("{id}");

    // GET /pets
    pets.addMethod(
      "GET",
      new apigw.Integration({
        type: apigw.IntegrationType.AWS_PROXY,
        uri: petsFnUri,
        integrationHttpMethod: "POST",
        options: {}
      }),
      { authorizationType: apigw.AuthorizationType.NONE }
    );

    // GET /pets/:id
    pet.addMethod(
      "GET",
      new apigw.Integration({
        type: apigw.IntegrationType.AWS_PROXY,
        uri: petsFnUri,
        integrationHttpMethod: "POST",
        options: {}
      }),
      {
        authorizationType: apigw.AuthorizationType.NONE,
        requestParameters: { "method.request.path.id": true }
      }
    );

    // POST /pets
    pets.addMethod(
      "POST",
      new apigw.Integration({
        type: apigw.IntegrationType.AWS_PROXY,
        uri: petsFnUri,
        integrationHttpMethod: "POST",
        options: {}
      }),
      {
        authorizationType: apigw.AuthorizationType.NONE,
        requestParameters: {
          "method.request.body.breed": true,
          "method.request.body.name": true,
          "method.request.body.dateOfBirth": true
        }
      }
    );
  }
}
