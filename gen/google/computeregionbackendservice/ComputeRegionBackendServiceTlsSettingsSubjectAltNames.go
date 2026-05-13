package computeregionbackendservice


type ComputeRegionBackendServiceTlsSettingsSubjectAltNames struct {
	// The SAN specified as a DNS Name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_region_backend_service#dns_name ComputeRegionBackendService#dns_name}
	DnsName *string `field:"optional" json:"dnsName" yaml:"dnsName"`
	// The SAN specified as a URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_region_backend_service#uniform_resource_identifier ComputeRegionBackendService#uniform_resource_identifier}
	UniformResourceIdentifier *string `field:"optional" json:"uniformResourceIdentifier" yaml:"uniformResourceIdentifier"`
}

