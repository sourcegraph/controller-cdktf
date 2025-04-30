package googletpuv2queuedresource


type GoogleTpuV2QueuedResourceTpuNodeSpecNode struct {
	// Runtime version for the TPU.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_tpu_v2_queued_resource#runtime_version GoogleTpuV2QueuedResource#runtime_version}
	RuntimeVersion *string `field:"required" json:"runtimeVersion" yaml:"runtimeVersion"`
	// TPU accelerator type for the TPU. If not specified, this defaults to 'v2-8'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_tpu_v2_queued_resource#accelerator_type GoogleTpuV2QueuedResource#accelerator_type}
	AcceleratorType *string `field:"optional" json:"acceleratorType" yaml:"acceleratorType"`
	// Text description of the TPU.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_tpu_v2_queued_resource#description GoogleTpuV2QueuedResource#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_tpu_v2_queued_resource#network_config GoogleTpuV2QueuedResource#network_config}
	NetworkConfig *GoogleTpuV2QueuedResourceTpuNodeSpecNodeNetworkConfig `field:"optional" json:"networkConfig" yaml:"networkConfig"`
}

