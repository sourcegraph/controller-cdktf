package googlebigquerydatatransferconfig


type GoogleBigqueryDataTransferConfigEmailPreferences struct {
	// If true, email notifications will be sent on transfer run failures.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_data_transfer_config#enable_failure_email GoogleBigqueryDataTransferConfig#enable_failure_email}
	EnableFailureEmail interface{} `field:"required" json:"enableFailureEmail" yaml:"enableFailureEmail"`
}

