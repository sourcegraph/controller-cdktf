package googlegkehubfeature


type GoogleGkeHubFeatureFleetDefaultMemberConfigPolicycontroller struct {
	// policy_controller_hub_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#policy_controller_hub_config GoogleGkeHubFeature#policy_controller_hub_config}
	PolicyControllerHubConfig *GoogleGkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfig `field:"required" json:"policyControllerHubConfig" yaml:"policyControllerHubConfig"`
	// Configures the version of Policy Controller.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#version GoogleGkeHubFeature#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

