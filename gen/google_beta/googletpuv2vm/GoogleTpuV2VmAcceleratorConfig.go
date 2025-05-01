package googletpuv2vm


type GoogleTpuV2VmAcceleratorConfig struct {
	// Topology of TPU in chips.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_tpu_v2_vm#topology GoogleTpuV2Vm#topology}
	Topology *string `field:"required" json:"topology" yaml:"topology"`
	// Type of TPU. Please select one of the allowed types: https://cloud.google.com/tpu/docs/reference/rest/v2/AcceleratorConfig#Type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_tpu_v2_vm#type GoogleTpuV2Vm#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

