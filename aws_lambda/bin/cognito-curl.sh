#!/usr/bin/env bash

set -e

CURL_OPTIONS="$@"
USER_POOL_ID=$(
  aws cognito-idp list-user-pools \
    --max-results 20 |
    jq '.UserPools[] | select(.Name=="AwsLambdaApigatewayStack-UserPool") | .Id' |
    sed -E 's/"//g'
)
CLIENT_ID=$(
  aws cognito-idp list-user-pool-clients \
    --user-pool-id ${USER_POOL_ID} |
    jq '.UserPoolClients[] | select(.ClientName=="userPoolClient") | .ClientId' |
    sed -E 's/"//g'
)
USER_NAME="test-$(date +%s)@example.com"
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
  local session=$(getSession ${PASSWORD})
  aws cognito-idp admin-respond-to-auth-challenge \
    --user-pool-id ${USER_POOL_ID} \
    --client-id ${CLIENT_ID} \
    --challenge-name NEW_PASSWORD_REQUIRED \
    --challenge-responses USERNAME=${USER_NAME},NEW_PASSWORD=${NEW_PASSWORD} \
    --session ${session}
}

afterDeleteUser() {
  echo ""
  echo ""
  echo "EXECUTE: ${FUNCNAME[0]}"
  aws cognito-idp admin-delete-user \
    --user-pool-id ${USER_POOL_ID} \
    --username ${USER_NAME}
}

cognitoCurl() {
  echo "START: ${FUNCNAME[0]}"
  beforeCreateUser
  beforeChangePassword
  doCognitoCurl
  afterDeleteUser
  echo "END: ${FUNCNAME[0]}"
}

doCognitoCurl() {
  local session=$(getSession ${NEW_PASSWORD})
  eval "curl ${CURL_OPTIONS} -H 'Authorization: ${session}'"
}

getSession() {
  local pass=$1
  local rc=$(
    aws cognito-idp admin-initiate-auth \
      --user-pool-id ${USER_POOL_ID} \
      --client-id ${CLIENT_ID} \
      --auth-flow ADMIN_NO_SRP_AUTH \
      --auth-parameters USERNAME=${USER_NAME},PASSWORD=${pass} |
      jq -r .Session
  )
  echo $rc
}

cognitoCurl
