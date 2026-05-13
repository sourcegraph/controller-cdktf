package containercluster


type ContainerClusterNodeConfigContainerdConfigRegistryHostsHostsClientKey struct {
	// URI for the Secret Manager secret that hosts the private key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#gcp_secret_manager_secret_uri ContainerCluster#gcp_secret_manager_secret_uri}
	GcpSecretManagerSecretUri *string `field:"optional" json:"gcpSecretManagerSecretUri" yaml:"gcpSecretManagerSecretUri"`
}

