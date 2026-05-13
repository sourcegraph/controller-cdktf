package googlechroniclefeed


type GoogleChronicleFeedDetailsAzureEventHubSettings struct {
	// Event hub consumer group to read from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#consumer_group GoogleChronicleFeed#consumer_group}
	ConsumerGroup *string `field:"required" json:"consumerGroup" yaml:"consumerGroup"`
	// Event hub connection string for authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#event_hub_connection_string GoogleChronicleFeed#event_hub_connection_string}
	EventHubConnectionString *string `field:"required" json:"eventHubConnectionString" yaml:"eventHubConnectionString"`
	// Event hub to read from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#name GoogleChronicleFeed#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// SAS token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#azure_sas_token GoogleChronicleFeed#azure_sas_token}
	AzureSasToken *string `field:"optional" json:"azureSasToken" yaml:"azureSasToken"`
	// Blob store connection string for authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#azure_storage_connection_string GoogleChronicleFeed#azure_storage_connection_string}
	AzureStorageConnectionString *string `field:"optional" json:"azureStorageConnectionString" yaml:"azureStorageConnectionString"`
	// Blob storage container name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#azure_storage_container GoogleChronicleFeed#azure_storage_container}
	AzureStorageContainer *string `field:"optional" json:"azureStorageContainer" yaml:"azureStorageContainer"`
}

