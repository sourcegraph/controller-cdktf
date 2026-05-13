package googlediscoveryenginecontrol


type GoogleDiscoveryEngineControlBoostAction struct {
	// The data store to boost.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_control#data_store GoogleDiscoveryEngineControl#data_store}
	DataStore *string `field:"required" json:"dataStore" yaml:"dataStore"`
	// The filter to apply to the search results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_control#filter GoogleDiscoveryEngineControl#filter}
	Filter *string `field:"required" json:"filter" yaml:"filter"`
	// The fixed boost value to apply to the search results.
	//
	// Positive values will increase the relevance of the results, while negative values will decrease the relevance. The value must be between -100 and 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_control#fixed_boost GoogleDiscoveryEngineControl#fixed_boost}
	FixedBoost *float64 `field:"optional" json:"fixedBoost" yaml:"fixedBoost"`
	// interpolation_boost_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_control#interpolation_boost_spec GoogleDiscoveryEngineControl#interpolation_boost_spec}
	InterpolationBoostSpec *GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpec `field:"optional" json:"interpolationBoostSpec" yaml:"interpolationBoostSpec"`
}

