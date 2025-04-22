package googletpuv2vm


type GoogleTpuV2VmSchedulingConfig struct {
	// Defines whether the node is preemptible.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_tpu_v2_vm#preemptible GoogleTpuV2Vm#preemptible}
	Preemptible interface{} `field:"optional" json:"preemptible" yaml:"preemptible"`
	// Whether the node is created under a reservation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_tpu_v2_vm#reserved GoogleTpuV2Vm#reserved}
	Reserved interface{} `field:"optional" json:"reserved" yaml:"reserved"`
	// Optional. Defines whether the node is Spot VM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_tpu_v2_vm#spot GoogleTpuV2Vm#spot}
	Spot interface{} `field:"optional" json:"spot" yaml:"spot"`
}

