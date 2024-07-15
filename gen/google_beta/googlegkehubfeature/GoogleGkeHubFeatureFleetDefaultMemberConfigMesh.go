package googlegkehubfeature


type GoogleGkeHubFeatureFleetDefaultMemberConfigMesh struct {
	// Whether to automatically manage Service Mesh Possible values: ["MANAGEMENT_UNSPECIFIED", "MANAGEMENT_AUTOMATIC", "MANAGEMENT_MANUAL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#management GoogleGkeHubFeature#management}
	Management *string `field:"required" json:"management" yaml:"management"`
}

