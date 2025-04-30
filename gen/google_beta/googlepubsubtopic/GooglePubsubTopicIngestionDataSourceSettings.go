package googlepubsubtopic


type GooglePubsubTopicIngestionDataSourceSettings struct {
	// aws_kinesis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_pubsub_topic#aws_kinesis GooglePubsubTopic#aws_kinesis}
	AwsKinesis *GooglePubsubTopicIngestionDataSourceSettingsAwsKinesis `field:"optional" json:"awsKinesis" yaml:"awsKinesis"`
	// aws_msk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_pubsub_topic#aws_msk GooglePubsubTopic#aws_msk}
	AwsMsk *GooglePubsubTopicIngestionDataSourceSettingsAwsMsk `field:"optional" json:"awsMsk" yaml:"awsMsk"`
	// azure_event_hubs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_pubsub_topic#azure_event_hubs GooglePubsubTopic#azure_event_hubs}
	AzureEventHubs *GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubs `field:"optional" json:"azureEventHubs" yaml:"azureEventHubs"`
	// cloud_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_pubsub_topic#cloud_storage GooglePubsubTopic#cloud_storage}
	CloudStorage *GooglePubsubTopicIngestionDataSourceSettingsCloudStorage `field:"optional" json:"cloudStorage" yaml:"cloudStorage"`
	// confluent_cloud block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_pubsub_topic#confluent_cloud GooglePubsubTopic#confluent_cloud}
	ConfluentCloud *GooglePubsubTopicIngestionDataSourceSettingsConfluentCloud `field:"optional" json:"confluentCloud" yaml:"confluentCloud"`
	// platform_logs_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_pubsub_topic#platform_logs_settings GooglePubsubTopic#platform_logs_settings}
	PlatformLogsSettings *GooglePubsubTopicIngestionDataSourceSettingsPlatformLogsSettings `field:"optional" json:"platformLogsSettings" yaml:"platformLogsSettings"`
}

