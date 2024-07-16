package googlegkehubfeature


type GoogleGkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigPolicyContentBundles struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#bundle GoogleGkeHubFeature#bundle}.
	Bundle *string `field:"required" json:"bundle" yaml:"bundle"`
	// The set of namespaces to be exempted from the bundle.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#exempted_namespaces GoogleGkeHubFeature#exempted_namespaces}
	ExemptedNamespaces *[]*string `field:"optional" json:"exemptedNamespaces" yaml:"exemptedNamespaces"`
}

