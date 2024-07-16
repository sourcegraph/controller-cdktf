package googlegkehubfeaturemembership


type GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigs struct {
	// The name for the key in the map for which this object is mapped to in the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#component_name GoogleGkeHubFeatureMembership#component_name}
	ComponentName *string `field:"required" json:"componentName" yaml:"componentName"`
	// container_resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#container_resources GoogleGkeHubFeatureMembership#container_resources}
	ContainerResources *GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResources `field:"optional" json:"containerResources" yaml:"containerResources"`
	// Pod affinity configuration. Possible values: AFFINITY_UNSPECIFIED, NO_AFFINITY, ANTI_AFFINITY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#pod_affinity GoogleGkeHubFeatureMembership#pod_affinity}
	PodAffinity *string `field:"optional" json:"podAffinity" yaml:"podAffinity"`
	// pod_tolerations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#pod_tolerations GoogleGkeHubFeatureMembership#pod_tolerations}
	PodTolerations interface{} `field:"optional" json:"podTolerations" yaml:"podTolerations"`
	// Pod replica count.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#replica_count GoogleGkeHubFeatureMembership#replica_count}
	ReplicaCount *float64 `field:"optional" json:"replicaCount" yaml:"replicaCount"`
}

