package googlestoragecontrolfolderintelligenceconfig


type GoogleStorageControlFolderIntelligenceConfigFilter struct {
	// excluded_cloud_storage_buckets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_storage_control_folder_intelligence_config#excluded_cloud_storage_buckets GoogleStorageControlFolderIntelligenceConfig#excluded_cloud_storage_buckets}
	ExcludedCloudStorageBuckets *GoogleStorageControlFolderIntelligenceConfigFilterExcludedCloudStorageBuckets `field:"optional" json:"excludedCloudStorageBuckets" yaml:"excludedCloudStorageBuckets"`
	// excluded_cloud_storage_locations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_storage_control_folder_intelligence_config#excluded_cloud_storage_locations GoogleStorageControlFolderIntelligenceConfig#excluded_cloud_storage_locations}
	ExcludedCloudStorageLocations *GoogleStorageControlFolderIntelligenceConfigFilterExcludedCloudStorageLocations `field:"optional" json:"excludedCloudStorageLocations" yaml:"excludedCloudStorageLocations"`
	// included_cloud_storage_buckets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_storage_control_folder_intelligence_config#included_cloud_storage_buckets GoogleStorageControlFolderIntelligenceConfig#included_cloud_storage_buckets}
	IncludedCloudStorageBuckets *GoogleStorageControlFolderIntelligenceConfigFilterIncludedCloudStorageBuckets `field:"optional" json:"includedCloudStorageBuckets" yaml:"includedCloudStorageBuckets"`
	// included_cloud_storage_locations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_storage_control_folder_intelligence_config#included_cloud_storage_locations GoogleStorageControlFolderIntelligenceConfig#included_cloud_storage_locations}
	IncludedCloudStorageLocations *GoogleStorageControlFolderIntelligenceConfigFilterIncludedCloudStorageLocations `field:"optional" json:"includedCloudStorageLocations" yaml:"includedCloudStorageLocations"`
}

