package googlegkehubfeature


type GoogleGkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesRequests struct {
	// CPU requirement expressed in Kubernetes resource units.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#cpu GoogleGkeHubFeature#cpu}
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// Memory requirement expressed in Kubernetes resource units.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#memory GoogleGkeHubFeature#memory}
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
}

