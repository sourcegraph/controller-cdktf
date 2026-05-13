package datalineageconfig


type DataLineageConfigIngestionRuleIntegrationSelector struct {
	// Integration to which the rule applies. Possible values: ["DATAPROC", "LOOKER_CORE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/data_lineage_config#integration DataLineageConfig#integration}
	Integration *string `field:"required" json:"integration" yaml:"integration"`
}

