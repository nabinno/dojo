import cdk = require("@aws-cdk/core");
import apigw = require("@aws-cdk/aws-apigateway");
import lambda = require("@aws-cdk/aws-lambda");
import s3 = require("@aws-cdk/aws-s3");
import iam = require("@aws-cdk/aws-iam");

export interface PetsLambdaProps {
  stageName: string;
  stackName: string;
  account: string;
  region: string;
  api: apigw.RestApi;
  requestValidator: apigw.RequestValidator;
}

export class PetsLambdaConstruct extends cdk.Construct {
  private readonly uri: string;

  constructor(scope: cdk.Construct, id: string, props: PetsLambdaProps) {
    super(scope, id);

    const stageName = props.stageName;
    const stackName = props.stackName;
    const account = props.account;
    const region = props.region;
    const api = props.api;
    const requestValidator = props.requestValidator;

    const petsFn = new lambda.Function(scope, "petsLambda", {
      functionName: `${stackName}-Pets`,
      runtime: lambda.Runtime.GO_1_X,
      code: lambda.Code.fromBucket(
        s3.Bucket.fromBucketName(scope, "s3bucket", `${stageName}-awslambdas3deploystack`),
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

    this.uri = `arn:aws:apigateway:${region}:lambda:path/2015-03-31/functions/${petsFn.functionArn}/invocations`;
    const pets = api.root.addResource("pets");
    const pet = pets.addResource("{id}");

    // GET /pets
    this.addMethod(pets, "GET", { authorizationType: apigw.AuthorizationType.NONE });

    // GET /pets/:id
    this.addMethod(pet, "GET", { authorizationType: apigw.AuthorizationType.NONE });

    // POST /pets
    this.addMethod(pets, "POST", {
      // authorizationType: apigw.AuthorizationType.CUSTOM,
      // authorizer: { authorizerId: "1" },
      authorizationType: apigw.AuthorizationType.NONE,
      requestValidator: requestValidator,
      requestModels: {
        "application/json": new apigw.Model(scope, "petModel", {
          restApi: api,
          modelName: "petModel",
          contentType: "application/json",
          schema: {
            type: apigw.JsonSchemaType.OBJECT,
            properties: {
              breed: { type: apigw.JsonSchemaType.STRING },
              name: { type: apigw.JsonSchemaType.STRING },
              dateOfBirth: { type: apigw.JsonSchemaType.STRING }
            },
            required: ["breed", "name", "dateOfBirth"]
          }
        })
      }
    });
  }

  private addMethod(resource: apigw.Resource, method: string, options?: apigw.MethodOptions) {
    resource.addMethod(
      method,
      new apigw.Integration({
        type: apigw.IntegrationType.AWS_PROXY,
        uri: this.uri,
        integrationHttpMethod: "POST"
      }),
      options
    );
  }
}
