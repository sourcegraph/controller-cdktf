package googlestorageinsightsreportconfig


type GoogleStorageInsightsReportConfigCsvOptions struct {
	// The delimiter used to separate the fields in the inventory report CSV file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_insights_report_config#delimiter GoogleStorageInsightsReportConfig#delimiter}
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// The boolean that indicates whether or not headers are included in the inventory report CSV file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_insights_report_config#header_required GoogleStorageInsightsReportConfig#header_required}
	HeaderRequired interface{} `field:"optional" json:"headerRequired" yaml:"headerRequired"`
	// The character used to separate the records in the inventory report CSV file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_insights_report_config#record_separator GoogleStorageInsightsReportConfig#record_separator}
	RecordSeparator *string `field:"optional" json:"recordSeparator" yaml:"recordSeparator"`
}

