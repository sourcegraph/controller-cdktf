package googlecloudtasksqueue


type GoogleCloudTasksQueueHttpTargetUriOverrideQueryOverride struct {
	// The query parameters (e.g., qparam1=123&qparam2=456). Default is an empty string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_cloud_tasks_queue#query_params GoogleCloudTasksQueue#query_params}
	QueryParams *string `field:"optional" json:"queryParams" yaml:"queryParams"`
}

