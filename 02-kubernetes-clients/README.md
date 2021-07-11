Kubernets clients
=================

After finishing this notebook, you'll learn:
* What other alternative Kubernetes clients exist
* How can they make it easier to work with Kubernetes

## Overview

Some graphical clients:
* [the Dashboard](https://kubernetes.io/docs/tasks/access-application-cluster/web-ui-dashboard/) - the official web UI
* [Lens](https://k8slens.dev/) - the Kubernetes IDE - an alternative GUI

## Excercise

### Dashboard

```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.2.0/aio/deploy/recommended.yaml
kubectl proxy
```

```bash
cat <<EOF >dashboard-user.yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: admin-user
  namespace: kubernetes-dashboard
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: admin-user
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
- kind: ServiceAccount
  name: admin-user
  namespace: kubernetes-dashboard
EOF
kubectl apply -f dashboard-user.yaml

kubectl -n kubernetes-dashboard get secret $(kubectl -n kubernetes-dashboard get sa/admin-user -o jsonpath="{.secrets[0].name}") -o go-template="{{.data.token | base64decode}}"
```

Deploy a new app:
1. Press "+"
1. Choose "Create from form"
1. Use `gcr.io/google-samples/hello-app:1.0` as the container image
1. Use an exteral service, choose any ports.
1. Check logs
1. Check associated services

### Lens

1. Download Lens from [the official site](https://k8slens.dev/) and install it.
1. Copy the Kubernetes config from `~/.kube/config`.

1. Find the app deployed using the dashboard
1. Edit the deployment, setting `.spec.replicas` to 2.
1. Scale the deployment to 100.

1. Check Network->Services, notice a LoadBalancer service in `Pending` state.

### CLI

```bash
# scale
kubectl scale deploy/hello --replicas 1
# wait for scalining to be done
kubectl rollout status deploy/hello -w
# expose the service!
kubectl create service nodeport hello2 --tcp 8080
# get the assigned random port
kubectl get service hello2 -o jsonpath='{.spec.ports[0].nodePort}'
```

## Test

1. Which is simpler to use, the CLI or GUI?
