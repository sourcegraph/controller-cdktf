package googlecestool


type GoogleCesToolDataStoreToolModalityConfigsRewriterConfig struct {
	// model_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_tool#model_settings GoogleCesTool#model_settings}
	ModelSettings *GoogleCesToolDataStoreToolModalityConfigsRewriterConfigModelSettings `field:"required" json:"modelSettings" yaml:"modelSettings"`
	// Whether the rewriter is disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_tool#disabled GoogleCesTool#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// The prompt definition. If not set, default prompt will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_tool#prompt GoogleCesTool#prompt}
	Prompt *string `field:"optional" json:"prompt" yaml:"prompt"`
}

