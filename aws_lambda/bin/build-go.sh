#!/usr/bin/env bash

buildGo() {
  local pkg=$1
  (
    cd lambda/${pkg}
    rm -fr _build
    mkdir -p _build
    GOOS=linux GOARCH=amd64 go build -o _build/${pkg}-${CDK_UID}
    zip _build/${pkg}-${CDK_UID}.zip _build/${pkg}-${CDK_UID}
  )
}

CDK_UID=$(nanoid --alphabet "1234567890abcdefghijklmnopqrstuvwxyz" --size 13)

mkdir -p tmp/cache && echo ${CDK_UID} >tmp/cache/uid

buildGo "authorizer"
buildGo "pets"
