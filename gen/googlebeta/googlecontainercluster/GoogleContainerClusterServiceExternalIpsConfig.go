package googlecontainercluster


type GoogleContainerClusterServiceExternalIpsConfig struct {
	// When enabled, services with exterenal ips specified will be allowed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#enabled GoogleContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

