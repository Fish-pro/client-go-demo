# k8s安装

## 基本环境配置

### 1.关闭selinux

```shell
setenforce 0
sed -i "s/SELINUX=enforcing/SELINUX=disabled/g" /etc/selinux/config
```

### 2.关闭swap分区或禁用swap文件

```shell
swapoff -a
# 注释掉关于swap分区的行
yes | cp /etc/fstab /etc/fstab_bak
cat /etc/fstab_bak |grep -v swap > /etc/fstab
```

### 3.修改网卡配置

```shell
$ vim /etc/sysctl.conf
net.ipv4.ip_forward = 1
net.bridge.bridge-nf-call-iptables = 1
net.bridge.bridge-nf-call-ip6tables = 1
$ sysctl -p
```

### 4.启用内核模块

```shell
$ modprobe -- ip_vs
$ modprobe -- ip_vs_rr
$ modprobe -- ip_vs_wrr
$ modprobe -- ip_vs_sh
$ modprobe -- nf_conntrack_ipv4
$ cut -f1 -d " "  /proc/modules | grep -e ip_vs -e nf_conntrack_ipv4
$ vim /etc/sysconfig/modules/ipvs.modules
modprobe -- ip_vs
modprobe -- ip_vs_rr
modprobe -- ip_vs_wrr
modprobe -- ip_vs_sh
modprobe -- nf_conntrack_ipv4
```

### 5.关闭防火墙

```shell
$ systemctl stop firewalld
$ systemctl disable firewalld
```

### 6.配置hosts

#### kubectl、kubeadm、kubelet的安装

#### 添加Kubernetes的yum源

此处使用alibaba的镜像源

```shell
$ vim /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=http://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=0
repo_gpgcheck=0
gpgkey=http://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg
	http://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg
```

#### 安装kubelet、kubeadm、kubectl

```shell
$ yum install -y kubelet-1.20.1 kubeadm-1.20.1 kubectl-1.20.1
```

#### 启动kubelet服务

```shell
$ systemctl enable kubelet
$ systemctl start kubelet
```

此时执行`systemctl status kubelet`查看服务状态，服务状态应为Error(255)， 如果是其他错误可使用`journalctl -xe`查看错误信息。

## Docker安装和配置

### Docker安装

docker的安装请查看官网文档(Overview of Docker editions)[https://docs.docker.com/install/overview/]

#### OS requirements

To install Docker Engine, you need a maintained version of CentOS 7 or 8. Archived versions aren’t supported or tested.

The `centos-extras` repository must be enabled. This repository is enabled by default, but if you have disabled it, you need to [re-enable it](https://wiki.centos.org/AdditionalResources/Repositories).

The `overlay2` storage driver is recommended.

#### Uninstall old versions

Older versions of Docker were called `docker` or `docker-engine`. If these are installed, uninstall them, along with associated dependencies.

```shell
$ sudo yum remove docker \
                  docker-client \
                  docker-client-latest \
                  docker-common \
                  docker-latest \
                  docker-latest-logrotate \
                  docker-logrotate \
                  docker-engine
```

It’s OK if `yum` reports that none of these packages are installed.

The contents of `/var/lib/docker/`, including images, containers, volumes, and networks, are preserved. The Docker Engine package is now called `docker-ce`.

#### Installation methods

You can install Docker Engine in different ways, depending on your needs:

- Most users [set up Docker’s repositories](https://docs.docker.com/engine/install/centos/#install-using-the-repository) and install from them, for ease of installation and upgrade tasks. This is the recommended approach.
- Some users download the RPM package and [install it manually](https://docs.docker.com/engine/install/centos/#install-from-a-package) and manage upgrades completely manually. This is useful in situations such as installing Docker on air-gapped systems with no access to the internet.
- In testing and development environments, some users choose to use automated [convenience scripts](https://docs.docker.com/engine/install/centos/#install-using-the-convenience-script) to install Docker.

#### Install using the repository

Before you install Docker Engine for the first time on a new host machine, you need to set up the Docker repository. Afterward, you can install and update Docker from the repository.

#### SET UP THE REPOSITORY

Install the `yum-utils` package (which provides the `yum-config-manager` utility) and set up the **stable** repository.

```shell
$ sudo yum install -y yum-utils

$ sudo yum-config-manager \
    --add-repo \
    https://download.docker.com/linux/centos/docker-ce.repo
```

> **Optional**: Enable the **nightly** or **test** repositories.
>
> These repositories are included in the `docker.repo` file above but are disabled by default. You can enable them alongside the stable repository. The following command enables the **nightly** repository.
>
> ```shell
> $ sudo yum-config-manager --enable docker-ce-nightly
> ```
>
> To enable the **test** channel, run the following command:
>
> ```shell
> $ sudo yum-config-manager --enable docker-ce-test
> ```
>
> You can disable the **nightly** or **test** repository by running the `yum-config-manager` command with the `--disable` flag. To re-enable it, use the `--enable` flag. The following command disables the **nightly** repository.
>
> ```shell
> $ sudo yum-config-manager --disable docker-ce-nightly
> ```

#### INSTALL DOCKER ENGINE

1. ##### Install the *latest version* of Docker Engine and containerd, or go to the next step to install a specific version:

   ```shell
   $ sudo yum install docker-ce docker-ce-cli containerd.io
   ```

   If prompted to accept the GPG key, verify that the fingerprint matches `060A 61C5 1B55 8A7F 742B 77AA C52F EB6B 621E 9F35`, and if so, accept it.

   > Got multiple Docker repositories?
   >
   > If you have multiple Docker repositories enabled, installing or updating without specifying a version in the `yum install` or `yum update` command always installs the highest possible version, which may not be appropriate for your stability needs.

   Docker is installed but not started. The `docker` group is created, but no users are added to the group.

2. ##### To install a *specific version* of Docker Engine, list the available versions in the repo, then select and install:

   a. List and sort the versions available in your repo. This example sorts results by version number, highest to lowest, and is truncated:

   ```shell
   $ yum list docker-ce --showduplicates | sort -r
   
   docker-ce.x86_64  3:18.09.1-3.el7                     docker-ce-stable
   docker-ce.x86_64  3:18.09.0-3.el7                     docker-ce-stable
   docker-ce.x86_64  18.06.1.ce-3.el7                    docker-ce-stable
   docker-ce.x86_64  18.06.0.ce-3.el7                    docker-ce-stable
   ```

   The list returned depends on which repositories are enabled, and is specific to your version of CentOS (indicated by the `.el7` suffix in this example).

   b. Install a specific version by its fully qualified package name, which is the package name (`docker-ce`) plus the version string (2nd column) starting at the first colon (`:`), up to the first hyphen, separated by a hyphen (`-`). For example, `docker-ce-18.09.1`.

   ```shell
   $ sudo yum install docker-ce-<VERSION_STRING> docker-ce-cli-<VERSION_STRING> containerd.io
   ```

   Docker is installed but not started. The `docker` group is created, but no users are added to the group.

3. ##### Start Docker.

   ```shell
   $ sudo systemctl start docker
   ```

4. ##### Verify that Docker Engine is installed correctly by running the `hello-world` image.

   ```shell
   $ sudo docker run hello-world
   ```

   This command downloads a test image and runs it in a container. When the container runs, it prints an informational message and exits.

### Docker配置

1. #### 配置cgroup-driver为systemd

   ```shell
   # 查看cgroup-driver
   $ docker info | grep -i cgroup
   # 追加 --exec-opt native.cgroupdriver=systemd 参数
   $ sed -i "s#^ExecStart=/usr/bin/dockerd.*#ExecStart=/usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock --exec-opt native.cgroupdriver=systemd#g" /usr/lib/systemd/system/docker.service
   $ systemctl daemon-reload # 重新加载服务
   $ systemctl enable docker # 启用docker服务(开机自起)
   $ systemctl restart docker # 启动docker服务
   # 或者修改docker配置文件
   $ vim /etc/docker/daemon.json
   {
     "exec-opts": ["native.cgroupdriver=systemd"]
   }
   ```

2. #### 预先拉取所需镜像

   ```shell
   # 查看kubeadm所需镜像
   $ kubeadm config images list
   k8s.gcr.io/kube-apiserver:v1.16.3
   k8s.gcr.io/kube-controller-manager:v1.16.3
   k8s.gcr.io/kube-scheduler:v1.16.3
   k8s.gcr.io/kube-proxy:v1.16.3
   k8s.gcr.io/pause:3.1
   k8s.gcr.io/etcd:3.3.15-0
   k8s.gcr.io/coredns:1.6.2
   # 拉取镜像
   $ docker pull kubeimage/kube-apiserver-amd64:v1.16.3
   $ docker pull kubeimage/kube-controller-manager-amd64:v1.16.3
   $ docker pull kubeimage/kube-scheduler-amd64:v1.16.3
   $ docker pull kubeimage/kube-proxy-amd64:v1.16.3
   $ docker pull kubeimage/pause-amd64:3.1
   $ docker pull kubeimage/etcd-amd64:3.3.15-0
   $ docker pull coredns/coredns:1.6.2
   ```

3. #### 对预先拉取的镜像重新打tag

   ```shell
   $ docker tag kubeimage/kube-apiserver-amd64:v1.16.3  k8s.gcr.io/kube-apiserver:v1.16.3
   $ docker tag kubeimage/kube-controller-manager-amd64:v1.16.3  k8s.gcr.io/kube-controller-manager:v1.16.3
   $ docker tag kubeimage/kube-scheduler-amd64:v1.16.3  k8s.gcr.io/kube-scheduler:v1.16.3
   $ docker tag kubeimage/kube-proxy-amd64:v1.16.3 k8s.gcr.io/kube-proxy:v1.16.3
   $ docker tag kubeimage/pause-amd64:3.1 k8s.gcr.io/pause:3.1
   $ docker tag kubeimage/etcd-amd64:3.3.15-0 k8s.gcr.io/etcd:3.3.15-0
   $ docker tag coredns/coredns:1.6.2 k8s.gcr.io/coredns:1.6.2
   ```

## Master节点的配置

以上步骤需要在node节点和master节点执行，当前步骤仅需在master节点执行

### Master节点的初始化

```shell
# 初始化master节点，
# --pod-network-cidr=192.168.0.0/16 指定使用Calico网络
# --apiserver-advertise-address=10.0.0.5 指向master节点IP，此处也可以使用hosts
$ kubeadm init --pod-network-cidr=10.244.0.0/16 \
  --kubernetes-version=v1.16.3 \
  --apiserver-advertise-address=10.0.0.5
```

执行上述命令的输出为：

```shell
[init] Using Kubernetes version: v1.16.3
[preflight] Running pre-flight checks
	[WARNING SystemVerification]: this Docker version is not on the list of validated versions: 19.03.4. Latest validated version: 18.09
[preflight] Pulling images required for setting up a Kubernetes cluster
[preflight] This might take a minute or two, depending on the speed of your internet connection
[preflight] You can also perform this action in beforehand using 'kubeadm config images pull'
[kubelet-start] Writing kubelet environment file with flags to file "/var/lib/kubelet/kubeadm-flags.env"
[kubelet-start] Writing kubelet configuration to file "/var/lib/kubelet/config.yaml"
[kubelet-start] Activating the kubelet service
[certs] Using certificateDir folder "/etc/kubernetes/pki"
[certs] Generating "ca" certificate and key
[certs] Generating "apiserver" certificate and key
[certs] apiserver serving cert is signed for DNS names [master kubernetes kubernetes.default kubernetes.default.svc kubernetes.default.svc.cluster.local] and IPs [10.96.0.1 10.0.0.5]
[certs] Generating "apiserver-kubelet-client" certificate and key
[certs] Generating "front-proxy-ca" certificate and key
[certs] Generating "front-proxy-client" certificate and key
[certs] Generating "etcd/ca" certificate and key
[certs] Generating "etcd/server" certificate and key
[certs] etcd/server serving cert is signed for DNS names [master localhost] and IPs [10.0.0.5 127.0.0.1 ::1]
[certs] Generating "etcd/peer" certificate and key
[certs] etcd/peer serving cert is signed for DNS names [master localhost] and IPs [10.0.0.5 127.0.0.1 ::1]
[certs] Generating "etcd/healthcheck-client" certificate and key
[certs] Generating "apiserver-etcd-client" certificate and key
[certs] Generating "sa" key and public key
[kubeconfig] Using kubeconfig folder "/etc/kubernetes"
[kubeconfig] Writing "admin.conf" kubeconfig file
[kubeconfig] Writing "kubelet.conf" kubeconfig file
[kubeconfig] Writing "controller-manager.conf" kubeconfig file
[kubeconfig] Writing "scheduler.conf" kubeconfig file
[control-plane] Using manifest folder "/etc/kubernetes/manifests"
[control-plane] Creating static Pod manifest for "kube-apiserver"
[control-plane] Creating static Pod manifest for "kube-controller-manager"
[control-plane] Creating static Pod manifest for "kube-scheduler"
[etcd] Creating static Pod manifest for local etcd in "/etc/kubernetes/manifests"
[wait-control-plane] Waiting for the kubelet to boot up the control plane as static Pods from directory "/etc/kubernetes/manifests". This can take up to 4m0s
[apiclient] All control plane components are healthy after 13.002108 seconds
[upload-config] Storing the configuration used in ConfigMap "kubeadm-config" in the "kube-system" Namespace
[kubelet] Creating a ConfigMap "kubelet-config-1.16" in namespace kube-system with the configuration for the kubelets in the cluster
[upload-certs] Skipping phase. Please see --upload-certs
[mark-control-plane] Marking the node master as control-plane by adding the label "node-role.kubernetes.io/master=''"
[mark-control-plane] Marking the node master as control-plane by adding the taints [node-role.kubernetes.io/master:NoSchedule]
[bootstrap-token] Using token: kt58np.djd3youoqb0bnz4r
[bootstrap-token] Configuring bootstrap tokens, cluster-info ConfigMap, RBAC Roles
[bootstrap-token] configured RBAC rules to allow Node Bootstrap tokens to post CSRs in order for nodes to get long term certificate credentials
[bootstrap-token] configured RBAC rules to allow the csrapprover controller automatically approve CSRs from a Node Bootstrap Token
[bootstrap-token] configured RBAC rules to allow certificate rotation for all node client certificates in the cluster
[bootstrap-token] Creating the "cluster-info" ConfigMap in the "kube-public" namespace
[addons] Applied essential addon: CoreDNS
[addons] Applied essential addon: kube-proxy

Your Kubernetes control-plane has initialized successfully!

To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/

Then you can join any number of worker nodes by running the following on each as root:

kubeadm join 10.0.0.5:6443 --token kt58np.djd3youoqb0bnz4r \
    --discovery-token-ca-cert-hash sha256:37a3924142dc6d57eac2714e539c174ee3b0cda723746ada2464ac9e8a2091ce
```

保存输出中的`kubeadm join`部分内容，用于添加node节点，或者使用`kubeadm token list` 和`kubeadm token create --print-join-command`查看

```shell
$ kubeadm join 10.0.0.5:6443 --token kt58np.djd3youoqb0bnz4r \
		--discovery-token-ca-cert-hash sha256:37a3924142dc6d57eac2714e539c174ee3b0cda723746ada2464ac9e8a2091ce
```

接下来执行剩余的初始化步骤

```shell
$ mkdir -p $HOME/.kube
$ sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
$ sudo chown $(id -u):$(id -g) $HOME/.kube/config
```

#### Calico网络插件的配置

Calico的官方文档地址为： https://docs.projectcalico.org/v3.10/getting-started/kubernetes/。 具体安装步骤：

1. ##### 安装Calico

   ```shell
   $ kubectl apply -f https://docs.projectcalico.org/v3.10/manifests/calico.yaml
   ```

2. ##### 监听安装进度

   ```shell
   $ watch kubectl get pods --all-namespaces
   ```

出现以下内容时为安装成功

```shell
NAMESPACE    NAME                                       READY  STATUS   RESTARTS  AGE
kube-system  calico-kube-controllers-6ff88bf6d4-tgtzb   1/1    Running  0         2m45s
kube-system  calico-node-24h85                          1/1    Running  0         2m43s
kube-system  coredns-846jhw23g9-9af73                   1/1    Running  0         4m5s
kube-system  coredns-846jhw23g9-hmswk                   1/1    Running  0         4m5s
kube-system  etcd-jbaker-1                              1/1    Running  0         6m22s
kube-system  kube-apiserver-jbaker-1                    1/1    Running  0         6m12s
kube-system  kube-controller-manager-jbaker-1           1/1    Running  0         6m16s
kube-system  kube-proxy-8fzp2                           1/1    Running  0         5m16s
kube-system  kube-scheduler-jbaker-1                    1/1    Running  0         5m41s
```

测试

```shell
$ kubectl get nodes -o wide
NAME                STATUS     ROLES    AGE     VERSION   INTERNAL-IP      EXTERNAL-IP   OS-IMAGE                KERNEL-VERSION           CONTAINER-RUNTIME
kubernetes-master   Ready      master   4d12h   v1.16.3   192.168.56.101   <none>        CentOS Linux 7 (Core)   3.10.0-1062.el7.x86_64   docker://19.3.4
```

### Node节点的初始化

1. #### kubeadm获取join命令

   ```shell
   $ kubeadm token create --print-join-command
   ```

2. #### 登录node节点，执行加入集群的命令，完成加入集群操作

   ```shell
   $ kubeadm join 10.0.0.5:6443 --token kt58np.djd3youoqb0bnz4r \
       --discovery-token-ca-cert-hash sha256:37a3924142dc6d57eac2714e539c174ee3b0cda723746ada2464ac9e8a2091ce
   ```

3. #### 在master节点上查看添加结果

   ```shell
   $ kubectl get nodes -o wide
   NAME                STATUS     ROLES    AGE     VERSION   INTERNAL-IP      EXTERNAL-IP   OS-IMAGE                KERNEL-VERSION           CONTAINER-RUNTIME
   kubernetes-master   Ready      master   4d12h   v1.16.3   192.168.56.101   <none>        CentOS Linux 7 (Core)   3.10.0-1062.el7.x86_64   docker://19.3.4
   kubernetes-node-1   Ready      <none>   4d12h   v1.16.3   192.168.56.102   <none>        CentOS Linux 7 (Core)   3.10.0-1062.el7.x86_64   docker://19.3.4
   ```


