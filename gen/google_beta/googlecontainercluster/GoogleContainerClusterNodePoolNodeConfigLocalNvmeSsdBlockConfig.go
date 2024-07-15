package googlecontainercluster


type GoogleContainerClusterNodePoolNodeConfigLocalNvmeSsdBlockConfig struct {
	// Number of raw-block local NVMe SSD disks to be attached to the node.
	//
	// Each local SSD is 375 GB in size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#local_ssd_count GoogleContainerCluster#local_ssd_count}
	LocalSsdCount *float64 `field:"required" json:"localSsdCount" yaml:"localSsdCount"`
}

