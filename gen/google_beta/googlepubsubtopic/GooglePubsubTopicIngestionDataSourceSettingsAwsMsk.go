package googlepubsubtopic


type GooglePubsubTopicIngestionDataSourceSettingsAwsMsk struct {
	// AWS role ARN to be used for Federated Identity authentication with MSK.
	//
	// Check the Pub/Sub docs for how to set up this role and the
	// required permissions that need to be attached to it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_pubsub_topic#aws_role_arn GooglePubsubTopic#aws_role_arn}
	AwsRoleArn *string `field:"required" json:"awsRoleArn" yaml:"awsRoleArn"`
	// ARN that uniquely identifies the MSK cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_pubsub_topic#cluster_arn GooglePubsubTopic#cluster_arn}
	ClusterArn *string `field:"required" json:"clusterArn" yaml:"clusterArn"`
	// The GCP service account to be used for Federated Identity authentication with MSK (via a 'AssumeRoleWithWebIdentity' call for the provided role).
	//
	// The 'awsRoleArn' must be set up with 'accounts.google.com:sub'
	// equals to this service account number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_pubsub_topic#gcp_service_account GooglePubsubTopic#gcp_service_account}
	GcpServiceAccount *string `field:"required" json:"gcpServiceAccount" yaml:"gcpServiceAccount"`
	// The name of the MSK topic that Pub/Sub will import from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_pubsub_topic#topic GooglePubsubTopic#topic}
	Topic *string `field:"required" json:"topic" yaml:"topic"`
}

