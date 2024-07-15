package googlecomputeinstancefrommachineimage


type GoogleComputeInstanceFromMachineImageSchedulingNodeAffinities struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_machine_image#key GoogleComputeInstanceFromMachineImage#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_machine_image#operator GoogleComputeInstanceFromMachineImage#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_machine_image#values GoogleComputeInstanceFromMachineImage#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

