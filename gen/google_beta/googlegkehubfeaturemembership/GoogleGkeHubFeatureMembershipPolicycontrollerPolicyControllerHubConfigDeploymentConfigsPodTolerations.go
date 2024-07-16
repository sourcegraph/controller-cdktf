package googlegkehubfeaturemembership


type GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodTolerations struct {
	// Matches a taint effect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#effect GoogleGkeHubFeatureMembership#effect}
	Effect *string `field:"optional" json:"effect" yaml:"effect"`
	// Matches a taint key (not necessarily unique).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#key GoogleGkeHubFeatureMembership#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Matches a taint operator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#operator GoogleGkeHubFeatureMembership#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// Matches a taint value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#value GoogleGkeHubFeatureMembership#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

