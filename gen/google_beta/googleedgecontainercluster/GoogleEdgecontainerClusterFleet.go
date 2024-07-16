package googleedgecontainercluster


type GoogleEdgecontainerClusterFleet struct {
	// The name of the Fleet host project where this cluster will be registered. Project names are formatted as 'projects/<project-number>'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_edgecontainer_cluster#project GoogleEdgecontainerCluster#project}
	Project *string `field:"required" json:"project" yaml:"project"`
}

