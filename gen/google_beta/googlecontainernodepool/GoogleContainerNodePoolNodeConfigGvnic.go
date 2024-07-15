package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigGvnic struct {
	// Whether or not gvnic is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_node_pool#enabled GoogleContainerNodePool#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

