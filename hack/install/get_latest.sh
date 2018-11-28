#!/usr/bin/env bash


set -ex

# Expects supergloo to be installed in the standard place
SUPERGLOO_DIR="${GOPATH}/src/github.com/solo-io/supergloo"
HACK_DIR="${SUPERGLOO_DIR}/hack/install"
TMP_YML="${HACK_DIR}/supergloo.tmp.yml"

cd $SUPERGLOO_DIR

sh $HACK_DIR/recompile.sh

sed 's/imagePullPolicy: Always/imagePullPolicy: IfNotPresent/g' $HACK_DIR/supergloo.yaml > $TMP_YML
kubectl apply -f $TMP_YML
make install-cli

rm $TMP_YML

