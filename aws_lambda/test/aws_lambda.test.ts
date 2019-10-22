import { expect as expectCDK, matchTemplate, MatchStyle } from "@aws-cdk/assert";
import cdk = require("@aws-cdk/core");
import { AwsLambdaApigatewayStack } from "../lib/apigateway/apigateway-stack";

describe("AwsLambdaApigatewayStack", () => {
  const app = new cdk.App();
  const stack = new AwsLambdaApigatewayStack(app, "MyTestStack");

  it("should match template", () => {
    expectCDK(stack).to(
      matchTemplate(
        {
          Resources: {}
        },
        MatchStyle.EXACT
      )
    );
  });
});
