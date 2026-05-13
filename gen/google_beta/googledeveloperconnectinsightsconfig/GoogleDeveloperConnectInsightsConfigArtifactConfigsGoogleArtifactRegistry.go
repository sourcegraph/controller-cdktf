package googledeveloperconnectinsightsconfig


type GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistry struct {
	// The name of the artifact registry package.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_developer_connect_insights_config#artifact_registry_package GoogleDeveloperConnectInsightsConfig#artifact_registry_package}
	ArtifactRegistryPackage *string `field:"required" json:"artifactRegistryPackage" yaml:"artifactRegistryPackage"`
	// The host project of Artifact Registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_developer_connect_insights_config#project_id GoogleDeveloperConnectInsightsConfig#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

