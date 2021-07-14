#!/usr/bin/env bash

# Adapted from
# https://github.com/kubernetes/kubernetes/blob/master/hack/verify-golangci-lint.sh.

set -o errexit
set -o nounset
set -o pipefail

# Install golangci-lint
echo "Installing golangci-lint"
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.41.1

echo "running golangci-lint..."
./bin/golangci-lint run
