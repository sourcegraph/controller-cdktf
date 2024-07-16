package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigTaint struct {
	// Effect for taint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_node_pool#effect GoogleContainerNodePool#effect}
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// Key for taint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_node_pool#key GoogleContainerNodePool#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Value for taint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_node_pool#value GoogleContainerNodePool#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

