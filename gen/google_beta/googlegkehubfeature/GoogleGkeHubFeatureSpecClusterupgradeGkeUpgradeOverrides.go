package googlegkehubfeature


type GoogleGkeHubFeatureSpecClusterupgradeGkeUpgradeOverrides struct {
	// post_conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#post_conditions GoogleGkeHubFeature#post_conditions}
	PostConditions *GoogleGkeHubFeatureSpecClusterupgradeGkeUpgradeOverridesPostConditions `field:"required" json:"postConditions" yaml:"postConditions"`
	// upgrade block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#upgrade GoogleGkeHubFeature#upgrade}
	Upgrade *GoogleGkeHubFeatureSpecClusterupgradeGkeUpgradeOverridesUpgrade `field:"required" json:"upgrade" yaml:"upgrade"`
}

