Create a pod
============

After finishing this notebook, you'll learn:
* How to write a Kubernetes resource manifest in yaml
* What's the difference between `kubectl apply` and `kubectl create`

## Overview

Kubernetes documentation describes [the concept of a pod](https://kubernetes.io/docs/reference/kubectl/cheatsheet/),
which, in short, is a group of containers, sharing same network and storage. So creating a pod is the simplest way
to run an application in a container on Kubernetes.

## Excercise

```bash
kubectl run --image nginx app
kubectl get pods
# this will keep running, so open up a second terminal
kubectl port-forward app 8080:80
firefox http://localhost:8080
# compare the pod created by the `kubectl run` command with a minimal example
kubectl get pod app -o yaml > $TMPDIR/run-app.yaml
diff $TMPDIR/run-app.yaml 03-create-a-pod/pod.yaml

# compare if the result is the same
kubectl delete pod app
kubectl apply -f 03-create-a-pod/pod.yaml
kubectl port-forward app 8080:80
firefox http://localhost:8080
kubectl get pod app -o yaml > $TMPDIR/manifest-app.yaml
diff $TMPDIR/manifest-app.yaml 03-create-a-pod/pod.yaml

kubectl logs app
# this might throw an error
kubectl apply -f 03-create-a-pod/pod-pages.yaml
# recreate the pod
kubectl delete pod app
kubectl apply -f 03-create-a-pod/pod-pages.yaml
```

## Test

1. Why Kubernetes adds other fields to a pod definition?
1. Does port-forward always needs two port numbers?
1. Why some parts of a pod cannot be updated?

## Extras

1. Try to create a secret using only the `kubectl create` command.
1. Find out what's the largest possible file a secret can hold.
