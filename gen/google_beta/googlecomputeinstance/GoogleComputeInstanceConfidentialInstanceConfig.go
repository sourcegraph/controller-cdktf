package googlecomputeinstance


type GoogleComputeInstanceConfidentialInstanceConfig struct {
	// Specifies which confidential computing technology to use.
	//
	// This could be one of the following values: SEV, SEV_SNP.
	// If SEV_SNP, min_cpu_platform = "AMD Milan" is currently required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance#confidential_instance_type GoogleComputeInstance#confidential_instance_type}
	ConfidentialInstanceType *string `field:"optional" json:"confidentialInstanceType" yaml:"confidentialInstanceType"`
	// Defines whether the instance should have confidential compute enabled. Field will be deprecated in a future release.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance#enable_confidential_compute GoogleComputeInstance#enable_confidential_compute}
	EnableConfidentialCompute interface{} `field:"optional" json:"enableConfidentialCompute" yaml:"enableConfidentialCompute"`
}

