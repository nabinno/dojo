#!/usr/bin/env bash

set -e

buildGo() {
  local pkg=$1
  local old_last_modified_at=0
  local cdk_uid=""
  if [ -f "tmp/cache/uid-${pkg}" ]; then
    old_last_modified_at=$(cat "tmp/cache/uid-${pkg}" | sed -E 's/:.+$//g')
    cdk_uid=$(cat "tmp/cache/uid-${pkg}" | sed -E "s/^.+?:(.+?)$/\1/g")
  fi
  local new_last_modified_at=$(
    find "lambda/${pkg}" -maxdepth 1 -type f |
      xargs stat -c %Y |
      sort |
      tail -1
  )
  if [[ "${new_last_modified_at}" -le "${old_last_modified_at}" && -f "lambda/${pkg}/_build/${pkg}-${cdk_uid}" ]]; then
    return 0
  fi

  cdk_uid=$(nanoid --alphabet "1234567890abcdefghijklmnopqrstuvwxyz" --size 13)
  echo "${new_last_modified_at}:${cdk_uid}" >"tmp/cache/uid-${pkg}"

  (
    cd "lambda/${pkg}"
    rm -fr _build
    mkdir -p _build
    GOOS=linux GOARCH=amd64 go build -o "_build/${pkg}-${cdk_uid}"
    zip "_build/${pkg}-${cdk_uid}.zip" "_build/${pkg}-${cdk_uid}"
  )
}

if [ ! -f tmp/cache ]; then mkdir -p tmp/cache; fi

buildGo "pretokengen"
buildGo "authorizer"
buildGo "pets"
