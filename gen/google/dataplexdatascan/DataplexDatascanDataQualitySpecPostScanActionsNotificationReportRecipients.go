package dataplexdatascan


type DataplexDatascanDataQualitySpecPostScanActionsNotificationReportRecipients struct {
	// The email recipients who will receive the DataQualityScan results report.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/dataplex_datascan#emails DataplexDatascan#emails}
	Emails *[]*string `field:"optional" json:"emails" yaml:"emails"`
}

