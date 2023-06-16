package googlecontainercluster


type GoogleContainerClusterNodePoolPlacementPolicy struct {
	// Type defines the type of placement policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_cluster#type GoogleContainerCluster#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

