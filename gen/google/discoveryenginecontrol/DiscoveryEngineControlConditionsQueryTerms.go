package discoveryenginecontrol


type DiscoveryEngineControlConditionsQueryTerms struct {
	// If true, the query term must be an exact match. Otherwise, the query term can be a partial match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#full_match DiscoveryEngineControl#full_match}
	FullMatch interface{} `field:"optional" json:"fullMatch" yaml:"fullMatch"`
	// The value of the query term.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#value DiscoveryEngineControl#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

