package googlepubsubtopic


type GooglePubsubTopicIngestionDataSourceSettingsAwsKinesis struct {
	// AWS role ARN to be used for Federated Identity authentication with Kinesis.
	//
	// Check the Pub/Sub docs for how to set up this role and the
	// required permissions that need to be attached to it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_pubsub_topic#aws_role_arn GooglePubsubTopic#aws_role_arn}
	AwsRoleArn *string `field:"required" json:"awsRoleArn" yaml:"awsRoleArn"`
	// The Kinesis consumer ARN to used for ingestion in Enhanced Fan-Out mode.
	//
	// The consumer must be already
	// created and ready to be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_pubsub_topic#consumer_arn GooglePubsubTopic#consumer_arn}
	ConsumerArn *string `field:"required" json:"consumerArn" yaml:"consumerArn"`
	// The GCP service account to be used for Federated Identity authentication with Kinesis (via a 'AssumeRoleWithWebIdentity' call for the provided role).
	//
	// The 'awsRoleArn' must be set up with 'accounts.google.com:sub'
	// equals to this service account number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_pubsub_topic#gcp_service_account GooglePubsubTopic#gcp_service_account}
	GcpServiceAccount *string `field:"required" json:"gcpServiceAccount" yaml:"gcpServiceAccount"`
	// The Kinesis stream ARN to ingest data from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_pubsub_topic#stream_arn GooglePubsubTopic#stream_arn}
	StreamArn *string `field:"required" json:"streamArn" yaml:"streamArn"`
}

