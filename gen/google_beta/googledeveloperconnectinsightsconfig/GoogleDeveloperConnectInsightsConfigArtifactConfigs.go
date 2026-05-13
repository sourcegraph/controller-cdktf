package googledeveloperconnectinsightsconfig


type GoogleDeveloperConnectInsightsConfigArtifactConfigs struct {
	// google_artifact_analysis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_developer_connect_insights_config#google_artifact_analysis GoogleDeveloperConnectInsightsConfig#google_artifact_analysis}
	GoogleArtifactAnalysis *GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactAnalysis `field:"optional" json:"googleArtifactAnalysis" yaml:"googleArtifactAnalysis"`
	// google_artifact_registry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_developer_connect_insights_config#google_artifact_registry GoogleDeveloperConnectInsightsConfig#google_artifact_registry}
	GoogleArtifactRegistry *GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistry `field:"optional" json:"googleArtifactRegistry" yaml:"googleArtifactRegistry"`
	// The URI of the artifact that is deployed.
	//
	// e.g. 'us-docker.pkg.dev/my-project/my-repo/image'.
	// The URI does not include the tag / digest because it captures a lineage of
	// artifacts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_developer_connect_insights_config#uri GoogleDeveloperConnectInsightsConfig#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

