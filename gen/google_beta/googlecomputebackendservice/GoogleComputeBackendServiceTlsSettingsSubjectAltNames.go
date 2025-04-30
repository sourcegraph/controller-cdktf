package googlecomputebackendservice


type GoogleComputeBackendServiceTlsSettingsSubjectAltNames struct {
	// The SAN specified as a DNS Name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_backend_service#dns_name GoogleComputeBackendService#dns_name}
	DnsName *string `field:"optional" json:"dnsName" yaml:"dnsName"`
	// The SAN specified as a URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_backend_service#uniform_resource_identifier GoogleComputeBackendService#uniform_resource_identifier}
	UniformResourceIdentifier *string `field:"optional" json:"uniformResourceIdentifier" yaml:"uniformResourceIdentifier"`
}

