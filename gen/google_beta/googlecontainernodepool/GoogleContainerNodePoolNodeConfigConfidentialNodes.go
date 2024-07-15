package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigConfidentialNodes struct {
	// Whether Confidential Nodes feature is enabled for all nodes in this pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_node_pool#enabled GoogleContainerNodePool#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

