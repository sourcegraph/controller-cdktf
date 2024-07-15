package googlegkehubfeature


type GoogleGkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSync struct {
	// Enables the installation of ConfigSync.
	//
	// If set to true, ConfigSync resources will be created and the other ConfigSync fields will be applied if exist. If set to false, all other ConfigSync fields will be ignored, ConfigSync resources will be deleted. If omitted, ConfigSync resources will be managed depends on the presence of the git or oci field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#enabled GoogleGkeHubFeature#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// git block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#git GoogleGkeHubFeature#git}
	Git *GoogleGkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGit `field:"optional" json:"git" yaml:"git"`
	// oci block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#oci GoogleGkeHubFeature#oci}
	Oci *GoogleGkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOci `field:"optional" json:"oci" yaml:"oci"`
	// Set to true to enable the Config Sync admission webhook to prevent drifts.
	//
	// If set to 'false', disables the Config Sync admission webhook and does not prevent drifts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#prevent_drift GoogleGkeHubFeature#prevent_drift}
	PreventDrift interface{} `field:"optional" json:"preventDrift" yaml:"preventDrift"`
	// Specifies whether the Config Sync Repo is in hierarchical or unstructured mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#source_format GoogleGkeHubFeature#source_format}
	SourceFormat *string `field:"optional" json:"sourceFormat" yaml:"sourceFormat"`
}

