package googlegkehubfeaturemembership


type GoogleGkeHubFeatureMembershipPolicycontroller struct {
	// policy_controller_hub_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#policy_controller_hub_config GoogleGkeHubFeatureMembership#policy_controller_hub_config}
	PolicyControllerHubConfig *GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfig `field:"required" json:"policyControllerHubConfig" yaml:"policyControllerHubConfig"`
	// Optional. Version of Policy Controller to install. Defaults to the latest version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#version GoogleGkeHubFeatureMembership#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

