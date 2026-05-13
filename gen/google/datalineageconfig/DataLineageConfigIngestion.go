package datalineageconfig


type DataLineageConfigIngestion struct {
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/data_lineage_config#rule DataLineageConfig#rule}
	Rule interface{} `field:"required" json:"rule" yaml:"rule"`
}

