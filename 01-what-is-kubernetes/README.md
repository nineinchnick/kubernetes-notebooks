What is Kubernetes
==================

After finishing this notebook, you'll learn:
* What problem Kubernetes solves?
* What are the components of Kubernetes?
* How to interact with a Kubernetes cluster?

## Overview

The official documentation starts with this question - [What is Kubernetes](https://kubernetes.io/docs/concepts/overview/what-is-kubernetes/).
But it might be easier to grasp it after seeing the [Kubernetes Components](https://kubernetes.io/docs/concepts/overview/components/) it consists of.
But if all of this is already overwhelming, don't worry, it's not required to know all of this to do the following excercise.

## Excercise

First excercise is creating a Kubernetes cluster. You can either do this:
* using `minikube`, [in your web broswer](https://kubernetes.io/docs/tutorials/kubernetes-basics/create-cluster/cluster-interactive/)
* or using `kind` locally, on your laptop

* [install kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/)
* [install kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)

### Create a cluster

```bash
# list existing clusters, there should be none
kind get clusters

# create a new cluster
kind create cluster

# list existing clusters, you should see a new one, with a default name
kind get clusters
```

### Look around the cluster

```bash
kubectl get all
kubectl get all -A
kubectl get namespaces
kubectl get pods
kubectl get pods -A
kubectl describe pod
kubectl get pods -o yaml

# get specific pods using label selectors
kubectl get pod -l k8s-app=kube-dns
```

> Tip: tired of writing `kubectl` over and over again? Create
> a shorter alias using `alias k=kubectl`!

### Change something on the cluster

```bash
# delete a pod
kubectl delete pod -l k8s-app=kube-dns
kubectl get pods
# watch pods for changes
kubectl get pods -w
# or run kubectl every few seconds
watch kubectl get pods

# now delete the whole cluster
kind delete cluster
# try creating a new cluster again
kind create cluster
```

### Explore Kubernetes components

```bash
# exec into the node
docker exec -it kind-control-plane bash
# see all the Kubernetes components running
ps aux | grep kube
# some of them are containers, but running using containerd, not Docker
ctr -n k8s.io containers ls
```

## Test

1. Which Kubernetes component talks to the Docker daemon? On which nodes it runs?
1. Why control-plane components should not run on worker nodes?
1. What does `kind` mean?

## Extras

Try creating two different kind clusters at once. How to talk to both at once?

Try creating a multi-node cluster. Which components run only on worker nodes?
