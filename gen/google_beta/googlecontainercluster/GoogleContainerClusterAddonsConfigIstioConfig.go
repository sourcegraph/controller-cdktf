package googlecontainercluster


type GoogleContainerClusterAddonsConfigIstioConfig struct {
	// The status of the Istio addon, which makes it easy to set up Istio for services in a cluster.
	//
	// It is disabled by default. Set disabled = false to enable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#disabled GoogleContainerCluster#disabled}
	Disabled interface{} `field:"required" json:"disabled" yaml:"disabled"`
	// The authentication type between services in Istio. Available options include AUTH_MUTUAL_TLS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#auth GoogleContainerCluster#auth}
	Auth *string `field:"optional" json:"auth" yaml:"auth"`
}

