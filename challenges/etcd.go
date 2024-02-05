package challenges

import "fmt"

var alert = `
           Labels:
 - alertname = etcdMembersDown
 - job = etcd
 - namespace = openshift-etcd
 - openshift_io_alert_source = platform
 - prometheus = openshift-monitoring/k8s
 - service = etcd
 - severity = critical
Annotations:
 - description = etcd cluster "etcd": members are down (1).
 - runbook_url = https://github.com/openshift/runbooks/blob/master/alerts/cluster-etcd-operator/etcdMembersDown.md
 - summary = etcd cluster members are down.
`

func InitEtcdMembersDown() {
	fmt.Println("[INFO] Breaking cluster...")
	fmt.Println("[SUCCESS] Cluster is ready for breakfix")
	fmt.Println("You receive the following alert - ")
	fmt.Println(alert)
	fmt.Println("As Primary On Call you need to fix the problem that is present in the cluster. Good Luck!")
}

func BreakCluster() {

}

func ValidateCluster() {

}
