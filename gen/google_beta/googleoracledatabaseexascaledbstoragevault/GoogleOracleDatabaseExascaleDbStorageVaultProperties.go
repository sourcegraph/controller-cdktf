package googleoracledatabaseexascaledbstoragevault


type GoogleOracleDatabaseExascaleDbStorageVaultProperties struct {
	// exascale_db_storage_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_exascale_db_storage_vault#exascale_db_storage_details GoogleOracleDatabaseExascaleDbStorageVault#exascale_db_storage_details}
	ExascaleDbStorageDetails *GoogleOracleDatabaseExascaleDbStorageVaultPropertiesExascaleDbStorageDetails `field:"required" json:"exascaleDbStorageDetails" yaml:"exascaleDbStorageDetails"`
	// The size of additional flash cache in percentage of high capacity database storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_exascale_db_storage_vault#additional_flash_cache_percent GoogleOracleDatabaseExascaleDbStorageVault#additional_flash_cache_percent}
	AdditionalFlashCachePercent *float64 `field:"optional" json:"additionalFlashCachePercent" yaml:"additionalFlashCachePercent"`
	// time_zone block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_exascale_db_storage_vault#time_zone GoogleOracleDatabaseExascaleDbStorageVault#time_zone}
	TimeZone *GoogleOracleDatabaseExascaleDbStorageVaultPropertiesTimeZone `field:"optional" json:"timeZone" yaml:"timeZone"`
}

