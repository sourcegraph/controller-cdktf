package googlecontainercluster


type GoogleContainerClusterNodeConfigEphemeralStorageLocalSsdConfig struct {
	// Number of local SSDs to use to back ephemeral storage.
	//
	// Uses NVMe interfaces. Each local SSD must be 375 or 3000 GB in size, and all local SSDs must share the same size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#local_ssd_count GoogleContainerCluster#local_ssd_count}
	LocalSsdCount *float64 `field:"required" json:"localSsdCount" yaml:"localSsdCount"`
	// Number of local SSDs to be utilized for GKE Data Cache. Uses NVMe interfaces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#data_cache_count GoogleContainerCluster#data_cache_count}
	DataCacheCount *float64 `field:"optional" json:"dataCacheCount" yaml:"dataCacheCount"`
}

