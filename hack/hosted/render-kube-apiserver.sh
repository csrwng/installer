#!/bin/bash

RELEASE_IMAGE_DIGEST="registry.svc.ci.openshift.org/origin/release:4.2"
	
image_for() {
   docker run --rm --net=none "${RELEASE_IMAGE_DIGEST}" image "${1}"
}

KUBE_APISERVER_OPERATOR_IMAGE=$(image_for cluster-kube-apiserver-operator)
OPENSHIFT_HYPERKUBE_IMAGE=$(image_for hyperkube)

docker run  \
	--volume "$(pwd)/bootstrap/opt/openshift:/assets" \
	"${KUBE_APISERVER_OPERATOR_IMAGE}" \
	/usr/bin/cluster-kube-apiserver-operator render \
	--manifest-etcd-serving-ca=etcd-ca-bundle.crt \
	--manifest-etcd-server-urls=https://etcd-client:2379 \
	--manifest-image="${OPENSHIFT_HYPERKUBE_IMAGE}" \
	--manifest-operator-image="${KUBE_APISERVER_OPERATOR_IMAGE}" \
	--asset-input-dir=/assets/tls \
	--asset-output-dir=/assets/kube-apiserver-bootstrap \
	--config-output-file=/assets/kube-apiserver-bootstrap/config \
	--cluster-config-file=/assets/manifests/cluster-network-02-config.yml
