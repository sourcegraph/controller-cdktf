package googlediscoveryenginerecommendationengine


type GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigMostPopularConfig struct {
	// The time window of which the engine is queried at training and prediction time.
	//
	// Positive integers only. The value translates to the
	// last X days of events. Currently required for the 'most-popular-items'
	// engine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_recommendation_engine#time_window_days GoogleDiscoveryEngineRecommendationEngine#time_window_days}
	TimeWindowDays *float64 `field:"optional" json:"timeWindowDays" yaml:"timeWindowDays"`
}

