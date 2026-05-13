package oracledatabaseexascaledbstoragevault


type OracleDatabaseExascaleDbStorageVaultProperties struct {
	// exascale_db_storage_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_exascale_db_storage_vault#exascale_db_storage_details OracleDatabaseExascaleDbStorageVault#exascale_db_storage_details}
	ExascaleDbStorageDetails *OracleDatabaseExascaleDbStorageVaultPropertiesExascaleDbStorageDetails `field:"required" json:"exascaleDbStorageDetails" yaml:"exascaleDbStorageDetails"`
	// The size of additional flash cache in percentage of high capacity database storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_exascale_db_storage_vault#additional_flash_cache_percent OracleDatabaseExascaleDbStorageVault#additional_flash_cache_percent}
	AdditionalFlashCachePercent *float64 `field:"optional" json:"additionalFlashCachePercent" yaml:"additionalFlashCachePercent"`
	// time_zone block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_exascale_db_storage_vault#time_zone OracleDatabaseExascaleDbStorageVault#time_zone}
	TimeZone *OracleDatabaseExascaleDbStorageVaultPropertiesTimeZone `field:"optional" json:"timeZone" yaml:"timeZone"`
}

