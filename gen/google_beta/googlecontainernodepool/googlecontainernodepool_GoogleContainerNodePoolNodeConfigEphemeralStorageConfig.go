package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigEphemeralStorageConfig struct {
	// Number of local SSDs to use to back ephemeral storage.
	//
	// Uses NVMe interfaces. Each local SSD is 375 GB in size.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_node_pool#local_ssd_count GoogleContainerNodePool#local_ssd_count}
	LocalSsdCount *float64 `field:"required" json:"localSsdCount" yaml:"localSsdCount"`
}

