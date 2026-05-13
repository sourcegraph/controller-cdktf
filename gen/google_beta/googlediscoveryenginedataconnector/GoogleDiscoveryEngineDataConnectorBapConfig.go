package googlediscoveryenginedataconnector


type GoogleDiscoveryEngineDataConnectorBapConfig struct {
	// The list of enabled actions for this connector. Supported values include: 'create_issue', 'update_issue', 'change_issue_status', 'create_comment', 'update_comment', 'upload_attachment'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_connector#enabled_actions GoogleDiscoveryEngineDataConnector#enabled_actions}
	EnabledActions *[]*string `field:"optional" json:"enabledActions" yaml:"enabledActions"`
	// The connector modes supported by the BAP configuration. The possible values include: 'ACTIONS'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_connector#supported_connector_modes GoogleDiscoveryEngineDataConnector#supported_connector_modes}
	SupportedConnectorModes *[]*string `field:"optional" json:"supportedConnectorModes" yaml:"supportedConnectorModes"`
}

