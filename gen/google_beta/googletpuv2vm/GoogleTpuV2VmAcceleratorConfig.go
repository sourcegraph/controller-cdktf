package googletpuv2vm


type GoogleTpuV2VmAcceleratorConfig struct {
	// Topology of TPU in chips.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_tpu_v2_vm#topology GoogleTpuV2Vm#topology}
	Topology *string `field:"required" json:"topology" yaml:"topology"`
	// Type of TPU. Possible values: ["V2", "V3", "V4", "V5P"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_tpu_v2_vm#type GoogleTpuV2Vm#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

