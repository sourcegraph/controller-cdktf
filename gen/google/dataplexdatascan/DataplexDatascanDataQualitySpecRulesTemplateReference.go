package dataplexdatascan


type DataplexDatascanDataQualitySpecRulesTemplateReference struct {
	// The resource name of the template entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dataplex_datascan#name DataplexDatascan#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dataplex_datascan#values DataplexDatascan#values}
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

