package containercluster


type ContainerClusterControlPlaneEndpointsConfigDnsEndpointConfig struct {
	// Controls whether user traffic is allowed over this endpoint.
	//
	// Note that GCP-managed services may still use the endpoint even if this is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#allow_external_traffic ContainerCluster#allow_external_traffic}
	AllowExternalTraffic interface{} `field:"optional" json:"allowExternalTraffic" yaml:"allowExternalTraffic"`
	// Controls whether the k8s certs auth is allowed via dns.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#enable_k8s_certs_via_dns ContainerCluster#enable_k8s_certs_via_dns}
	EnableK8SCertsViaDns interface{} `field:"optional" json:"enableK8SCertsViaDns" yaml:"enableK8SCertsViaDns"`
	// Controls whether the k8s token auth is allowed via dns.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#enable_k8s_tokens_via_dns ContainerCluster#enable_k8s_tokens_via_dns}
	EnableK8STokensViaDns interface{} `field:"optional" json:"enableK8STokensViaDns" yaml:"enableK8STokensViaDns"`
	// The cluster's DNS endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#endpoint ContainerCluster#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
}

