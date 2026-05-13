package cestool


type CesToolDataStoreToolModalityConfigsGroundingConfig struct {
	// Whether grounding is disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#disabled CesTool#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// The groundedness threshold of the answer based on the retrieved sources.
	//
	// The value has a configurable range of [1, 5]. The level is used to
	// threshold the groundedness of the answer, meaning that all responses with
	// a groundedness score below the threshold will fall back to returning
	// relevant snippets only.
	// For example, a level of 3 means that the groundedness score must be
	// 3 or higher for the response to be returned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#grounding_level CesTool#grounding_level}
	GroundingLevel *float64 `field:"optional" json:"groundingLevel" yaml:"groundingLevel"`
}

