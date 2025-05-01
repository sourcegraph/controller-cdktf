package googlepubsubtopic


type GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubs struct {
	// The Azure event hub client ID to use for ingestion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_pubsub_topic#client_id GooglePubsubTopic#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The Azure event hub to ingest data from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_pubsub_topic#event_hub GooglePubsubTopic#event_hub}
	EventHub *string `field:"optional" json:"eventHub" yaml:"eventHub"`
	// The GCP service account to be used for Federated Identity authentication with Azure (via a 'AssumeRoleWithWebIdentity' call for the provided role).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_pubsub_topic#gcp_service_account GooglePubsubTopic#gcp_service_account}
	GcpServiceAccount *string `field:"optional" json:"gcpServiceAccount" yaml:"gcpServiceAccount"`
	// The Azure event hub namespace to ingest data from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_pubsub_topic#namespace GooglePubsubTopic#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The name of the resource group within an Azure subscription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_pubsub_topic#resource_group GooglePubsubTopic#resource_group}
	ResourceGroup *string `field:"optional" json:"resourceGroup" yaml:"resourceGroup"`
	// The Azure event hub subscription ID to use for ingestion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_pubsub_topic#subscription_id GooglePubsubTopic#subscription_id}
	SubscriptionId *string `field:"optional" json:"subscriptionId" yaml:"subscriptionId"`
	// The Azure event hub tenant ID to use for ingestion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_pubsub_topic#tenant_id GooglePubsubTopic#tenant_id}
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

