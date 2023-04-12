package googledataprocworkflowtemplate


type GoogleDataprocWorkflowTemplatePlacement struct {
	// cluster_selector block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_workflow_template#cluster_selector GoogleDataprocWorkflowTemplate#cluster_selector}
	ClusterSelector *GoogleDataprocWorkflowTemplatePlacementClusterSelector `field:"optional" json:"clusterSelector" yaml:"clusterSelector"`
	// managed_cluster block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_workflow_template#managed_cluster GoogleDataprocWorkflowTemplate#managed_cluster}
	ManagedCluster *GoogleDataprocWorkflowTemplatePlacementManagedCluster `field:"optional" json:"managedCluster" yaml:"managedCluster"`
}

