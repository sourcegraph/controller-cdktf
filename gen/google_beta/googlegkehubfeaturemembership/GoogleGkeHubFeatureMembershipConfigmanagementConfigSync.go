package googlegkehubfeaturemembership


type GoogleGkeHubFeatureMembershipConfigmanagementConfigSync struct {
	// deployment_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gke_hub_feature_membership#deployment_overrides GoogleGkeHubFeatureMembership#deployment_overrides}
	DeploymentOverrides interface{} `field:"optional" json:"deploymentOverrides" yaml:"deploymentOverrides"`
	// Enables the installation of ConfigSync.
	//
	// If set to true, ConfigSync resources will be created and the other ConfigSync fields will be applied if exist. If set to false, all other ConfigSync fields will be ignored, ConfigSync resources will be deleted. If omitted, ConfigSync resources will be managed depends on the presence of the git or oci field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gke_hub_feature_membership#enabled GoogleGkeHubFeatureMembership#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// git block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gke_hub_feature_membership#git GoogleGkeHubFeatureMembership#git}
	Git *GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncGit `field:"optional" json:"git" yaml:"git"`
	// Deprecated: If Workload Identity Federation for GKE is enabled, Google Cloud Service Account is no longer needed for exporting Config Sync metrics: https://cloud.google.com/kubernetes-engine/enterprise/config-sync/docs/how-to/monitor-config-sync-cloud-monitoring#custom-monitoring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gke_hub_feature_membership#metrics_gcp_service_account_email GoogleGkeHubFeatureMembership#metrics_gcp_service_account_email}
	MetricsGcpServiceAccountEmail *string `field:"optional" json:"metricsGcpServiceAccountEmail" yaml:"metricsGcpServiceAccountEmail"`
	// oci block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gke_hub_feature_membership#oci GoogleGkeHubFeatureMembership#oci}
	Oci *GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOci `field:"optional" json:"oci" yaml:"oci"`
	// Set to true to enable the Config Sync admission webhook to prevent drifts.
	//
	// If set to `false`, disables the Config Sync admission webhook and does not prevent drifts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gke_hub_feature_membership#prevent_drift GoogleGkeHubFeatureMembership#prevent_drift}
	PreventDrift interface{} `field:"optional" json:"preventDrift" yaml:"preventDrift"`
	// Specifies whether the Config Sync Repo is in "hierarchical" or "unstructured" mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gke_hub_feature_membership#source_format GoogleGkeHubFeatureMembership#source_format}
	SourceFormat *string `field:"optional" json:"sourceFormat" yaml:"sourceFormat"`
	// Set to true to stop syncing configs for a single cluster. Default: false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gke_hub_feature_membership#stop_syncing GoogleGkeHubFeatureMembership#stop_syncing}
	StopSyncing interface{} `field:"optional" json:"stopSyncing" yaml:"stopSyncing"`
}

