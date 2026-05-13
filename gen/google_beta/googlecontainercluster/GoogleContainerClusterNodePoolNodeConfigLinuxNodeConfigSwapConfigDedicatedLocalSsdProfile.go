package googlecontainercluster


type GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfile struct {
	// The number of physical local NVMe SSD disks to attach.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#disk_count GoogleContainerCluster#disk_count}
	DiskCount *float64 `field:"optional" json:"diskCount" yaml:"diskCount"`
}

