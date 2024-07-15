package googlecontainercluster


type GoogleContainerClusterNodePoolNodeConfigContainerdConfig struct {
	// private_registry_access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#private_registry_access_config GoogleContainerCluster#private_registry_access_config}
	PrivateRegistryAccessConfig *GoogleContainerClusterNodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfig `field:"optional" json:"privateRegistryAccessConfig" yaml:"privateRegistryAccessConfig"`
}

