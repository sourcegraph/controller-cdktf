package googlediscoveryenginecontrol


type GoogleDiscoveryEngineControlSynonymsAction struct {
	// The synonyms to apply to the search results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_control#synonyms GoogleDiscoveryEngineControl#synonyms}
	Synonyms *[]*string `field:"optional" json:"synonyms" yaml:"synonyms"`
}

