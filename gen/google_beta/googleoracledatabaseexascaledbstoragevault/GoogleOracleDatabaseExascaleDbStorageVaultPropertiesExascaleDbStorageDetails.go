package googleoracledatabaseexascaledbstoragevault


type GoogleOracleDatabaseExascaleDbStorageVaultPropertiesExascaleDbStorageDetails struct {
	// The total storage allocation for the ExascaleDbStorageVault, in gigabytes (GB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_exascale_db_storage_vault#total_size_gbs GoogleOracleDatabaseExascaleDbStorageVault#total_size_gbs}
	TotalSizeGbs *float64 `field:"required" json:"totalSizeGbs" yaml:"totalSizeGbs"`
}

