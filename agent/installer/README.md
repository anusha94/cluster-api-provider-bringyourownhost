# Building a BYOH Bundle
```sh
cd agent/installer/bundle_builder
docker build -t byoh-build-push-bundle .
```

Build a BYOH bundle and publish to an OCI compliant repo
In the below command, we will push to Harbor registry. You can push to any OCI registry of your choice.
The supported environment variables are -
BUILD_ONLY
# If set to 1 bundle is built and available as bundle/bundle.tar
# If set to 0 bundle is build and pushed to repo

CONTAINERD_VERSION
KUBERNETES_VERSION
```sh
docker run --rm -v `pwd`/byoh-ingredients-download:/ingredients --env BUILD_ONLY=0 --env CONTAINERD_VERSION=1.5.7 --env KUBERNETES_VERSION=1.21.2-00 build-push-bundle projects.registry.vmware.com/cluster_api_provider_bringyourownhost/byoh-bundle-ubuntu_20.04.1_x86-64_k8s:v1.21.2
```