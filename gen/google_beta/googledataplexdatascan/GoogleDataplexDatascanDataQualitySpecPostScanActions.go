package googledataplexdatascan


type GoogleDataplexDatascanDataQualitySpecPostScanActions struct {
	// bigquery_export block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_datascan#bigquery_export GoogleDataplexDatascan#bigquery_export}
	BigqueryExport *GoogleDataplexDatascanDataQualitySpecPostScanActionsBigqueryExport `field:"optional" json:"bigqueryExport" yaml:"bigqueryExport"`
	// notification_report block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_datascan#notification_report GoogleDataplexDatascan#notification_report}
	NotificationReport *GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReport `field:"optional" json:"notificationReport" yaml:"notificationReport"`
}

