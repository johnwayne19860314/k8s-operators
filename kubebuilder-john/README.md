# kubebuilder-john
// TODO(user): Add simple overview of use/purpose

## Description
// TODO(user): An in-depth paragraph about your project and overview of use

## create api
    kubebuilder create api --group john --version v1beta1 --kind App
## use web-hook
    kubebuilder create webhook --group john --version v1beta1 --kind App --defaulting --programmatic-validation
    api/v1beta1/app_webhook.go webhook对应的handler，我们添加业务逻辑的地方

    api/v1beta1/webhook_suite_test.go 测试

    config/certmanager 自动生成自签名的证书，用于webhook server提供https服务

    config/webhook 用于注册webhook到k8s中

    config/crd/patches 为conversion自动注入caBoundle

    config/default/manager_webhook_patch.yaml 让manager的deployment支持webhook请求

    config/default/webhookcainjection_patch.yaml 为webhook server注入caBoundle


    kubectl get secrets cert-manager-webhook-ca -n cert-manager -o jsonpath='{..tls\.crt}' |base64 -d > certs/tls.crt
    kubectl get secrets cert-manager-webhook-ca -n cert-manager -o jsonpath='{..tls\.key}' |base64 -d > certs/tls.key
## Getting Started

### Prerequisites
- go version v1.20.0+
- docker version 17.03+.
- kubectl version v1.11.3+.
- Access to a Kubernetes v1.11.3+ cluster.

### To Deploy on the cluster
**Build and push your image to the location specified by `IMG`:**

```sh
make docker-build docker-push IMG=<some-registry>/kubebuilder-john:tag
```

**NOTE:** This image ought to be published in the personal registry you specified. 
And it is required to have access to pull the image from the working environment. 
Make sure you have the proper permission to the registry if the above commands don’t work.

**Install the CRDs into the cluster:**

```sh
make install
```

**Deploy the Manager to the cluster with the image specified by `IMG`:**

```sh
make deploy IMG=<some-registry>/kubebuilder-john:tag
```

> **NOTE**: If you encounter RBAC errors, you may need to grant yourself cluster-admin 
privileges or be logged in as admin.

**Create instances of your solution**
You can apply the samples (examples) from the config/sample:

```sh
kubectl apply -k config/samples/
```

>**NOTE**: Ensure that the samples has default values to test it out.

### To Uninstall
**Delete the instances (CRs) from the cluster:**

```sh
kubectl delete -k config/samples/
```

**Delete the APIs(CRDs) from the cluster:**

```sh
make uninstall
```

**UnDeploy the controller from the cluster:**

```sh
make undeploy
```

## Contributing
// TODO(user): Add detailed information on how you would like others to contribute to this project

**NOTE:** Run `make --help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

