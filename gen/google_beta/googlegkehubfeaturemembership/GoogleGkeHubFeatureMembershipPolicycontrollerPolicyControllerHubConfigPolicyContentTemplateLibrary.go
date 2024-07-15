package googlegkehubfeaturemembership


type GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary struct {
	// Configures the manner in which the template library is installed on the cluster. Possible values: INSTALLATION_UNSPECIFIED, NOT_INSTALLED, ALL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#installation GoogleGkeHubFeatureMembership#installation}
	Installation *string `field:"optional" json:"installation" yaml:"installation"`
}

