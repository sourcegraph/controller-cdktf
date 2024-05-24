package googlegkehubfeature


type GoogleGkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSync struct {
	// git block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.30.0/docs/resources/google_gke_hub_feature#git GoogleGkeHubFeature#git}
	Git *GoogleGkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGit `field:"optional" json:"git" yaml:"git"`
	// oci block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.30.0/docs/resources/google_gke_hub_feature#oci GoogleGkeHubFeature#oci}
	Oci *GoogleGkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOci `field:"optional" json:"oci" yaml:"oci"`
	// Specifies whether the Config Sync Repo is in hierarchical or unstructured mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.30.0/docs/resources/google_gke_hub_feature#source_format GoogleGkeHubFeature#source_format}
	SourceFormat *string `field:"optional" json:"sourceFormat" yaml:"sourceFormat"`
}

