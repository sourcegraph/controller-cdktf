package googlestorageinsightsreportconfig


type GoogleStorageInsightsReportConfigFrequencyOptions struct {
	// end_date block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_insights_report_config#end_date GoogleStorageInsightsReportConfig#end_date}
	EndDate *GoogleStorageInsightsReportConfigFrequencyOptionsEndDate `field:"required" json:"endDate" yaml:"endDate"`
	// The frequency in which inventory reports are generated. Values are DAILY or WEEKLY. Possible values: ["DAILY", "WEEKLY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_insights_report_config#frequency GoogleStorageInsightsReportConfig#frequency}
	Frequency *string `field:"required" json:"frequency" yaml:"frequency"`
	// start_date block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_insights_report_config#start_date GoogleStorageInsightsReportConfig#start_date}
	StartDate *GoogleStorageInsightsReportConfigFrequencyOptionsStartDate `field:"required" json:"startDate" yaml:"startDate"`
}

