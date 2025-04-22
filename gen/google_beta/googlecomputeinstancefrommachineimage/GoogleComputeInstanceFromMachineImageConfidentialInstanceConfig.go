package googlecomputeinstancefrommachineimage


type GoogleComputeInstanceFromMachineImageConfidentialInstanceConfig struct {
	// The confidential computing technology the instance uses.
	//
	// SEV is an AMD feature. TDX is an Intel feature. One of the following
	// values is required: SEV, SEV_SNP, TDX. If SEV_SNP, min_cpu_platform =
	// "AMD Milan" is currently required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_instance_from_machine_image#confidential_instance_type GoogleComputeInstanceFromMachineImage#confidential_instance_type}
	ConfidentialInstanceType *string `field:"optional" json:"confidentialInstanceType" yaml:"confidentialInstanceType"`
	// Defines whether the instance should have confidential compute enabled. Field will be deprecated in a future release.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_instance_from_machine_image#enable_confidential_compute GoogleComputeInstanceFromMachineImage#enable_confidential_compute}
	EnableConfidentialCompute interface{} `field:"optional" json:"enableConfidentialCompute" yaml:"enableConfidentialCompute"`
}

