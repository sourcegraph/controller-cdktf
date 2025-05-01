package googlegkehubfeaturemembership


type GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainers struct {
	// The name of the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gke_hub_feature_membership#container_name GoogleGkeHubFeatureMembership#container_name}
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// The CPU limit of the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gke_hub_feature_membership#cpu_limit GoogleGkeHubFeatureMembership#cpu_limit}
	CpuLimit *string `field:"optional" json:"cpuLimit" yaml:"cpuLimit"`
	// The CPU request of the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gke_hub_feature_membership#cpu_request GoogleGkeHubFeatureMembership#cpu_request}
	CpuRequest *string `field:"optional" json:"cpuRequest" yaml:"cpuRequest"`
	// The memory limit of the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gke_hub_feature_membership#memory_limit GoogleGkeHubFeatureMembership#memory_limit}
	MemoryLimit *string `field:"optional" json:"memoryLimit" yaml:"memoryLimit"`
	// The memory request of the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gke_hub_feature_membership#memory_request GoogleGkeHubFeatureMembership#memory_request}
	MemoryRequest *string `field:"optional" json:"memoryRequest" yaml:"memoryRequest"`
}

