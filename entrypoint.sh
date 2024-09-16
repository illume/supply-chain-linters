#!/usr/bin/env bash

# array of strings for github actions
ARGS=("$@")

if [[ ! -z "${GITHUB_AUTHENTICATION_TOKEN}" ]]; then
  git config --global --add url."https://x-access-token:${GITHUB_AUTHENTICATION_TOKEN}@github.com/".insteadOf "https://github.com/"
fi

/bin/supply-chain-linters ${ARGS[*]}
