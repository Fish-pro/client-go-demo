# client-go demo 
just one cluster
# get cluster-info
```go
 1008  ps | grep kube-apiserver
 1009  ps aux | grep kube-apiserver
 1010  cd /proc/2979/root/etc/ssl/
 1011  ll
 1012  cd private/
 1013  ll
 1014  kubectl config --help
 1015  kubectl config set-cluster --help
 1016  kubectl get ep
 1017  kubectl config set-cluster dce --server=https://10.6.124.52:16443 --certificate-authority=ca.crt --embed-certs --config /tmp/config
 1018  kubectl config set-cluster dce --server=https://10.6.124.52:16443 --certificate-authority=ca.crt --embed-certs --kubeconfig /tmp/config
 1019  cat /tmp/config
 1020  ll
 1021  kubectl config set-user
 1022  kubectl config
 1023  kubectl config  set-credentials --help
 1024  kubectl config set-credentials kube-admin --client-certificate kube-admin.crt --client-key kube-admin.key --embed-certs --kubeconfig /tmp/config
 1025  kubectl config
 1026  kubectl config set-context --help
 1027  kubectl config set-context --cluster dce --user kube-admin
 1028  kubectl config set-context dce --cluster dce --user kube-admin
 1029  kubectl config remove-context dce --cluster dce --user kube-admin
 1030  kubectl config del-context dce --cluster dce --user kube-admin
 1031  kubectl config
 1032  kubectl config delete-context dce --cluster dce --user kube-admin
 1033  kubectl config set-context dce --cluster dce --user kube-admin --kubeconfig /tmp/config
 1034  kubectl config use-context dce
 1035  kubectl config use-context dce --kubeconfig /tmp/config
 1036  kubectl --kubeconfig /tmp/config get no
 1037  kubectl config view
 1038  cat /tmp/config
```
