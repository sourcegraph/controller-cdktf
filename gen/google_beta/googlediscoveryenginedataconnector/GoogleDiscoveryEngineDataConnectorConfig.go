package googlediscoveryenginedataconnector

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDiscoveryEngineDataConnectorConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The display name of the Collection.
	//
	// Should be human readable, used to display collections in the Console
	// Dashboard. UTF-8 encoded string with limit of 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_connector#collection_display_name GoogleDiscoveryEngineDataConnector#collection_display_name}
	CollectionDisplayName *string `field:"required" json:"collectionDisplayName" yaml:"collectionDisplayName"`
	// The ID to use for the Collection, which will become the final component of the Collection's resource name.
	//
	// A new Collection is created as
	// part of the DataConnector setup. DataConnector is a singleton
	// resource under Collection, managing all DataStores of the Collection.
	// This field must conform to [RFC-1034](https://tools.ietf.org/html/rfc1034)
	// standard with a length limit of 63 characters. Otherwise, an
	// INVALID_ARGUMENT error is returned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_connector#collection_id GoogleDiscoveryEngineDataConnector#collection_id}
	CollectionId *string `field:"required" json:"collectionId" yaml:"collectionId"`
	// The identifier for the data source.
	//
	// This is a partial list of supported connectors. Please refer to the
	// [documentation](https://docs.cloud.google.com/gemini/enterprise/docs/connectors/introduction-to-connectors-and-data-stores)
	// for the full list of connectors.
	//
	// Supported first-party connectors include:
	//
	// *   'bigquery'
	// *   'gcp_fhir'
	// *   'google_mail'
	// *   'google_drive'
	// *   'google_calendar'
	// *   'google_chat'
	//
	// Supported third-party connectors include:
	// Generally available (GA) connectors:
	//
	// *   'onedrive'
	// *   'outlook'
	// *   'confluence'
	// *   'jira'
	// *   'servicenow'
	// *   'sharepoint'
	//
	// Preview connectors:
	//
	// *   'asana'
	// *   'azure_active_directory'
	// *   'box'
	// *   'canva'
	// *   'confluence_server'
	// *   'custom_connector'
	// *   'docusign'
	// *   'dropbox'
	// *   'dynamics365'
	// *   'github'
	// *   'gitlab'
	// *   'hubspot'
	// *   'jira_server'
	// *   'linear'
	// *   'native_cloud_identity'
	// *   'notion'
	// *   'okta'
	// *   'pagerduty'
	// *   'peoplesoft'
	// *   'salesforce'
	// *   'shopify'
	// *   'slack'
	// *   'snowflake'
	// *   'teams'
	// *   'trello'
	// *   'workday'
	// *   'zendesk'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_connector#data_source GoogleDiscoveryEngineDataConnector#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// The geographic location where the data store should reside. The value can only be one of "global", "us" and "eu".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_connector#location GoogleDiscoveryEngineDataConnector#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The refresh interval for data sync.
	//
	// If duration is set to 0, the data will
	// be synced in real time. The streaming feature is not supported yet. The
	// minimum is 30 minutes and maximum is 7 days. When the refresh interval is
	// set to the same value as the incremental refresh interval, incremental
	// sync will be disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_connector#refresh_interval GoogleDiscoveryEngineDataConnector#refresh_interval}
	RefreshInterval *string `field:"required" json:"refreshInterval" yaml:"refreshInterval"`
	// action_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_connector#action_config GoogleDiscoveryEngineDataConnector#action_config}
	ActionConfig *GoogleDiscoveryEngineDataConnectorActionConfig `field:"optional" json:"actionConfig" yaml:"actionConfig"`
	// Indicates whether full syncs are paused for this connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_connector#auto_run_disabled GoogleDiscoveryEngineDataConnector#auto_run_disabled}
	AutoRunDisabled interface{} `field:"optional" json:"autoRunDisabled" yaml:"autoRunDisabled"`
	// bap_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_connector#bap_config GoogleDiscoveryEngineDataConnector#bap_config}
	BapConfig *GoogleDiscoveryEngineDataConnectorBapConfig `field:"optional" json:"bapConfig" yaml:"bapConfig"`
	// The modes enabled for this connector. The possible value can be: 'DATA_INGESTION', 'ACTIONS', 'FEDERATED' 'EUA', 'FEDERATED_AND_EUA'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_connector#connector_modes GoogleDiscoveryEngineDataConnector#connector_modes}
	ConnectorModes *[]*string `field:"optional" json:"connectorModes" yaml:"connectorModes"`
	// The version of the data source. For example, '3' for Jira v3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_connector#data_source_version GoogleDiscoveryEngineDataConnector#data_source_version}
	DataSourceVersion *float64 `field:"optional" json:"dataSourceVersion" yaml:"dataSourceVersion"`
	// destination_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_connector#destination_configs GoogleDiscoveryEngineDataConnector#destination_configs}
	DestinationConfigs interface{} `field:"optional" json:"destinationConfigs" yaml:"destinationConfigs"`
	// entities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_connector#entities GoogleDiscoveryEngineDataConnector#entities}
	Entities interface{} `field:"optional" json:"entities" yaml:"entities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_connector#id GoogleDiscoveryEngineDataConnector#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The refresh interval specifically for incremental data syncs.
	//
	// If unset,
	// incremental syncs will use the default from env, set to 3hrs.
	// The minimum is 30 minutes and maximum is 7 days. Applicable to only 3P
	// connectors. When the refresh interval is
	// set to the same value as the incremental refresh interval, incremental
	// sync will be disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_connector#incremental_refresh_interval GoogleDiscoveryEngineDataConnector#incremental_refresh_interval}
	IncrementalRefreshInterval *string `field:"optional" json:"incrementalRefreshInterval" yaml:"incrementalRefreshInterval"`
	// Indicates whether incremental syncs are paused for this connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_connector#incremental_sync_disabled GoogleDiscoveryEngineDataConnector#incremental_sync_disabled}
	IncrementalSyncDisabled interface{} `field:"optional" json:"incrementalSyncDisabled" yaml:"incrementalSyncDisabled"`
	// Params needed to access the source in the format of json string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_connector#json_params GoogleDiscoveryEngineDataConnector#json_params}
	JsonParams *string `field:"optional" json:"jsonParams" yaml:"jsonParams"`
	// The KMS key to be used to protect the DataStores managed by this connector.
	//
	// Must be set for requests that need to comply with CMEK Org Policy
	// protections.
	// If this field is set and processed successfully, the DataStores created by
	// this connector will be protected by the KMS key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_connector#kms_key_name GoogleDiscoveryEngineDataConnector#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
	// Params needed to access the source in the format of String-to-String (Key, Value) pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_connector#params GoogleDiscoveryEngineDataConnector#params}
	Params *map[string]*string `field:"optional" json:"params" yaml:"params"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_connector#project GoogleDiscoveryEngineDataConnector#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Whether customer has enabled static IP addresses for this connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_connector#static_ip_enabled GoogleDiscoveryEngineDataConnector#static_ip_enabled}
	StaticIpEnabled interface{} `field:"optional" json:"staticIpEnabled" yaml:"staticIpEnabled"`
	// The data synchronization mode supported by the data connector. The possible value can be: 'PERIODIC', 'STREAMING'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_connector#sync_mode GoogleDiscoveryEngineDataConnector#sync_mode}
	SyncMode *string `field:"optional" json:"syncMode" yaml:"syncMode"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_connector#timeouts GoogleDiscoveryEngineDataConnector#timeouts}
	Timeouts *GoogleDiscoveryEngineDataConnectorTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

