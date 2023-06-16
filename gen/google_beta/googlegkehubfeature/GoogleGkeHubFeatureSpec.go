package googlegkehubfeature


type GoogleGkeHubFeatureSpec struct {
	// fleetobservability block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_gke_hub_feature#fleetobservability GoogleGkeHubFeature#fleetobservability}
	Fleetobservability *GoogleGkeHubFeatureSpecFleetobservability `field:"optional" json:"fleetobservability" yaml:"fleetobservability"`
	// multiclusteringress block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_gke_hub_feature#multiclusteringress GoogleGkeHubFeature#multiclusteringress}
	Multiclusteringress *GoogleGkeHubFeatureSpecMulticlusteringress `field:"optional" json:"multiclusteringress" yaml:"multiclusteringress"`
}

