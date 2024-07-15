package googlecontainercluster


type GoogleContainerClusterServiceExternalIpsConfig struct {
	// When enabled, services with external ips specified will be allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#enabled GoogleContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

