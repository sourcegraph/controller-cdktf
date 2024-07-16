package googlecloudschedulerjob


type GoogleCloudSchedulerJobHttpTargetOidcToken struct {
	// Service account email to be used for generating OAuth token.
	//
	// The service account must be within the same project as the job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_scheduler_job#service_account_email GoogleCloudSchedulerJob#service_account_email}
	ServiceAccountEmail *string `field:"required" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// Audience to be used when generating OIDC token. If not specified, the URI specified in target will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_scheduler_job#audience GoogleCloudSchedulerJob#audience}
	Audience *string `field:"optional" json:"audience" yaml:"audience"`
}

