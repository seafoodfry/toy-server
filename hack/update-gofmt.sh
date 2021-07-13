#!/usr/bin/env bash

# Adapted from
# https://github.com/kubernetes/kubernetes/blob/master/hack/update-gofmt.sh.

set -o errexit
set -o nounset
set -o pipefail


find_files() {
  find . -not \( \
      \( \
        -wholename './.git' \
      \) -prune \
    \) -name '*.go'
}

find_files | xargs gofmt -s -w
