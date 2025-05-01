package googlegkehubfeaturemembership


type GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverrides struct {
	// containers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gke_hub_feature_membership#containers GoogleGkeHubFeatureMembership#containers}
	Containers interface{} `field:"optional" json:"containers" yaml:"containers"`
	// The name of the Deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gke_hub_feature_membership#deployment_name GoogleGkeHubFeatureMembership#deployment_name}
	DeploymentName *string `field:"optional" json:"deploymentName" yaml:"deploymentName"`
	// The namespace of the Deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gke_hub_feature_membership#deployment_namespace GoogleGkeHubFeatureMembership#deployment_namespace}
	DeploymentNamespace *string `field:"optional" json:"deploymentNamespace" yaml:"deploymentNamespace"`
}

