package googlecontainernodepool


type GoogleContainerNodePoolPlacementPolicy struct {
	// Type defines the type of placement policy.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_node_pool#type GoogleContainerNodePool#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

