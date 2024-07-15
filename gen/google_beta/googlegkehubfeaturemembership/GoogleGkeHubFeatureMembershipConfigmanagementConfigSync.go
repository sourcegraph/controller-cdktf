package googlegkehubfeaturemembership


type GoogleGkeHubFeatureMembershipConfigmanagementConfigSync struct {
	// git block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#git GoogleGkeHubFeatureMembership#git}
	Git *GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncGit `field:"optional" json:"git" yaml:"git"`
	// The Email of the Google Cloud Service Account (GSA) used for exporting Config Sync metrics to Cloud Monitoring.
	//
	// The GSA should have the Monitoring Metric Writer(roles/monitoring.metricWriter) IAM role. The Kubernetes ServiceAccount `default` in the namespace `config-management-monitoring` should be bound to the GSA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#metrics_gcp_service_account_email GoogleGkeHubFeatureMembership#metrics_gcp_service_account_email}
	MetricsGcpServiceAccountEmail *string `field:"optional" json:"metricsGcpServiceAccountEmail" yaml:"metricsGcpServiceAccountEmail"`
	// oci block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#oci GoogleGkeHubFeatureMembership#oci}
	Oci *GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOci `field:"optional" json:"oci" yaml:"oci"`
	// Set to true to enable the Config Sync admission webhook to prevent drifts.
	//
	// If set to `false`, disables the Config Sync admission webhook and does not prevent drifts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#prevent_drift GoogleGkeHubFeatureMembership#prevent_drift}
	PreventDrift interface{} `field:"optional" json:"preventDrift" yaml:"preventDrift"`
	// Specifies whether the Config Sync Repo is in "hierarchical" or "unstructured" mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#source_format GoogleGkeHubFeatureMembership#source_format}
	SourceFormat *string `field:"optional" json:"sourceFormat" yaml:"sourceFormat"`
}

