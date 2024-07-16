package googlestorageinsightsreportconfig


type GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptions struct {
	// The destination bucket that stores the generated inventory reports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_insights_report_config#bucket GoogleStorageInsightsReportConfig#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The path within the destination bucket to store generated inventory reports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_insights_report_config#destination_path GoogleStorageInsightsReportConfig#destination_path}
	DestinationPath *string `field:"optional" json:"destinationPath" yaml:"destinationPath"`
}

