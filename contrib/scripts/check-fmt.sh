#!/usr/bin/env bash

set -e
set -o pipefail

diff="$(find . ! \( -path './contrib' -prune \) \
        ! \( -path './vendor' -prune \) \
        ! \( -path './_build' -prune \) \
        ! \( -path './.git' -prune \) \
        ! \( -path '*.validate.go' -prune \) \
        -type f -name '*.go' | grep -Ev "(pkg/k8s/apis/cilium.io/v2/client/bindata.go)" | \
        xargs gofmt -d -l -s )"

if [ -n "$diff" ]; then
	echo "Unformatted Go source code:"
	echo "$diff"
	exit 1
fi

exit 0
