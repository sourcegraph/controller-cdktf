package googledatalineageconfig


type GoogleDataLineageConfigIngestion struct {
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_data_lineage_config#rule GoogleDataLineageConfig#rule}
	Rule interface{} `field:"required" json:"rule" yaml:"rule"`
}

