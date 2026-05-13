package cestool


type CesToolDataStoreToolModalityConfigsRewriterConfig struct {
	// model_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#model_settings CesTool#model_settings}
	ModelSettings *CesToolDataStoreToolModalityConfigsRewriterConfigModelSettings `field:"required" json:"modelSettings" yaml:"modelSettings"`
	// Whether the rewriter is disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#disabled CesTool#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// The prompt definition. If not set, default prompt will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#prompt CesTool#prompt}
	Prompt *string `field:"optional" json:"prompt" yaml:"prompt"`
}

