package googlebigquerydatatransferconfig


type GoogleBigqueryDataTransferConfigSensitiveParams struct {
	// The Secret Access Key of the AWS account transferring data from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_data_transfer_config#secret_access_key GoogleBigqueryDataTransferConfig#secret_access_key}
	SecretAccessKey *string `field:"required" json:"secretAccessKey" yaml:"secretAccessKey"`
}

