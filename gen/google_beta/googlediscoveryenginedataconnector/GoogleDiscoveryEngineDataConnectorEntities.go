package googlediscoveryenginedataconnector


type GoogleDiscoveryEngineDataConnectorEntities struct {
	// The name of the entity.
	//
	// Supported values by data source:
	// * Salesforce: 'Lead', 'Opportunity', 'Contact', 'Account', 'Case', 'Contract', 'Campaign'
	// * Jira: project, issue, attachment, comment, worklog
	// * Confluence: 'Content', 'Space'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_connector#entity_name GoogleDiscoveryEngineDataConnector#entity_name}
	EntityName *string `field:"optional" json:"entityName" yaml:"entityName"`
	// Attributes for indexing.
	//
	// Key: Field name.
	// Value: The key property to map a field to, such as 'title', and
	// 'description'. Supported key properties:
	// * 'title': The title for data record. This would be displayed on search
	//   results.
	// * 'description': The description for data record. This would be displayed
	//   on search results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_connector#key_property_mappings GoogleDiscoveryEngineDataConnector#key_property_mappings}
	KeyPropertyMappings *map[string]*string `field:"optional" json:"keyPropertyMappings" yaml:"keyPropertyMappings"`
	// The parameters for the entity to facilitate data ingestion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_connector#params GoogleDiscoveryEngineDataConnector#params}
	Params *string `field:"optional" json:"params" yaml:"params"`
}

