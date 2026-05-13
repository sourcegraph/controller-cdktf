package googlecomputeregionbackendservice


type GoogleComputeRegionBackendServiceTlsSettingsSubjectAltNames struct {
	// The SAN specified as a DNS Name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_backend_service#dns_name GoogleComputeRegionBackendService#dns_name}
	DnsName *string `field:"optional" json:"dnsName" yaml:"dnsName"`
	// The SAN specified as a URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_backend_service#uniform_resource_identifier GoogleComputeRegionBackendService#uniform_resource_identifier}
	UniformResourceIdentifier *string `field:"optional" json:"uniformResourceIdentifier" yaml:"uniformResourceIdentifier"`
}

