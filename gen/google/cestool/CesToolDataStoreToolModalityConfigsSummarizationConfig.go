package cestool


type CesToolDataStoreToolModalityConfigsSummarizationConfig struct {
	// Whether summarization is disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#disabled CesTool#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// model_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#model_settings CesTool#model_settings}
	ModelSettings *CesToolDataStoreToolModalityConfigsSummarizationConfigModelSettings `field:"optional" json:"modelSettings" yaml:"modelSettings"`
	// The prompt definition. If not set, default prompt will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#prompt CesTool#prompt}
	Prompt *string `field:"optional" json:"prompt" yaml:"prompt"`
}

