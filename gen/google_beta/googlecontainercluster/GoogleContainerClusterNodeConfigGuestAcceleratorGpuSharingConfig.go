package googlecontainercluster


type GoogleContainerClusterNodeConfigGuestAcceleratorGpuSharingConfig struct {
	// The type of GPU sharing strategy to enable on the GPU node.
	//
	// Possible values are described in the API package (https://pkg.go.dev/google.golang.org/api/container/v1#GPUSharingConfig)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_container_cluster#gpu_sharing_strategy GoogleContainerCluster#gpu_sharing_strategy}
	GpuSharingStrategy *string `field:"required" json:"gpuSharingStrategy" yaml:"gpuSharingStrategy"`
	// The maximum number of containers that can share a GPU.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_container_cluster#max_shared_clients_per_gpu GoogleContainerCluster#max_shared_clients_per_gpu}
	MaxSharedClientsPerGpu *float64 `field:"required" json:"maxSharedClientsPerGpu" yaml:"maxSharedClientsPerGpu"`
}

