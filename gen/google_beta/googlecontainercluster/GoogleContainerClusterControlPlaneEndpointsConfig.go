package googlecontainercluster


type GoogleContainerClusterControlPlaneEndpointsConfig struct {
	// dns_endpoint_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_container_cluster#dns_endpoint_config GoogleContainerCluster#dns_endpoint_config}
	DnsEndpointConfig *GoogleContainerClusterControlPlaneEndpointsConfigDnsEndpointConfig `field:"optional" json:"dnsEndpointConfig" yaml:"dnsEndpointConfig"`
	// ip_endpoints_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_container_cluster#ip_endpoints_config GoogleContainerCluster#ip_endpoints_config}
	IpEndpointsConfig *GoogleContainerClusterControlPlaneEndpointsConfigIpEndpointsConfig `field:"optional" json:"ipEndpointsConfig" yaml:"ipEndpointsConfig"`
}

