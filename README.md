# client-go demo 
It's a `client-go` demo by golang, only a single cluster is supported
## get cluster-info
If you already have a cluster, you can get cluster connection information in the following ways
```shell script
ps aux | grep kube-apiserver
cd /proc/2979/root/etc/ssl/private/
kubectl get ep
kubectl config set-cluster dce --server=https://10.6.124.52:16443 --certificate-authority=ca.crt --embed-certs --kubeconfig /tmp/config
cat /tmp/config
kubectl config set-user
kubectl config set-credentials kube-admin --client-certificate kube-admin.crt --client-key kube-admin.key --embed-certs --kubeconfig /tmp/config
kubectl config set-context dce --cluster dce --user kube-admin --kubeconfig /tmp/config
kubectl config use-context dce --kubeconfig /tmp/config
kubectl --kubeconfig /tmp/config get no
kubectl config view
cat /tmp/config
```
## unit test
The test only contains GET method
```shell script
go test ./pkg/test
```
## struct
```shell script
├── Dockerfile
├── README.md
├── cmd
│   └── appserver
│       └── main.go
├── config
│   ├── 32.yaml
│   ├── 52.yaml
│   ├── 55.yaml
│   └── config.go
├── docs
│   ├── client-go-demo.postman_collection.json
│   └── client-go-demo.postman_environment.json
├── go.mod
├── go.sum
└── pkg
    ├── app
    │   ├── deployment
    │   │   ├── handler.go
    │   │   ├── router.go
    │   │   └── service.go
    │   ├── node
    │   │   ├── handler.go
    │   │   ├── router.go
    │   │   └── service.go
    │   ├── pod
    │   │   ├── handler.go
    │   │   ├── router.go
    │   │   └── servcie.go
    │   └── service
    │       ├── handler.go
    │       ├── router.go
    │       └── service.go
    ├── middleware
    │   ├── corsmiddleware.go
    │   ├── logmiddleware.go
    │   ├── reqidmiddleware.go
    │   └── setup.go
    ├── test
    │   ├── demo_test.go
    │   ├── deploy_test.go
    │   ├── node_test.go
    │   ├── pod_test.go
    │   ├── service_test.go
    │   └── test.go
    └── util
        ├── log.go
        └── util.go
```