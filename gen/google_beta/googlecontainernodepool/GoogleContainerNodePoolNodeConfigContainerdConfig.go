package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigContainerdConfig struct {
	// private_registry_access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#private_registry_access_config GoogleContainerNodePool#private_registry_access_config}
	PrivateRegistryAccessConfig *GoogleContainerNodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfig `field:"optional" json:"privateRegistryAccessConfig" yaml:"privateRegistryAccessConfig"`
	// registry_hosts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#registry_hosts GoogleContainerNodePool#registry_hosts}
	RegistryHosts interface{} `field:"optional" json:"registryHosts" yaml:"registryHosts"`
	// writable_cgroups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#writable_cgroups GoogleContainerNodePool#writable_cgroups}
	WritableCgroups *GoogleContainerNodePoolNodeConfigContainerdConfigWritableCgroups `field:"optional" json:"writableCgroups" yaml:"writableCgroups"`
}

