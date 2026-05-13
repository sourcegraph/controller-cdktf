package discoveryenginecontrol


type DiscoveryEngineControlBoostAction struct {
	// The data store to boost.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#data_store DiscoveryEngineControl#data_store}
	DataStore *string `field:"required" json:"dataStore" yaml:"dataStore"`
	// The filter to apply to the search results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#filter DiscoveryEngineControl#filter}
	Filter *string `field:"required" json:"filter" yaml:"filter"`
	// The fixed boost value to apply to the search results.
	//
	// Positive values will increase the relevance of the results, while negative values will decrease the relevance. The value must be between -100 and 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#fixed_boost DiscoveryEngineControl#fixed_boost}
	FixedBoost *float64 `field:"optional" json:"fixedBoost" yaml:"fixedBoost"`
	// interpolation_boost_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#interpolation_boost_spec DiscoveryEngineControl#interpolation_boost_spec}
	InterpolationBoostSpec *DiscoveryEngineControlBoostActionInterpolationBoostSpec `field:"optional" json:"interpolationBoostSpec" yaml:"interpolationBoostSpec"`
}

