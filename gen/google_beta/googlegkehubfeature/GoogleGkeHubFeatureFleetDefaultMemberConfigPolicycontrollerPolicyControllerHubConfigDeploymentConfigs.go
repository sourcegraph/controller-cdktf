package googlegkehubfeature


type GoogleGkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigDeploymentConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#component GoogleGkeHubFeature#component}.
	Component *string `field:"required" json:"component" yaml:"component"`
	// container_resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#container_resources GoogleGkeHubFeature#container_resources}
	ContainerResources *GoogleGkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResources `field:"optional" json:"containerResources" yaml:"containerResources"`
	// Pod affinity configuration. Possible values: ["AFFINITY_UNSPECIFIED", "NO_AFFINITY", "ANTI_AFFINITY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#pod_affinity GoogleGkeHubFeature#pod_affinity}
	PodAffinity *string `field:"optional" json:"podAffinity" yaml:"podAffinity"`
	// pod_toleration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#pod_toleration GoogleGkeHubFeature#pod_toleration}
	PodToleration interface{} `field:"optional" json:"podToleration" yaml:"podToleration"`
	// Pod replica count.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#replica_count GoogleGkeHubFeature#replica_count}
	ReplicaCount *float64 `field:"optional" json:"replicaCount" yaml:"replicaCount"`
}

