package googledatalineageconfig


type GoogleDataLineageConfigIngestionRuleIntegrationSelector struct {
	// Integration to which the rule applies. Possible values: ["DATAPROC", "LOOKER_CORE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_data_lineage_config#integration GoogleDataLineageConfig#integration}
	Integration *string `field:"required" json:"integration" yaml:"integration"`
}

