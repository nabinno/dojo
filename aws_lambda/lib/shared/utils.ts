import {
  CloudFormationClient,
  DescribeStacksCommand,
  Output,
} from "@aws-sdk/client-cloudformation";
import apigateway = require("@aws-cdk/aws-apigateway");
import { URL } from "url";

export class Utils {
  static async getStackOutputs(
    stackName: string,
    stackRegion: string
  ): Promise<Output[]> {
    const client = new CloudFormationClient({ region: stackRegion });
    const result = await client.send(
      new DescribeStacksCommand({ StackName: stackName })
    );
    return result.Stacks?.[0]?.Outputs ?? [];
  }

  static getEnv(variableName: string, defaultValue?: string) {
    const variable = process.env[variableName];
    if (!variable) {
      if (defaultValue !== undefined) {
        return defaultValue;
      }
      throw new Error(`${variableName} environment variable must be defined`);
    }
    return variable;
  }

  static addCorsOptions(
    apiResource: apigateway.IResource,
    origin: string,
    allowCredentials: boolean = false,
    allowMethods: string = "OPTIONS,GET,PUT,POST,DELETE"
  ) {
    apiResource.addMethod(
      "OPTIONS",
      new apigateway.MockIntegration({
        integrationResponses: [
          {
            statusCode: "200",
            responseParameters: {
              "method.response.header.Access-Control-Allow-Headers":
                "'Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token,X-Amz-User-Agent'",
              "method.response.header.Access-Control-Allow-Origin": "'" + origin + "'",
              "method.response.header.Access-Control-Allow-Credentials":
                "'" + allowCredentials.toString() + "'",
              "method.response.header.Access-Control-Allow-Methods": "'" + allowMethods + "'",
              "method.response.header.Access-Control-Max-Age": "'7200'"
            }
          }
        ],
        passthroughBehavior: apigateway.PassthroughBehavior.NEVER,
        requestTemplates: {
          "application/json": '{"statusCode": 200}'
        }
      }),
      {
        methodResponses: [
          {
            statusCode: "200",
            responseParameters: {
              "method.response.header.Access-Control-Allow-Headers": true,
              "method.response.header.Access-Control-Allow-Methods": true,
              "method.response.header.Access-Control-Allow-Credentials": true,
              "method.response.header.Access-Control-Allow-Origin": true,
              "method.response.header.Access-Control-Max-Age": true
            }
          }
        ]
      }
    );
  }

  static isURL(identityProviderMetadataURLOrFile: string) {
    try {
      new URL(identityProviderMetadataURLOrFile);
      return true;
    } catch {
      return false;
    }
  }
}
