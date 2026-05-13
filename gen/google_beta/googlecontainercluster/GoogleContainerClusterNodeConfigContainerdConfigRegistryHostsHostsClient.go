package googlecontainercluster


type GoogleContainerClusterNodeConfigContainerdConfigRegistryHostsHostsClient struct {
	// cert block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#cert GoogleContainerCluster#cert}
	Cert *GoogleContainerClusterNodeConfigContainerdConfigRegistryHostsHostsClientCert `field:"required" json:"cert" yaml:"cert"`
	// key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#key GoogleContainerCluster#key}
	Key *GoogleContainerClusterNodeConfigContainerdConfigRegistryHostsHostsClientKey `field:"optional" json:"key" yaml:"key"`
}

