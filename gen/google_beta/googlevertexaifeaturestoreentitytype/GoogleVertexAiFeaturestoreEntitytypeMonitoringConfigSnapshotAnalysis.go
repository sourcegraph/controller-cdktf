package googlevertexaifeaturestoreentitytype


type GoogleVertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysis struct {
	// The monitoring schedule for snapshot analysis.
	//
	// For EntityType-level config: unset / disabled = true indicates disabled by default for Features under it; otherwise by default enable snapshot analysis monitoring with monitoringInterval for Features under it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_featurestore_entitytype#disabled GoogleVertexAiFeaturestoreEntitytype#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Configuration of the snapshot analysis based monitoring pipeline running interval. The value is rolled up to full day.
	//
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_featurestore_entitytype#monitoring_interval GoogleVertexAiFeaturestoreEntitytype#monitoring_interval}
	MonitoringInterval *string `field:"optional" json:"monitoringInterval" yaml:"monitoringInterval"`
	// Configuration of the snapshot analysis based monitoring pipeline running interval.
	//
	// The value indicates number of days. The default value is 1.
	// If both FeaturestoreMonitoringConfig.SnapshotAnalysis.monitoring_interval_days and [FeaturestoreMonitoringConfig.SnapshotAnalysis.monitoring_interval][] are set when creating/updating EntityTypes/Features, FeaturestoreMonitoringConfig.SnapshotAnalysis.monitoring_interval_days will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_featurestore_entitytype#monitoring_interval_days GoogleVertexAiFeaturestoreEntitytype#monitoring_interval_days}
	MonitoringIntervalDays *float64 `field:"optional" json:"monitoringIntervalDays" yaml:"monitoringIntervalDays"`
	// Customized export features time window for snapshot analysis.
	//
	// Unit is one day. The default value is 21 days. Minimum value is 1 day. Maximum value is 4000 days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_featurestore_entitytype#staleness_days GoogleVertexAiFeaturestoreEntitytype#staleness_days}
	StalenessDays *float64 `field:"optional" json:"stalenessDays" yaml:"stalenessDays"`
}

