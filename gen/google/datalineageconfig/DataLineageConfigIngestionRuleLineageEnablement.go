package datalineageconfig


type DataLineageConfigIngestionRuleLineageEnablement struct {
	// Whether ingestion of lineage should be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/data_lineage_config#enabled DataLineageConfig#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

