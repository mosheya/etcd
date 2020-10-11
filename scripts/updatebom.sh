#!/usr/bin/env bash

set -e

if ! [[ "$0" =~ scripts/updatebom.sh ]]; then
	echo "must be run from repository root"
	exit 255
fi

echo "installing 'bill-of-materials.json'"
go get -v -u github.com/mosheya/license-bill-of-materials

echo "generating bill-of-materials.json"
license-bill-of-materials \
    --override-file ./bill-of-materials.override.json \
    github.com/mosheya/etcd github.com/mosheya/etcd/etcdctl > bill-of-materials.json

echo "generated bill-of-materials.json"
