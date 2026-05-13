package cestool


type CesToolDataStoreToolModalityConfigs struct {
	// The modality type. Possible values: TEXT AUDIO.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#modality_type CesTool#modality_type}
	ModalityType *string `field:"required" json:"modalityType" yaml:"modalityType"`
	// grounding_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#grounding_config CesTool#grounding_config}
	GroundingConfig *CesToolDataStoreToolModalityConfigsGroundingConfig `field:"optional" json:"groundingConfig" yaml:"groundingConfig"`
	// rewriter_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#rewriter_config CesTool#rewriter_config}
	RewriterConfig *CesToolDataStoreToolModalityConfigsRewriterConfig `field:"optional" json:"rewriterConfig" yaml:"rewriterConfig"`
	// summarization_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#summarization_config CesTool#summarization_config}
	SummarizationConfig *CesToolDataStoreToolModalityConfigsSummarizationConfig `field:"optional" json:"summarizationConfig" yaml:"summarizationConfig"`
}

