package googlegkehubfeature


type GoogleGkeHubFeatureSpecFleetobservabilityLoggingConfig struct {
	// default_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#default_config GoogleGkeHubFeature#default_config}
	DefaultConfig *GoogleGkeHubFeatureSpecFleetobservabilityLoggingConfigDefaultConfig `field:"optional" json:"defaultConfig" yaml:"defaultConfig"`
	// fleet_scope_logs_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#fleet_scope_logs_config GoogleGkeHubFeature#fleet_scope_logs_config}
	FleetScopeLogsConfig *GoogleGkeHubFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig `field:"optional" json:"fleetScopeLogsConfig" yaml:"fleetScopeLogsConfig"`
}

