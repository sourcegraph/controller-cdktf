package googlestorageinsightsreportconfig


type GoogleStorageInsightsReportConfigFrequencyOptionsEndDate struct {
	// The day of the month to stop generating inventory reports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_insights_report_config#day GoogleStorageInsightsReportConfig#day}
	Day *float64 `field:"required" json:"day" yaml:"day"`
	// The month to stop generating inventory reports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_insights_report_config#month GoogleStorageInsightsReportConfig#month}
	Month *float64 `field:"required" json:"month" yaml:"month"`
	// The year to stop generating inventory reports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_insights_report_config#year GoogleStorageInsightsReportConfig#year}
	Year *float64 `field:"required" json:"year" yaml:"year"`
}

