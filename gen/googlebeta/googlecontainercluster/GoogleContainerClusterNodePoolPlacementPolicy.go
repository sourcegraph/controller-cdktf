package googlecontainercluster


type GoogleContainerClusterNodePoolPlacementPolicy struct {
	// Type defines the type of placement policy.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#type GoogleContainerCluster#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

