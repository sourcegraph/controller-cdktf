package googlecomputeinstancefromtemplate


type GoogleComputeInstanceFromTemplateSchedulingNodeAffinities struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_from_template#key GoogleComputeInstanceFromTemplate#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_from_template#operator GoogleComputeInstanceFromTemplate#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_from_template#values GoogleComputeInstanceFromTemplate#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

