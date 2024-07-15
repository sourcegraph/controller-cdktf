package googlegkehubfeaturemembership


type GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentBundles struct {
	// The name for the key in the map for which this object is mapped to in the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#bundle_name GoogleGkeHubFeatureMembership#bundle_name}
	BundleName *string `field:"required" json:"bundleName" yaml:"bundleName"`
	// The set of namespaces to be exempted from the bundle.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#exempted_namespaces GoogleGkeHubFeatureMembership#exempted_namespaces}
	ExemptedNamespaces *[]*string `field:"optional" json:"exemptedNamespaces" yaml:"exemptedNamespaces"`
}

