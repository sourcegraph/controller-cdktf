package containernodepool


type ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfile struct {
	// The number of physical local NVMe SSD disks to attach.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_node_pool#disk_count ContainerNodePool#disk_count}
	DiskCount *float64 `field:"optional" json:"diskCount" yaml:"diskCount"`
}

