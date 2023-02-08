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

### Steps to configure password authenticator in Scylla
1. Open CQLSH session and execute below query:
```
ALTER KEYSPACE system_auth WITH REPLICATION =
  { 'class' : 'SimpleStrategy', 'replication_factor' : <number_of_nodes> };
```
2. Add **PasswordAuthentication** in scylla config.
```
kubectl create configmap scylla-config -n scylla --from-file=scylla.yaml
```
3. Wait for mount to propagate and restart the cluster.
```
kubectl rollout restart -n scylla statefulset/<deployment_name>
```
* Check operator log to be sure config applied successfully
4. Use following to access cqlsh in any pod
```
cqlsh -u cassandra -p cassandra
```
5. `system_auth` is a keyspace which contains all data related to authetication. Execute this command to ensure all data is sync in all nodes of cluster.
```
nodetool repair system_auth
```
6. Delete default superuser:
```
CREATE ROLE <user> WITH PASSWORD = 'password' AND SUPERUSER = true AND LOGIN = true;
DROP ROLE IF EXISTS  'cassandra';
```
More on authorization and authentication :
https://docs.scylladb.com/stable/operating-scylla/security/enable-authorization.html#authorizer


### Steps to set up Scylla manager
1. Deploy scylla manager for backup and repair related tasks. Make sure Scylla manager is not deployed on Scylla cluster Nodes. If you want to deploy it with operator and monitoring node make sure you increase the CPU and memory.
```
k apply -f manager.yaml
```
2. After scylla manager deployment is completed check Scylla Cluster registered with Scylla Manager or not.
```
k -n scylla describe ScyllaCluster
```
Check Manager Id output of above command.
3. Now set up periodic back up in S3 bucket. Either attach AWS IAM policy or set up secret for AWS bucket so that scylla agent can upload files to bucket. Update `access_key_id` and `secret_access_key` in the `scylla-manager-agent.yaml` file.
```
kubectl create secret -n scylla generic scylla-agent-config-secret --from-file scylla-manager-agent.yaml
```
4. Restart the cluster
```
kubectl rollout restart -n scylla statefulset/<deployment_name>
```
after restart you can do bash in scylla-manager-agent and check changes in .yaml file.
```
k exec -n scylla <pod_name> -ti -c scylla-manager-agent -- bash 
cat /mnt/scylla-agent-config/scylla-manager-agent.yaml
```
5. To schedule a backup, Add following keys under specs:
execute:
```
k -n scylla edit ScyllaCluster simple-cluster
```
then add this to configuration:
```
  repairs:
  - name: "weekly us-east-1 repair"
    intensity: "2"
    interval: "7d"
    dc: ["us-east-1"]
  backups:
  - name: "daily users backup"
    retention: 3
    location: ["s3:<BUCKET_NAME>"]
    interval: "1d"
  - name: "weekly full cluster backup"
    retention: 7
    location: ["s3:<BUCKET_NAME>"]
    interval: "7d"
```
6. Check whether task was successfully assigned or not:
```
kubectl -n scylla-manager exec -ti scylla-manager-scylla-manager-7bd9f968b9-w25jw -- sctool task list
```

### Steps to set up monitoring in Scylla DB
1. Add monitoring stack charts repository:
```
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
```
2. Install monitoring stack:
```
helm install monitoring prometheus-community/kube-prometheus-stack --values monitoring/values.yaml --create-namespace --namespace scylla-monitoring
```
3. Install service monitors:
```
kubectl apply -f monitoring/scylla-service-monitor.yaml
kubectl apply -f monitoring/scylla-manager-service-monitor.yaml
```
4. Download dashboards:
```
wget https://github.com/scylladb/scylla-monitoring/archive/scylla-monitoring-3.6.0.tar.gz
tar -xvf scylla-monitoring-3.6.0.tar.gz
```
5. Install dashboards:
```
# Scylla dashboards
kubectl -n scylla-monitoring create configmap scylla-dashboards --from-file=scylla-monitoring-scylla-monitoring-3.6.0/grafana/build/ver_4.3
kubectl -n scylla-monitoring patch configmap scylla-dashboards  -p '{"metadata":{"labels":{"grafana_dashboard": "1"}}}'

kubectl -n scylla-monitoring create configmap scylla-manager-dashboards --from-file=scylla-monitoring-scylla-monitoring-3.6.0/grafana/build/manager_2.2
kubectl -n scylla-monitoring patch configmap scylla-manager-dashboards  -p '{"metadata":{"labels":{"grafana_dashboard": "1"}}}'
```
5. To access dashboard configure ingress:
```
kubectl -n scylla-monitoring port-forward deployment.apps/monitoring-grafana 3000
```

