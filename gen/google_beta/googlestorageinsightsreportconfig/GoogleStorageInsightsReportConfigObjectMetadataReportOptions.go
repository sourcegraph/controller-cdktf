package googlestorageinsightsreportconfig


type GoogleStorageInsightsReportConfigObjectMetadataReportOptions struct {
	// The metadata fields included in an inventory report.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_insights_report_config#metadata_fields GoogleStorageInsightsReportConfig#metadata_fields}
	MetadataFields *[]*string `field:"required" json:"metadataFields" yaml:"metadataFields"`
	// storage_destination_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_insights_report_config#storage_destination_options GoogleStorageInsightsReportConfig#storage_destination_options}
	StorageDestinationOptions *GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptions `field:"required" json:"storageDestinationOptions" yaml:"storageDestinationOptions"`
	// storage_filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_insights_report_config#storage_filters GoogleStorageInsightsReportConfig#storage_filters}
	StorageFilters *GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageFilters `field:"optional" json:"storageFilters" yaml:"storageFilters"`
}

