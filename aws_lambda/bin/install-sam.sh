#!/usr/bin/env bash

get-brew() {
  case "${OSTYPE}" in
    darwin*) ;;
    linux*)
      sh -c "$(curl -fsSL https://raw.githubusercontent.com/Linuxbrew/install/master/install.sh)"
      brew vendor-install ruby
      ;;
  esac
}
if ! type -p brew >/dev/null; then get-brew; fi

get-aws-sam() {
  case "${OSTYPE}" in
    darwin* | linux*)
      brew tap aws/tap
      brew install aws-sam-cli
      ;;
  esac
}
if ! type -p sam >/dev/null; then get-aws-sam; fi
