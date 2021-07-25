Debug a pod
===========

After finishing this notebook, you'll learn:
* How to troubleshoot a pod that doesn't start

## Overview

Kubernetes documentation has a complete section dedicated to
[debugging pods](https://kubernetes.io/docs/tasks/debug-application-cluster/debug-running-pod/).

## Excercise

```bash
# apply a deployment with a wrong configmap mounted
kubectl create -f app.yaml
# check logs
kubectl logs -l app=config-printer
# change entrypoint, add `command: ["bash", "-c", "sleep 10000"]` to the container, next to its name
kubectl edit deploy/config-printer
# get list of pods and exec into the latest one
kubectl get pods
kubectl exec -it config-printer-ABCDE-XYZ -- bash
# check where the config is mounted
ls /etc/config*

# do all of the above with a single command:
kubectl debug config-printer-ABCDE-XYZ -it --copy-to=debug --container=app -- bash -c "sleep 10000"
kubectl get pod debug
kubectl exec -it debug -- sh
```

## Test

1. What can be the reasons for a pod to fail to start?
1. What `kubectl` command allows to look around in the container?

## Extras

1. What's the difference between a container command and args.
1. How to debug a container built from scratch, that doesn't have bash?
