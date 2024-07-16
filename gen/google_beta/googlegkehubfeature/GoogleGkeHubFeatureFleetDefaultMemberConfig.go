package googlegkehubfeature


type GoogleGkeHubFeatureFleetDefaultMemberConfig struct {
	// configmanagement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#configmanagement GoogleGkeHubFeature#configmanagement}
	Configmanagement *GoogleGkeHubFeatureFleetDefaultMemberConfigConfigmanagement `field:"optional" json:"configmanagement" yaml:"configmanagement"`
	// mesh block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#mesh GoogleGkeHubFeature#mesh}
	Mesh *GoogleGkeHubFeatureFleetDefaultMemberConfigMesh `field:"optional" json:"mesh" yaml:"mesh"`
	// policycontroller block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#policycontroller GoogleGkeHubFeature#policycontroller}
	Policycontroller *GoogleGkeHubFeatureFleetDefaultMemberConfigPolicycontroller `field:"optional" json:"policycontroller" yaml:"policycontroller"`
}

