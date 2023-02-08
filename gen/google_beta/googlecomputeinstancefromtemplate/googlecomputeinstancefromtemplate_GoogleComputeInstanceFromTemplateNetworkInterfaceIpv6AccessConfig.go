package googlecomputeinstancefromtemplate


type GoogleComputeInstanceFromTemplateNetworkInterfaceIpv6AccessConfig struct {
	// The service-level to be provided for IPv6 traffic when the subnet has an external subnet.
	//
	// Only PREMIUM tier is valid for IPv6
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_from_template#network_tier GoogleComputeInstanceFromTemplate#network_tier}
	NetworkTier *string `field:"required" json:"networkTier" yaml:"networkTier"`
	// The domain name to be used when creating DNSv6 records for the external IPv6 ranges.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_from_template#public_ptr_domain_name GoogleComputeInstanceFromTemplate#public_ptr_domain_name}
	PublicPtrDomainName *string `field:"optional" json:"publicPtrDomainName" yaml:"publicPtrDomainName"`
}

