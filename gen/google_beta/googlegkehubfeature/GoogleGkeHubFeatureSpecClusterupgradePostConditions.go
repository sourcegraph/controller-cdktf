package googlegkehubfeature


type GoogleGkeHubFeatureSpecClusterupgradePostConditions struct {
	// Amount of time to "soak" after a rollout has been finished before marking it COMPLETE. Cannot exceed 30 days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#soaking GoogleGkeHubFeature#soaking}
	Soaking *string `field:"required" json:"soaking" yaml:"soaking"`
}

