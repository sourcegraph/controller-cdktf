package googlecontainercluster


type GoogleContainerClusterNodeConfigContainerdConfig struct {
	// private_registry_access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#private_registry_access_config GoogleContainerCluster#private_registry_access_config}
	PrivateRegistryAccessConfig *GoogleContainerClusterNodeConfigContainerdConfigPrivateRegistryAccessConfig `field:"optional" json:"privateRegistryAccessConfig" yaml:"privateRegistryAccessConfig"`
	// registry_hosts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#registry_hosts GoogleContainerCluster#registry_hosts}
	RegistryHosts interface{} `field:"optional" json:"registryHosts" yaml:"registryHosts"`
	// writable_cgroups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#writable_cgroups GoogleContainerCluster#writable_cgroups}
	WritableCgroups *GoogleContainerClusterNodeConfigContainerdConfigWritableCgroups `field:"optional" json:"writableCgroups" yaml:"writableCgroups"`
}

