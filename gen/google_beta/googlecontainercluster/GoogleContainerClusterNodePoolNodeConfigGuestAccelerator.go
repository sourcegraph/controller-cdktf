package googlecontainercluster


type GoogleContainerClusterNodePoolNodeConfigGuestAccelerator struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_container_cluster#count GoogleContainerCluster#count}.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_container_cluster#gpu_partition_size GoogleContainerCluster#gpu_partition_size}.
	GpuPartitionSize *string `field:"optional" json:"gpuPartitionSize" yaml:"gpuPartitionSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_container_cluster#gpu_sharing_config GoogleContainerCluster#gpu_sharing_config}.
	GpuSharingConfig interface{} `field:"optional" json:"gpuSharingConfig" yaml:"gpuSharingConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_container_cluster#type GoogleContainerCluster#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

