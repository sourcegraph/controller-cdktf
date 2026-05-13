package googlecestool


type GoogleCesToolDataStoreToolModalityConfigs struct {
	// The modality type. Possible values: TEXT AUDIO.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_tool#modality_type GoogleCesTool#modality_type}
	ModalityType *string `field:"required" json:"modalityType" yaml:"modalityType"`
	// grounding_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_tool#grounding_config GoogleCesTool#grounding_config}
	GroundingConfig *GoogleCesToolDataStoreToolModalityConfigsGroundingConfig `field:"optional" json:"groundingConfig" yaml:"groundingConfig"`
	// rewriter_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_tool#rewriter_config GoogleCesTool#rewriter_config}
	RewriterConfig *GoogleCesToolDataStoreToolModalityConfigsRewriterConfig `field:"optional" json:"rewriterConfig" yaml:"rewriterConfig"`
	// summarization_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_tool#summarization_config GoogleCesTool#summarization_config}
	SummarizationConfig *GoogleCesToolDataStoreToolModalityConfigsSummarizationConfig `field:"optional" json:"summarizationConfig" yaml:"summarizationConfig"`
}

