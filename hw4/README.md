# Hw4
## Run in kind:
```
kind create cluster --config=kind-config.yaml
helm install --create-namespace -n <namespace> <name> .
```
## Run in own cluster:
```
helm install --create-namespace -n <namespace> --set db.node_name=<node_name> <name> .
```
node_name is the name of the node where persistent volume for the database will be deployed
