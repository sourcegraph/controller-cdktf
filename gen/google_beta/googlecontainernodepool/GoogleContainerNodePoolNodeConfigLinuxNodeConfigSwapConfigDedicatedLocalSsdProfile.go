package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfile struct {
	// The number of physical local NVMe SSD disks to attach.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#disk_count GoogleContainerNodePool#disk_count}
	DiskCount *float64 `field:"optional" json:"diskCount" yaml:"diskCount"`
}

