package googlepubsubtopic


type GooglePubsubTopicIngestionDataSourceSettingsConfluentCloud struct {
	// The Confluent Cloud bootstrap server. The format is url:port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_pubsub_topic#bootstrap_server GooglePubsubTopic#bootstrap_server}
	BootstrapServer *string `field:"required" json:"bootstrapServer" yaml:"bootstrapServer"`
	// The GCP service account to be used for Federated Identity authentication with Confluent Cloud.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_pubsub_topic#gcp_service_account GooglePubsubTopic#gcp_service_account}
	GcpServiceAccount *string `field:"required" json:"gcpServiceAccount" yaml:"gcpServiceAccount"`
	// Identity pool ID to be used for Federated Identity authentication with Confluent Cloud.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_pubsub_topic#identity_pool_id GooglePubsubTopic#identity_pool_id}
	IdentityPoolId *string `field:"required" json:"identityPoolId" yaml:"identityPoolId"`
	// Name of the Confluent Cloud topic that Pub/Sub will import from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_pubsub_topic#topic GooglePubsubTopic#topic}
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// The Confluent Cloud cluster ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_pubsub_topic#cluster_id GooglePubsubTopic#cluster_id}
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
}

