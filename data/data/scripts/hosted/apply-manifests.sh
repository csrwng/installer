#!/bin/bash

oc create configmap hosted-manifests --from-file=./hosted-manifests
oc create -f ./manifest-applier-pod.yaml
