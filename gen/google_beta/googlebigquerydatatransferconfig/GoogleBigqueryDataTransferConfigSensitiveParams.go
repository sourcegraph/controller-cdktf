package googlebigquerydatatransferconfig


type GoogleBigqueryDataTransferConfigSensitiveParams struct {
	// The Secret Access Key of the AWS account transferring data from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_data_transfer_config#secret_access_key GoogleBigqueryDataTransferConfig#secret_access_key}
	SecretAccessKey *string `field:"optional" json:"secretAccessKey" yaml:"secretAccessKey"`
	// The Secret Access Key of the AWS account transferring data from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_data_transfer_config#secret_access_key_wo GoogleBigqueryDataTransferConfig#secret_access_key_wo}
	SecretAccessKeyWo *string `field:"optional" json:"secretAccessKeyWo" yaml:"secretAccessKeyWo"`
	// The version of the sensitive params - used to trigger updates of the write-only params.
	//
	// For more info see [updating write-only arguments](/docs/providers/google/guides/using_write_only_arguments.html#updating-write-only-arguments)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_data_transfer_config#secret_access_key_wo_version GoogleBigqueryDataTransferConfig#secret_access_key_wo_version}
	SecretAccessKeyWoVersion *float64 `field:"optional" json:"secretAccessKeyWoVersion" yaml:"secretAccessKeyWoVersion"`
}

