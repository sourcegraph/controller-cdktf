package googlegkehubfeaturemembership


type GoogleGkeHubFeatureMembershipConfigmanagementConfigSync struct {
	// git block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_gke_hub_feature_membership#git GoogleGkeHubFeatureMembership#git}
	Git *GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncGit `field:"optional" json:"git" yaml:"git"`
	// Set to true to enable the Config Sync admission webhook to prevent drifts.
	//
	// If set to `false`, disables the Config Sync admission webhook and does not prevent drifts.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_gke_hub_feature_membership#prevent_drift GoogleGkeHubFeatureMembership#prevent_drift}
	PreventDrift interface{} `field:"optional" json:"preventDrift" yaml:"preventDrift"`
	// Specifies whether the Config Sync Repo is in "hierarchical" or "unstructured" mode.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_gke_hub_feature_membership#source_format GoogleGkeHubFeatureMembership#source_format}
	SourceFormat *string `field:"optional" json:"sourceFormat" yaml:"sourceFormat"`
}

