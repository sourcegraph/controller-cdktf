package googlecomputeinstance


type GoogleComputeInstanceNetworkInterfaceIpv6AccessConfig struct {
	// The service-level to be provided for IPv6 traffic when the subnet has an external subnet.
	//
	// Only PREMIUM tier is valid for IPv6
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_compute_instance#network_tier GoogleComputeInstance#network_tier}
	NetworkTier *string `field:"required" json:"networkTier" yaml:"networkTier"`
	// The domain name to be used when creating DNSv6 records for the external IPv6 ranges.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_compute_instance#public_ptr_domain_name GoogleComputeInstance#public_ptr_domain_name}
	PublicPtrDomainName *string `field:"optional" json:"publicPtrDomainName" yaml:"publicPtrDomainName"`
}

