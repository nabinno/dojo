import cdk = require("@aws-cdk/core");
import apigw = require("@aws-cdk/aws-apigateway");
import lambda = require("@aws-cdk/aws-lambda");
import { AwsLambdaApigatewayStack } from "./apigateway-stack";

export class PetsLambdaConstruct extends cdk.Construct {
  private readonly uri: string;

  constructor(scope: AwsLambdaApigatewayStack, id: string) {
    super(scope, id);

    const uid = process.env.CDK_UID;

    const petsFn = new lambda.Function(scope, "petsLambda", {
      functionName: `${scope.stackName}-Pets`,
      runtime: lambda.Runtime.GO_1_X,
      code: lambda.Code.fromBucket(scope.lambdaBucket, "pets/_build/pets.zip"),
      handler: `_build/pets-${uid}`,
      memorySize: 256,
      timeout: cdk.Duration.seconds(300),
      environment: {}
    });
    petsFn.addPermission(`${scope.stackName}-Pets`, scope.lambdaPermission);

    this.uri = `arn:aws:apigateway:${scope.region}:lambda:path/2015-03-31/functions/${petsFn.functionArn}/invocations`;
    const pets = scope.api.root.addResource("pets");
    const pet = pets.addResource("{id}");

    // GET /pets
    this.addMethod(pets, "GET", {
      authorizationType: apigw.AuthorizationType.CUSTOM,
      authorizer: { authorizerId: scope.apiLambdaAuthorizer.id }
    });

    // GET /pets/:id
    this.addMethod(pet, "GET", { authorizationType: apigw.AuthorizationType.NONE });

    // POST /pets
    this.addMethod(pets, "POST", {
      authorizationType: apigw.AuthorizationType.CUSTOM,
      authorizer: { authorizerId: scope.apiLambdaAuthorizer.id },
      requestValidator: scope.apiRequestValidator,
      requestModels: {
        "application/json": new apigw.Model(scope, "petModel", {
          restApi: scope.api,
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

  private addMethod(
    resource: apigw.Resource,
    method: string,
    methodOptions?: apigw.MethodOptions
  ) {
    resource.addMethod(
      method,
      new apigw.Integration({
        type: apigw.IntegrationType.AWS_PROXY,
        uri: this.uri,
        integrationHttpMethod: "POST"
      }),
      methodOptions
    );
  }
}
