Scylla operator advantage over statefulsets deployment.
Basic operations using scylla operator.
Scylla production config doc.
Scylla production deployment steps.
Scylla benching results.

### Scylla operator based deployments has below advantage over statefulset based deployment:
1. Scylla operator can put scylla node in maintenance mode whenever when we want to do some memory.
2. Scylla operator automatically clean up dead node. If dead is not removed or repaired other scylla nodes will try to sync data with dead node.
3. Scylla operator based deployment automatically configured scylla manager client (side car docker) so no additional config steps required to set up scylla manager with scylla nodes.
4. Upgradation of a scylla cluster is handled by taking system schema snapshot of all scylla node.

### System config required for scylla:
1. Scylla can handle 12.5k ops per physical core.
2. 2GB memory is needed per core. i.e. for 4vCPU / 4 cores CPU, 8GB memory is needed. 
3. Memory to storage ratio is 30:1 means for 4GB memory 120GB storage is required.
4. Scylla sizing calculator: https://price-calc.gh.scylladb.com/
* Scylla uses memory for caching of data. More detail here: 

### Steps to deploy using EKS:
1. Create 3 node group with taint: scylla-clusters, 1 node for scylla monitoring stack and operator stack.
2. Make changes resources like cpu and memory in `cluster.yaml` file.
3. Run `deploy.sh` file.
* `deploy.sh` file will deploy *node-setup-daemonset*, *provisioner*, *cert-manager*, *operator* and *scylla clusters*.
4. When scylla cluster deployment is completed then verify connection by accessing database:
```
kubectl exec -n scylla -it simple-cluster-us-east-1-us-east-1a-0 -- cqlsh
```
5. If everything worked as expected next step is to set up monitoring stack.
6. Add monitoring stack charts repository:
```
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
```
7. Install monitoring stack:
```
helm install monitoring prometheus-community/kube-prometheus-stack --values monitoring/values.yaml --create-namespace --namespace scylla-monitoring
```
8. Install service monitors:
```
kubectl apply -f monitoring/scylla-service-monitor.yaml
```
9. Download dashboards:
```
wget https://github.com/scylladb/scylla-monitoring/archive/scylla-monitoring-3.6.0.tar.gz
tar -xvf scylla-monitoring-3.6.0.tar.gz
```
10. Install dashboards:
```
# Scylla dashboards
kubectl -n scylla-monitoring create configmap scylla-dashboards --from-file=scylla-monitoring-scylla-monitoring-3.6.0/grafana/build/ver_4.3
kubectl -n scylla-monitoring patch configmap scylla-dashboards  -p '{"metadata":{"labels":{"grafana_dashboard": "1"}}}'
```
11. To access dashboard configure ingress:
```
kubectl -n scylla-monitoring port-forward deployment.apps/monitoring-grafana 3000
```