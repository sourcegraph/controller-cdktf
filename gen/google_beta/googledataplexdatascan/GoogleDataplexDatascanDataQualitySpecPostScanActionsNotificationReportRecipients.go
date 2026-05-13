package googledataplexdatascan


type GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportRecipients struct {
	// The email recipients who will receive the DataQualityScan results report.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_datascan#emails GoogleDataplexDatascan#emails}
	Emails *[]*string `field:"optional" json:"emails" yaml:"emails"`
}

