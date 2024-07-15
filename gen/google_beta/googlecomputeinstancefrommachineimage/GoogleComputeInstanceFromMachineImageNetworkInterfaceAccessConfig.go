package googlecomputeinstancefrommachineimage


type GoogleComputeInstanceFromMachineImageNetworkInterfaceAccessConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_machine_image#nat_ip GoogleComputeInstanceFromMachineImage#nat_ip}.
	NatIp *string `field:"optional" json:"natIp" yaml:"natIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_machine_image#network_tier GoogleComputeInstanceFromMachineImage#network_tier}.
	NetworkTier *string `field:"optional" json:"networkTier" yaml:"networkTier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_machine_image#public_ptr_domain_name GoogleComputeInstanceFromMachineImage#public_ptr_domain_name}.
	PublicPtrDomainName *string `field:"optional" json:"publicPtrDomainName" yaml:"publicPtrDomainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_machine_image#security_policy GoogleComputeInstanceFromMachineImage#security_policy}.
	SecurityPolicy *string `field:"optional" json:"securityPolicy" yaml:"securityPolicy"`
}

