package googletpuv2vm


type GoogleTpuV2VmSchedulingConfig struct {
	// Defines whether the node is preemptible.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_tpu_v2_vm#preemptible GoogleTpuV2Vm#preemptible}
	Preemptible interface{} `field:"optional" json:"preemptible" yaml:"preemptible"`
	// Whether the node is created under a reservation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_tpu_v2_vm#reserved GoogleTpuV2Vm#reserved}
	Reserved interface{} `field:"optional" json:"reserved" yaml:"reserved"`
}

