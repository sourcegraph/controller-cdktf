package googlegkehubfeaturemembership


type GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncGit struct {
	// The GCP Service Account Email used for auth when secretType is gcpServiceAccount.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#gcp_service_account_email GoogleGkeHubFeatureMembership#gcp_service_account_email}
	GcpServiceAccountEmail *string `field:"optional" json:"gcpServiceAccountEmail" yaml:"gcpServiceAccountEmail"`
	// URL for the HTTPS proxy to be used when communicating with the Git repo.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#https_proxy GoogleGkeHubFeatureMembership#https_proxy}
	HttpsProxy *string `field:"optional" json:"httpsProxy" yaml:"httpsProxy"`
	// The path within the Git repository that represents the top level of the repo to sync.
	//
	// Default: the root directory of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#policy_dir GoogleGkeHubFeatureMembership#policy_dir}
	PolicyDir *string `field:"optional" json:"policyDir" yaml:"policyDir"`
	// Type of secret configured for access to the Git repo.
	//
	// Must be one of ssh, cookiefile, gcenode, token, gcpserviceaccount or none. The validation of this is case-sensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#secret_type GoogleGkeHubFeatureMembership#secret_type}
	SecretType *string `field:"optional" json:"secretType" yaml:"secretType"`
	// The branch of the repository to sync from. Default: master.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#sync_branch GoogleGkeHubFeatureMembership#sync_branch}
	SyncBranch *string `field:"optional" json:"syncBranch" yaml:"syncBranch"`
	// The URL of the Git repository to use as the source of truth.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#sync_repo GoogleGkeHubFeatureMembership#sync_repo}
	SyncRepo *string `field:"optional" json:"syncRepo" yaml:"syncRepo"`
	// Git revision (tag or hash) to check out. Default HEAD.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#sync_rev GoogleGkeHubFeatureMembership#sync_rev}
	SyncRev *string `field:"optional" json:"syncRev" yaml:"syncRev"`
	// Period in seconds between consecutive syncs. Default: 15.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#sync_wait_secs GoogleGkeHubFeatureMembership#sync_wait_secs}
	SyncWaitSecs *string `field:"optional" json:"syncWaitSecs" yaml:"syncWaitSecs"`
}

