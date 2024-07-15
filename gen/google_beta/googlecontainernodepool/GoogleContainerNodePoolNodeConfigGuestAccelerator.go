package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigGuestAccelerator struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_node_pool#count GoogleContainerNodePool#count}.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_node_pool#gpu_driver_installation_config GoogleContainerNodePool#gpu_driver_installation_config}.
	GpuDriverInstallationConfig interface{} `field:"optional" json:"gpuDriverInstallationConfig" yaml:"gpuDriverInstallationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_node_pool#gpu_partition_size GoogleContainerNodePool#gpu_partition_size}.
	GpuPartitionSize *string `field:"optional" json:"gpuPartitionSize" yaml:"gpuPartitionSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_node_pool#gpu_sharing_config GoogleContainerNodePool#gpu_sharing_config}.
	GpuSharingConfig interface{} `field:"optional" json:"gpuSharingConfig" yaml:"gpuSharingConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_node_pool#type GoogleContainerNodePool#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

