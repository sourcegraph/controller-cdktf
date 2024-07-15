package googledataprocworkflowtemplate


type GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget struct {
	// Optional. A namespace within the GKE cluster to deploy into.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#cluster_namespace GoogleDataprocWorkflowTemplate#cluster_namespace}
	ClusterNamespace *string `field:"optional" json:"clusterNamespace" yaml:"clusterNamespace"`
	// Optional. The target GKE cluster to deploy to. Format: 'projects/{project}/locations/{location}/clusters/{cluster_id}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#target_gke_cluster GoogleDataprocWorkflowTemplate#target_gke_cluster}
	TargetGkeCluster *string `field:"optional" json:"targetGkeCluster" yaml:"targetGkeCluster"`
}

