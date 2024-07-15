package googlevertexaifeaturestoreentitytype


type GoogleVertexAiFeaturestoreEntitytypeMonitoringConfig struct {
	// categorical_threshold_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_featurestore_entitytype#categorical_threshold_config GoogleVertexAiFeaturestoreEntitytype#categorical_threshold_config}
	CategoricalThresholdConfig *GoogleVertexAiFeaturestoreEntitytypeMonitoringConfigCategoricalThresholdConfig `field:"optional" json:"categoricalThresholdConfig" yaml:"categoricalThresholdConfig"`
	// import_features_analysis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_featurestore_entitytype#import_features_analysis GoogleVertexAiFeaturestoreEntitytype#import_features_analysis}
	ImportFeaturesAnalysis *GoogleVertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysis `field:"optional" json:"importFeaturesAnalysis" yaml:"importFeaturesAnalysis"`
	// numerical_threshold_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_featurestore_entitytype#numerical_threshold_config GoogleVertexAiFeaturestoreEntitytype#numerical_threshold_config}
	NumericalThresholdConfig *GoogleVertexAiFeaturestoreEntitytypeMonitoringConfigNumericalThresholdConfig `field:"optional" json:"numericalThresholdConfig" yaml:"numericalThresholdConfig"`
	// snapshot_analysis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_featurestore_entitytype#snapshot_analysis GoogleVertexAiFeaturestoreEntitytype#snapshot_analysis}
	SnapshotAnalysis *GoogleVertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysis `field:"optional" json:"snapshotAnalysis" yaml:"snapshotAnalysis"`
}

