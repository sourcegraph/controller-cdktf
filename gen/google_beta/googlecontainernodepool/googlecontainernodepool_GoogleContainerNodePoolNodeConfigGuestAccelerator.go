package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigGuestAccelerator struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_node_pool#count GoogleContainerNodePool#count}.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_node_pool#gpu_partition_size GoogleContainerNodePool#gpu_partition_size}.
	GpuPartitionSize *string `field:"optional" json:"gpuPartitionSize" yaml:"gpuPartitionSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_node_pool#type GoogleContainerNodePool#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

