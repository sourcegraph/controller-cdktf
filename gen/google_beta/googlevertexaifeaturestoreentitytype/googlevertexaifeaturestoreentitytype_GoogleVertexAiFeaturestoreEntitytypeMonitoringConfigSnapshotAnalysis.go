package googlevertexaifeaturestoreentitytype


type GoogleVertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysis struct {
	// The monitoring schedule for snapshot analysis.
	//
	// For EntityType-level config: unset / disabled = true indicates disabled by default for Features under it; otherwise by default enable snapshot analysis monitoring with monitoringInterval for Features under it.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_vertex_ai_featurestore_entitytype#disabled GoogleVertexAiFeaturestoreEntitytype#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Configuration of the snapshot analysis based monitoring pipeline running interval. The value is rolled up to full day.
	//
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_vertex_ai_featurestore_entitytype#monitoring_interval GoogleVertexAiFeaturestoreEntitytype#monitoring_interval}
	MonitoringInterval *string `field:"optional" json:"monitoringInterval" yaml:"monitoringInterval"`
}

