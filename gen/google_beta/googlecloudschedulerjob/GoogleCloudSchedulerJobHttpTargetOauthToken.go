package googlecloudschedulerjob


type GoogleCloudSchedulerJobHttpTargetOauthToken struct {
	// Service account email to be used for generating OAuth token.
	//
	// The service account must be within the same project as the job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_scheduler_job#service_account_email GoogleCloudSchedulerJob#service_account_email}
	ServiceAccountEmail *string `field:"required" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// OAuth scope to be used for generating OAuth access token. If not specified, "https://www.googleapis.com/auth/cloud-platform" will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_scheduler_job#scope GoogleCloudSchedulerJob#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

