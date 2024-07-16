package googlestorageinsightsreportconfig


type GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageFilters struct {
	// The filter to use when specifying which bucket to generate inventory reports for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_insights_report_config#bucket GoogleStorageInsightsReportConfig#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
}

