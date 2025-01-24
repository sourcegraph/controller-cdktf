package googlecomputeinstancefromtemplate


type GoogleComputeInstanceFromTemplateConfidentialInstanceConfig struct {
	// The confidential computing technology the instance uses.
	//
	// SEV is an AMD feature. TDX is an Intel feature. One of the following
	// values is required: SEV, SEV_SNP, TDX. If SEV_SNP, min_cpu_platform =
	// "AMD Milan" is currently required. TDX is only available in beta.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_compute_instance_from_template#confidential_instance_type GoogleComputeInstanceFromTemplate#confidential_instance_type}
	ConfidentialInstanceType *string `field:"optional" json:"confidentialInstanceType" yaml:"confidentialInstanceType"`
	// Defines whether the instance should have confidential compute enabled. Field will be deprecated in a future release.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_compute_instance_from_template#enable_confidential_compute GoogleComputeInstanceFromTemplate#enable_confidential_compute}
	EnableConfidentialCompute interface{} `field:"optional" json:"enableConfidentialCompute" yaml:"enableConfidentialCompute"`
}

