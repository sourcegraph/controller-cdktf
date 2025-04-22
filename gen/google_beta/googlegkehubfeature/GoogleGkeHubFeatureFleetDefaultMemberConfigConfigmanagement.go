package googlegkehubfeature


type GoogleGkeHubFeatureFleetDefaultMemberConfigConfigmanagement struct {
	// config_sync block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gke_hub_feature#config_sync GoogleGkeHubFeature#config_sync}
	ConfigSync *GoogleGkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSync `field:"optional" json:"configSync" yaml:"configSync"`
	// Set this field to MANAGEMENT_AUTOMATIC to enable Config Sync auto-upgrades, and set this field to MANAGEMENT_MANUAL or MANAGEMENT_UNSPECIFIED to disable Config Sync auto-upgrades.
	//
	// Possible values: ["MANAGEMENT_UNSPECIFIED", "MANAGEMENT_AUTOMATIC", "MANAGEMENT_MANUAL"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gke_hub_feature#management GoogleGkeHubFeature#management}
	Management *string `field:"optional" json:"management" yaml:"management"`
	// Version of Config Sync installed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gke_hub_feature#version GoogleGkeHubFeature#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

