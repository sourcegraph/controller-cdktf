package googlediscoveryenginecontrol


type GoogleDiscoveryEngineControlConditionsQueryTerms struct {
	// If true, the query term must be an exact match. Otherwise, the query term can be a partial match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_control#full_match GoogleDiscoveryEngineControl#full_match}
	FullMatch interface{} `field:"optional" json:"fullMatch" yaml:"fullMatch"`
	// The value of the query term.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_control#value GoogleDiscoveryEngineControl#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

