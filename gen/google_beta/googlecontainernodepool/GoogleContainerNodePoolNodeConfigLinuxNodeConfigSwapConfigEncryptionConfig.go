package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEncryptionConfig struct {
	// If true, swap space will not be encrypted. Defaults to false (encrypted).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#disabled GoogleContainerNodePool#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
}

