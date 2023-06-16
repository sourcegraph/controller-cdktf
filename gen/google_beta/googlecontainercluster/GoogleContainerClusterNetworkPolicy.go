package googlecontainercluster


type GoogleContainerClusterNetworkPolicy struct {
	// Whether network policy is enabled on the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_cluster#enabled GoogleContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The selected network policy provider. Defaults to PROVIDER_UNSPECIFIED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_cluster#provider GoogleContainerCluster#provider}
	Provider *string `field:"optional" json:"provider" yaml:"provider"`
}

