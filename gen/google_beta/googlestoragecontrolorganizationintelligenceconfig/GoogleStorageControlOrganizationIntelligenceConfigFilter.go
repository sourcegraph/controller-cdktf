package googlestoragecontrolorganizationintelligenceconfig


type GoogleStorageControlOrganizationIntelligenceConfigFilter struct {
	// excluded_cloud_storage_buckets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_storage_control_organization_intelligence_config#excluded_cloud_storage_buckets GoogleStorageControlOrganizationIntelligenceConfig#excluded_cloud_storage_buckets}
	ExcludedCloudStorageBuckets *GoogleStorageControlOrganizationIntelligenceConfigFilterExcludedCloudStorageBuckets `field:"optional" json:"excludedCloudStorageBuckets" yaml:"excludedCloudStorageBuckets"`
	// excluded_cloud_storage_locations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_storage_control_organization_intelligence_config#excluded_cloud_storage_locations GoogleStorageControlOrganizationIntelligenceConfig#excluded_cloud_storage_locations}
	ExcludedCloudStorageLocations *GoogleStorageControlOrganizationIntelligenceConfigFilterExcludedCloudStorageLocations `field:"optional" json:"excludedCloudStorageLocations" yaml:"excludedCloudStorageLocations"`
	// included_cloud_storage_buckets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_storage_control_organization_intelligence_config#included_cloud_storage_buckets GoogleStorageControlOrganizationIntelligenceConfig#included_cloud_storage_buckets}
	IncludedCloudStorageBuckets *GoogleStorageControlOrganizationIntelligenceConfigFilterIncludedCloudStorageBuckets `field:"optional" json:"includedCloudStorageBuckets" yaml:"includedCloudStorageBuckets"`
	// included_cloud_storage_locations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_storage_control_organization_intelligence_config#included_cloud_storage_locations GoogleStorageControlOrganizationIntelligenceConfig#included_cloud_storage_locations}
	IncludedCloudStorageLocations *GoogleStorageControlOrganizationIntelligenceConfigFilterIncludedCloudStorageLocations `field:"optional" json:"includedCloudStorageLocations" yaml:"includedCloudStorageLocations"`
}

