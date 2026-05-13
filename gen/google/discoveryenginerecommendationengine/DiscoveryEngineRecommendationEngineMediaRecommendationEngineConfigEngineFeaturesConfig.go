package discoveryenginerecommendationengine


type DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfig struct {
	// most_popular_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_recommendation_engine#most_popular_config DiscoveryEngineRecommendationEngine#most_popular_config}
	MostPopularConfig *DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigMostPopularConfig `field:"optional" json:"mostPopularConfig" yaml:"mostPopularConfig"`
	// recommended_for_you_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_recommendation_engine#recommended_for_you_config DiscoveryEngineRecommendationEngine#recommended_for_you_config}
	RecommendedForYouConfig *DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigRecommendedForYouConfig `field:"optional" json:"recommendedForYouConfig" yaml:"recommendedForYouConfig"`
}

