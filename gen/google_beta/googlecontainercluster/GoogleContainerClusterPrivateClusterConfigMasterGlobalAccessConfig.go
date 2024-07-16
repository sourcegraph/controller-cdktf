package googlecontainercluster


type GoogleContainerClusterPrivateClusterConfigMasterGlobalAccessConfig struct {
	// Whether the cluster master is accessible globally or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#enabled GoogleContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

