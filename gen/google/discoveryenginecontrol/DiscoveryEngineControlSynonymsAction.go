package discoveryenginecontrol


type DiscoveryEngineControlSynonymsAction struct {
	// The synonyms to apply to the search results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#synonyms DiscoveryEngineControl#synonyms}
	Synonyms *[]*string `field:"optional" json:"synonyms" yaml:"synonyms"`
}

