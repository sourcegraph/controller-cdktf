package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigTaint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#effect GoogleContainerNodePool#effect}.
	Effect *string `field:"optional" json:"effect" yaml:"effect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#key GoogleContainerNodePool#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#value GoogleContainerNodePool#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

