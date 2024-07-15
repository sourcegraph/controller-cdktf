package googlecomputeregioncommitment


type GoogleComputeRegionCommitmentResources struct {
	// Name of the accelerator type resource. Applicable only when the type is ACCELERATOR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_commitment#accelerator_type GoogleComputeRegionCommitment#accelerator_type}
	AcceleratorType *string `field:"optional" json:"acceleratorType" yaml:"acceleratorType"`
	// The amount of the resource purchased (in a type-dependent unit, such as bytes).
	//
	// For vCPUs, this can just be an integer. For memory,
	// this must be provided in MB. Memory must be a multiple of 256 MB,
	// with up to 6.5GB of memory per every vCPU.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_commitment#amount GoogleComputeRegionCommitment#amount}
	Amount *string `field:"optional" json:"amount" yaml:"amount"`
	// Type of resource for which this commitment applies. Possible values are VCPU, MEMORY, LOCAL_SSD, and ACCELERATOR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_commitment#type GoogleComputeRegionCommitment#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

