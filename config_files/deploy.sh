#!/bin/bash
set -euo pipefail

#########
# Start #
#########

check_prerequisites() {

    echo "Checking if kubectl is present on the machine..."
    if ! hash kubectl 2>/dev/null; then
        echo "You need to install kubectl. See: https://kubernetes.io/docs/tasks/tools/install-kubectl/"
        exit 1
    fi

    echo "Checking if helm is present on the machine..."
    if ! hash helm 2>/dev/null; then
        echo "You need to install helm. See: https://docs.helm.sh/using_helm/#installing-helm"
        exit 1
    fi
}

function wait-for-object-creation {
    for i in {1..30}; do
        { kubectl -n "${1}" get "${2}" && break; } || sleep 1
    done
}

# Check if the environment has the prerequisites installed
check_prerequisites

# Configure node disks and network
kubectl apply -f node-setup-daemonset.yaml
wait-for-object-creation default daemonset.apps/node-setup
kubectl rollout status --timeout=5m daemonset.apps/node-setup

# Install local volume provisioner
echo "Installing local volume provisioner..."
helm install local-provisioner provisioner
echo "Your disks are ready to use."

echo "Starting the cert manger..."
kubectl apply -f cert-manager.yaml
kubectl wait --for condition=established --timeout=60s crd/certificates.cert-manager.io crd/issuers.cert-manager.io
wait-for-object-creation cert-manager deployment.apps/cert-manager-webhook
kubectl -n cert-manager rollout status --timeout=5m deployment.apps/cert-manager-webhook

echo "Starting the scylla operator..."
kubectl apply -f operator.yaml
kubectl wait --for condition=established crd/scyllaclusters.scylla.scylladb.com
wait-for-object-creation scylla-operator deployment.apps/scylla-operator
kubectl -n scylla-operator rollout status --timeout=5m deployment.apps/scylla-operator

echo "Starting the scylla cluster..."
kubectl apply -f cluster.yaml
