package googledataproccluster


type GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfig struct {
	// gke_cluster_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#gke_cluster_config GoogleDataprocCluster#gke_cluster_config}
	GkeClusterConfig *GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig `field:"required" json:"gkeClusterConfig" yaml:"gkeClusterConfig"`
	// kubernetes_software_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#kubernetes_software_config GoogleDataprocCluster#kubernetes_software_config}
	KubernetesSoftwareConfig *GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig `field:"required" json:"kubernetesSoftwareConfig" yaml:"kubernetesSoftwareConfig"`
	// A namespace within the Kubernetes cluster to deploy into.
	//
	// If this namespace does not exist, it is created. If it exists, Dataproc verifies that another Dataproc VirtualCluster is not installed into it. If not specified, the name of the Dataproc Cluster is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#kubernetes_namespace GoogleDataprocCluster#kubernetes_namespace}
	KubernetesNamespace *string `field:"optional" json:"kubernetesNamespace" yaml:"kubernetesNamespace"`
}

