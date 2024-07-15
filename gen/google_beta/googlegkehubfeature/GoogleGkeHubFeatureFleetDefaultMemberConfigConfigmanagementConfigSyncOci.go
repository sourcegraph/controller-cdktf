package googlegkehubfeature


type GoogleGkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOci struct {
	// Type of secret configured for access to the Git repo.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#secret_type GoogleGkeHubFeature#secret_type}
	SecretType *string `field:"required" json:"secretType" yaml:"secretType"`
	// The Google Cloud Service Account Email used for auth when secretType is gcpServiceAccount.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#gcp_service_account_email GoogleGkeHubFeature#gcp_service_account_email}
	GcpServiceAccountEmail *string `field:"optional" json:"gcpServiceAccountEmail" yaml:"gcpServiceAccountEmail"`
	// The absolute path of the directory that contains the local resources. Default: the root directory of the image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#policy_dir GoogleGkeHubFeature#policy_dir}
	PolicyDir *string `field:"optional" json:"policyDir" yaml:"policyDir"`
	// The OCI image repository URL for the package to sync from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#sync_repo GoogleGkeHubFeature#sync_repo}
	SyncRepo *string `field:"optional" json:"syncRepo" yaml:"syncRepo"`
	// Period in seconds between consecutive syncs. Default: 15.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#sync_wait_secs GoogleGkeHubFeature#sync_wait_secs}
	SyncWaitSecs *string `field:"optional" json:"syncWaitSecs" yaml:"syncWaitSecs"`
	// Version of ACM installed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature#version GoogleGkeHubFeature#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

