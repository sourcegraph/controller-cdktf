package googlediscoveryenginecontrol


type GoogleDiscoveryEngineControlConditions struct {
	// active_time_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_control#active_time_range GoogleDiscoveryEngineControl#active_time_range}
	ActiveTimeRange interface{} `field:"optional" json:"activeTimeRange" yaml:"activeTimeRange"`
	// The regular expression that the query must match for this condition to be met.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_control#query_regex GoogleDiscoveryEngineControl#query_regex}
	QueryRegex *string `field:"optional" json:"queryRegex" yaml:"queryRegex"`
	// query_terms block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_control#query_terms GoogleDiscoveryEngineControl#query_terms}
	QueryTerms interface{} `field:"optional" json:"queryTerms" yaml:"queryTerms"`
}

