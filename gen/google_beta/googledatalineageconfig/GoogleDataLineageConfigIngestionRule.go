package googledatalineageconfig


type GoogleDataLineageConfigIngestionRule struct {
	// integration_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_data_lineage_config#integration_selector GoogleDataLineageConfig#integration_selector}
	IntegrationSelector *GoogleDataLineageConfigIngestionRuleIntegrationSelector `field:"required" json:"integrationSelector" yaml:"integrationSelector"`
	// lineage_enablement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_data_lineage_config#lineage_enablement GoogleDataLineageConfig#lineage_enablement}
	LineageEnablement *GoogleDataLineageConfigIngestionRuleLineageEnablement `field:"required" json:"lineageEnablement" yaml:"lineageEnablement"`
}

