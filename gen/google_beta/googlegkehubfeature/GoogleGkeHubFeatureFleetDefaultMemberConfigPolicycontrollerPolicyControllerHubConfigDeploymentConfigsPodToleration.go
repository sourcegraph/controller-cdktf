package googlegkehubfeature


type GoogleGkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodToleration struct {
	// Matches a taint effect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#effect GoogleGkeHubFeature#effect}
	Effect *string `field:"optional" json:"effect" yaml:"effect"`
	// Matches a taint key (not necessarily unique).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#key GoogleGkeHubFeature#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Matches a taint operator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#operator GoogleGkeHubFeature#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// Matches a taint value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#value GoogleGkeHubFeature#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

