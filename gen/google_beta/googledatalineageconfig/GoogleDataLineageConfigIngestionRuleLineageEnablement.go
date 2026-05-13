package googledatalineageconfig


type GoogleDataLineageConfigIngestionRuleLineageEnablement struct {
	// Whether ingestion of lineage should be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_data_lineage_config#enabled GoogleDataLineageConfig#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

