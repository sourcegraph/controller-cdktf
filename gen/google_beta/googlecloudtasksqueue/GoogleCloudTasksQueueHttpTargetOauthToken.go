package googlecloudtasksqueue


type GoogleCloudTasksQueueHttpTargetOauthToken struct {
	// Service account email to be used for generating OAuth token.
	//
	// The service account must be within the same project as the queue.
	// The caller must have iam.serviceAccounts.actAs permission for the service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_cloud_tasks_queue#service_account_email GoogleCloudTasksQueue#service_account_email}
	ServiceAccountEmail *string `field:"required" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// OAuth scope to be used for generating OAuth access token. If not specified, "https://www.googleapis.com/auth/cloud-platform" will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_cloud_tasks_queue#scope GoogleCloudTasksQueue#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

