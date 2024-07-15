package googlegkehubfeature


type GoogleGkeHubFeatureFleetDefaultMemberConfigConfigmanagement struct {
	// config_sync block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#config_sync GoogleGkeHubFeature#config_sync}
	ConfigSync *GoogleGkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSync `field:"optional" json:"configSync" yaml:"configSync"`
	// Version of ACM installed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#version GoogleGkeHubFeature#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

