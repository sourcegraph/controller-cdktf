package googlecontainercluster


type GoogleContainerClusterControlPlaneEndpointsConfigDnsEndpointConfig struct {
	// Controls whether user traffic is allowed over this endpoint.
	//
	// Note that GCP-managed services may still use the endpoint even if this is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_container_cluster#allow_external_traffic GoogleContainerCluster#allow_external_traffic}
	AllowExternalTraffic interface{} `field:"optional" json:"allowExternalTraffic" yaml:"allowExternalTraffic"`
	// The cluster's DNS endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_container_cluster#endpoint GoogleContainerCluster#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
}

