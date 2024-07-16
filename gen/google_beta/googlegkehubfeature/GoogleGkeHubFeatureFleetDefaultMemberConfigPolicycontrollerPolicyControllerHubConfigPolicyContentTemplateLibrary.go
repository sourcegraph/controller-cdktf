package googlegkehubfeature


type GoogleGkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary struct {
	// Configures the manner in which the template library is installed on the cluster. Possible values: ["INSTALATION_UNSPECIFIED", "NOT_INSTALLED", "ALL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#installation GoogleGkeHubFeature#installation}
	Installation *string `field:"optional" json:"installation" yaml:"installation"`
}

