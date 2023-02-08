package googlegkehubfeature


type GoogleGkeHubFeatureSpec struct {
	// multiclusteringress block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_gke_hub_feature#multiclusteringress GoogleGkeHubFeature#multiclusteringress}
	Multiclusteringress *GoogleGkeHubFeatureSpecMulticlusteringress `field:"optional" json:"multiclusteringress" yaml:"multiclusteringress"`
}

