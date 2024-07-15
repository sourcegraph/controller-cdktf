package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigContainerdConfig struct {
	// private_registry_access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_node_pool#private_registry_access_config GoogleContainerNodePool#private_registry_access_config}
	PrivateRegistryAccessConfig *GoogleContainerNodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfig `field:"optional" json:"privateRegistryAccessConfig" yaml:"privateRegistryAccessConfig"`
}

