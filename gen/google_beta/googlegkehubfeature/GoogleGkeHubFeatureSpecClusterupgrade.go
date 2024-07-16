package googlegkehubfeature


type GoogleGkeHubFeatureSpecClusterupgrade struct {
	// Specified if other fleet should be considered as a source of upgrades.
	//
	// Currently, at most one upstream fleet is allowed. The fleet name should be either fleet project number or id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#upstream_fleets GoogleGkeHubFeature#upstream_fleets}
	UpstreamFleets *[]*string `field:"required" json:"upstreamFleets" yaml:"upstreamFleets"`
	// gke_upgrade_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#gke_upgrade_overrides GoogleGkeHubFeature#gke_upgrade_overrides}
	GkeUpgradeOverrides interface{} `field:"optional" json:"gkeUpgradeOverrides" yaml:"gkeUpgradeOverrides"`
	// post_conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#post_conditions GoogleGkeHubFeature#post_conditions}
	PostConditions *GoogleGkeHubFeatureSpecClusterupgradePostConditions `field:"optional" json:"postConditions" yaml:"postConditions"`
}

