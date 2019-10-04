import cdk = require('@aws-cdk/core')
import lambda = require('@aws-cdk/aws-lambda')
import dynamodb = require('@aws-cdk/aws-dynamodb')

export interface HitCounterProps {
  /**
   * the function for which we want to count url hits
   */
  downstream: lambda.IFunction
}

export class HitCounter extends cdk.Construct{
  /**
   * allows accesing the counter function
   */
  public readonly handler: lambda.Function

  constructor(scope: cdk.Construct, id: string, props: HitCounterProps){
    super(scope, id)

    const table = new dynamodb.Table(this, 'Hits', {
      partitionKey: {name: 'path', type: dynamodb.AttributeType.STRING}
    })

    this.handler = new lambda.Function(this, 'HitCounterHandler', {
      runtime: lambda.Runtime.NODEJS_8_10,
      handler: 'hitcounter.handler',
      code: lambda.Code.asset('lambda'),
      environment: {
        DOWNSTREAM_FUNCTION_NAME: props.downstream.functionName,
        HITS_TABLE_NAME: table.tableName
      }
    })
  }
}
