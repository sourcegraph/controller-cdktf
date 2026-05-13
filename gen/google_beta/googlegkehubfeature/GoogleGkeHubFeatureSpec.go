package googlegkehubfeature


type GoogleGkeHubFeatureSpec struct {
	// clusterupgrade block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gke_hub_feature#clusterupgrade GoogleGkeHubFeature#clusterupgrade}
	Clusterupgrade *GoogleGkeHubFeatureSpecClusterupgrade `field:"optional" json:"clusterupgrade" yaml:"clusterupgrade"`
	// fleetobservability block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gke_hub_feature#fleetobservability GoogleGkeHubFeature#fleetobservability}
	Fleetobservability *GoogleGkeHubFeatureSpecFleetobservability `field:"optional" json:"fleetobservability" yaml:"fleetobservability"`
	// multiclusteringress block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gke_hub_feature#multiclusteringress GoogleGkeHubFeature#multiclusteringress}
	Multiclusteringress *GoogleGkeHubFeatureSpecMulticlusteringress `field:"optional" json:"multiclusteringress" yaml:"multiclusteringress"`
	// rbacrolebindingactuation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gke_hub_feature#rbacrolebindingactuation GoogleGkeHubFeature#rbacrolebindingactuation}
	Rbacrolebindingactuation *GoogleGkeHubFeatureSpecRbacrolebindingactuation `field:"optional" json:"rbacrolebindingactuation" yaml:"rbacrolebindingactuation"`
	// workloadidentity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gke_hub_feature#workloadidentity GoogleGkeHubFeature#workloadidentity}
	Workloadidentity *GoogleGkeHubFeatureSpecWorkloadidentity `field:"optional" json:"workloadidentity" yaml:"workloadidentity"`
}

