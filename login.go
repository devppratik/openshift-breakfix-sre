package main

import (
	"fmt"

	bplogin "github.com/openshift/backplane-cli/cmd/ocm-backplane/login"
	bpconfig "github.com/openshift/backplane-cli/pkg/cli/config"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func getKubeConfigAndClient(clusterID string) (client.Client, *rest.Config, *kubernetes.Clientset, error) {
	bp, err := bpconfig.GetBackplaneConfiguration()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to load backplane-cli config: %v", err)
	}

	kubeconfig, err := bplogin.GetRestConfigAsUser(bp, clusterID, "backplane-cluster-admin")
	if err != nil {
		return nil, nil, nil, err
	}
	fmt.Println("[SUCCESS] Logged in to Cluster")
	// create the clientset
	clientset, err := kubernetes.NewForConfig(kubeconfig)
	if err != nil {
		return nil, nil, nil, err
	}
	kubeCli, err := client.New(kubeconfig, client.Options{})
	if err != nil {
		return nil, nil, nil, err
	}
	return kubeCli, kubeconfig, clientset, err
}

for cluster_id in $(ocm get subscriptions --parameter "search=organization_id='1pwwsfazToamNegaehP6eaDg80K'" | jq -r .items | jq -r '.[].cluster_id'); do
  # Check cluster support status
  echo "Checking cluster $cluster_id"
  cluster_support_status=$(osdctl cluster support status -S $cluster_id)
  	if [[ "$cluster_support_status" != "Cluster is fully supported" ]]; then
    echo $cluster_support_status
  fi
  echo ""
done
