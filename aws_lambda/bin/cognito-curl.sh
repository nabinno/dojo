#!/usr/bin/env bash

set -e

CURL_OPTIONS="$@"
USER_POOL_ID=$(
  aws cognito-idp list-user-pools \
    --max-results 20 |
    jq -r '.UserPools[] | select(.Name=="AwsLambdaApigatewayStack-UserPool") | .Id'
)
CLIENT_ID=$(
  aws cognito-idp list-user-pool-clients \
    --user-pool-id ${USER_POOL_ID} |
    jq -r '.UserPoolClients[] | select(.ClientName=="userPoolClient") | .ClientId'
)
ISO8601_TIME=$(date --iso-8601=seconds)
UNIX_TIME=$(date -d ${ISO8601_TIME} +%s)
USER_NAME="test-${UNIX_TIME}@example.com"
PASSWORD="P@ssw0rd"
NEW_PASSWORD="P@ssw0rD"

beforeCreateUser() {
  echo "EXECUTE: ${FUNCNAME[0]}"
  aws cognito-idp admin-create-user \
    --user-pool-id ${USER_POOL_ID} \
    --username ${USER_NAME} \
    --user-attributes Name="email",Value=${USER_NAME} \
    --temporary-password ${PASSWORD}
}

beforeChangePassword() {
  echo "EXECUTE: ${FUNCNAME[0]}"
  local session=$(initiateAuth ${PASSWORD} | jq -r .Session)

  aws cognito-idp admin-respond-to-auth-challenge \
    --user-pool-id ${USER_POOL_ID} \
    --client-id ${CLIENT_ID} \
    --challenge-name NEW_PASSWORD_REQUIRED \
    --challenge-responses USERNAME=${USER_NAME},NEW_PASSWORD=${NEW_PASSWORD} \
    --session ${session}
}

beforePutItem() {
  echo "EXECUTE: ${FUNCNAME[0]}"
  local username=$(getUsername)
  local item=$(
    cat <<JSON
{
  "id": {
    "S": "${UNIX_TIME}"
  },
  "breed": {
    "S": "Breed-${UNIX_TIME}"
  },
  "name": {
    "S": "Name-${UNIX_TIME}"
  },
  "dateOfBirth": {
    "S": "${ISO8601_TIME}"
  },
  "owner": {
    "S": "${username}"
  },
  "ownerDisplayName": {
    "S": "${username}"
  }
}
JSON
  )

  aws dynamodb put-item \
    --table-name AwsLambdaApigatewayStack-Items \
    --item "${item}"
}

afterDeleteUser() {
  echo "EXECUTE: ${FUNCNAME[0]}"
  aws cognito-idp admin-delete-user \
    --user-pool-id ${USER_POOL_ID} \
    --username ${USER_NAME}
}

cognitoCurl() {
  echo "START: ${FUNCNAME[0]}"
  beforeCreateUser
  beforeChangePassword
  beforePutItem
  doCognitoCurl
  # afterDeleteUser
  echo "END: ${FUNCNAME[0]}"
}

# @todo 2019-11-03
doCognitoCurl() {
  local session=$(initiateAuth ${NEW_PASSWORD} | jq -r '.AuthenticationResult.IdToken')
  eval "curl -H 'Authorization:Bearer ${session}' ${CURL_OPTIONS}"
  echo ""
  echo ""
}

initiateAuth() {
  local pass=$1
  local rc=$(
    aws cognito-idp admin-initiate-auth \
      --user-pool-id ${USER_POOL_ID} \
      --client-id ${CLIENT_ID} \
      --auth-flow ADMIN_NO_SRP_AUTH \
      --auth-parameters USERNAME=${USER_NAME},PASSWORD=${pass}
  )
  echo $rc
}

getUsername() {
  local rc=$(
    aws cognito-idp admin-get-user \
      --user-pool-id ${USER_POOL_ID} \
      --username ${USER_NAME} |
      jq -r '.Username'
  )
  echo $rc
}

cognitoCurl
