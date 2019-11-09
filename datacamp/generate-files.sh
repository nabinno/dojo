#!/usr/bin/env bash

set -e

generateContent() {
  local tags="$1"
  local title="$2"
  cat <<EOF
---
title: ${title}
tags: ${tags}
url:
---
EOF
}

convertToSnakeCase() {
  local string="$1"
  echo "${string}" |
    tr '[:upper:]' '[:lower:]' |
    sed -r 's/(\(|\))//g' |
    sed -r 's/( )([a-z0-9])/_\2/g'
}

while read TAGS TITLE; do
  generateContent "${TAGS}" "${TITLE}" >"$(convertToSnakeCase "${TITLE}").md"
done <title.tsv
