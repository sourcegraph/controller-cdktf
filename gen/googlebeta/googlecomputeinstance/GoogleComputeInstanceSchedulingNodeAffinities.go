package googlecomputeinstance


type GoogleComputeInstanceSchedulingNodeAffinities struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance#key GoogleComputeInstance#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance#operator GoogleComputeInstance#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance#values GoogleComputeInstance#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

