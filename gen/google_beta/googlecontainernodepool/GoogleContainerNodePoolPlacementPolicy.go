package googlecontainernodepool


type GoogleContainerNodePoolPlacementPolicy struct {
	// Type defines the type of placement policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#type GoogleContainerNodePool#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

