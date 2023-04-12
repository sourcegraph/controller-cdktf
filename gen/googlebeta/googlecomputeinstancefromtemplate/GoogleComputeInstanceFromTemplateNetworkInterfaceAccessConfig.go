package googlecomputeinstancefromtemplate


type GoogleComputeInstanceFromTemplateNetworkInterfaceAccessConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_from_template#nat_ip GoogleComputeInstanceFromTemplate#nat_ip}.
	NatIp *string `field:"optional" json:"natIp" yaml:"natIp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_from_template#network_tier GoogleComputeInstanceFromTemplate#network_tier}.
	NetworkTier *string `field:"optional" json:"networkTier" yaml:"networkTier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_from_template#public_ptr_domain_name GoogleComputeInstanceFromTemplate#public_ptr_domain_name}.
	PublicPtrDomainName *string `field:"optional" json:"publicPtrDomainName" yaml:"publicPtrDomainName"`
}

