package googledataprocworkflowtemplate


type GoogleDataprocWorkflowTemplatePlacementClusterSelector struct {
	// Required. The cluster labels. Cluster must have all labels to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#cluster_labels GoogleDataprocWorkflowTemplate#cluster_labels}
	ClusterLabels *map[string]*string `field:"required" json:"clusterLabels" yaml:"clusterLabels"`
	// Optional.
	//
	// The zone where workflow process executes. This parameter does not affect the selection of the cluster. If unspecified, the zone of the first cluster matching the selector is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#zone GoogleDataprocWorkflowTemplate#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

