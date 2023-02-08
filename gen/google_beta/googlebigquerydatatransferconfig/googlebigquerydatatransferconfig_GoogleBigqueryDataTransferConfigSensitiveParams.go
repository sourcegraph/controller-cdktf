package googlebigquerydatatransferconfig


type GoogleBigqueryDataTransferConfigSensitiveParams struct {
	// The Secret Access Key of the AWS account transferring data from.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_data_transfer_config#secret_access_key GoogleBigqueryDataTransferConfig#secret_access_key}
	SecretAccessKey *string `field:"required" json:"secretAccessKey" yaml:"secretAccessKey"`
}

