import fs = require("fs");
import cdk = require("@aws-cdk/core");
import apigw = require("@aws-cdk/aws-apigateway");
import lambda = require("@aws-cdk/aws-lambda");
import iam = require("@aws-cdk/aws-iam");
import { AwsLambdaApigatewayStack } from "./apigateway-stack";

export class PetsLambdaConstruct extends cdk.Construct {
  public readonly functionName: string;
  private readonly uri: string;

  constructor(scope: AwsLambdaApigatewayStack, id: string) {
    super(scope, id);

    let uid: string = "";
    if (fs.existsSync("tmp/cache/uid-pets")) {
      uid = fs
        .readFileSync("tmp/cache/uid-pets")
        .toString()
        .replace(/^.+?:(.+?)\n$/, "$1");
    }

    const petsFn = new lambda.Function(scope, "petsLambda", {
      functionName: `${scope.stackName}-Pets`,
      runtime: lambda.Runtime.GO_1_X,
      code: lambda.Code.fromBucket(scope.lambdaBucket, `pets/_build/pets-${uid}.zip`),
      handler: `_build/pets-${uid}`,
      memorySize: 128,
      timeout: cdk.Duration.seconds(30),
      environment: {
        REGION: scope.region,
        ITEMS_TABLE_NAME: scope.table.items.tableName,
        USERS_TABLE_NAME: scope.table.users.tableName,
        ALLOWED_ORIGIN: scope.envAppUrl,
        ADMINS_GROUP_NAME: scope.envAdminsGroupName,
        USERS_GROUP_NAME: scope.envUsersGroupName,
        USER_POOL_ID: scope.userPool.defaultChildId,
        AUTHORIZATION_HEADER_NAME: scope.envAuthorizationHeaderName
      }
    });
    // grant apigateway access to lambda
    petsFn.addPermission(`${scope.stackName}-Pets`, scope.apiLambdaPermission);
    // grant lambda access to tables
    scope.table.items.grantReadWriteData(petsFn.role!);
    scope.table.users.grantReadWriteData(petsFn.role!);
    // grant lambda access to cognito
    petsFn.addToRolePolicy(
      new iam.PolicyStatement({
        resources: [scope.userPool.arn],
        actions: ["cognito-idp:AdminUserGlobalSignOut", "cognito-idp:AdminGetUser"]
      })
    );

    this.functionName = petsFn.functionName;
    this.uri = `arn:aws:apigateway:${scope.region}:lambda:path/2015-03-31/functions/${petsFn.functionArn}/invocations`;
    const pets = scope.api.root.addResource("pets");
    const pet = pets.addResource("{id}");

    // GET /pets
    // @note the path require the cognito authorizer (validates the JWT and passes it to the lambda)
    this.addMethod(pets, "GET", {
      authorizationType: apigw.AuthorizationType.COGNITO,
      authorizer: { authorizerId: scope.apiCognitoAuthorizer.id }
      // authorizationType: apigw.AuthorizationType.CUSTOM,
      // authorizer: { authorizerId: scope.apiLambdaAuthorizer.id }
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
    const apigwMethod = resource.addMethod(
      method,
      new apigw.Integration({
        type: apigw.IntegrationType.AWS_PROXY,
        uri: this.uri,
        integrationHttpMethod: "POST"
      }),
      methodOptions
    );

    // API Gateway now expects an access token instead of an ID token
    (apigwMethod.node.defaultChild as apigw.CfnMethod).authorizationScopes = ["openid"];
  }
}
