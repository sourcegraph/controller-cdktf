package googlecomputeinstance


type GoogleComputeInstanceGuestAccelerator struct {
	// The number of the guest accelerator cards exposed to this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_instance#count GoogleComputeInstance#count}
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// The accelerator type resource exposed to this instance. E.g. nvidia-tesla-k80.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_instance#type GoogleComputeInstance#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

