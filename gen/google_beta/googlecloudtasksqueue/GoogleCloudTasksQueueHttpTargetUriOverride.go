package googlecloudtasksqueue


type GoogleCloudTasksQueueHttpTargetUriOverride struct {
	// Host override.
	//
	// When specified, replaces the host part of the task URL.
	// For example, if the task URL is "https://www.google.com", and host value
	// is set to "example.net", the overridden URI will be changed to "https://example.net".
	// Host value cannot be an empty string (INVALID_ARGUMENT).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_cloud_tasks_queue#host GoogleCloudTasksQueue#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// path_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_cloud_tasks_queue#path_override GoogleCloudTasksQueue#path_override}
	PathOverride *GoogleCloudTasksQueueHttpTargetUriOverridePathOverride `field:"optional" json:"pathOverride" yaml:"pathOverride"`
	// Port override.
	//
	// When specified, replaces the port part of the task URI.
	// For instance, for a URI http://www.google.com/foo and port=123, the overridden URI becomes http://www.google.com:123/foo.
	// Note that the port value must be a positive integer.
	// Setting the port to 0 (Zero) clears the URI port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_cloud_tasks_queue#port GoogleCloudTasksQueue#port}
	Port *string `field:"optional" json:"port" yaml:"port"`
	// query_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_cloud_tasks_queue#query_override GoogleCloudTasksQueue#query_override}
	QueryOverride *GoogleCloudTasksQueueHttpTargetUriOverrideQueryOverride `field:"optional" json:"queryOverride" yaml:"queryOverride"`
	// Scheme override.
	//
	// When specified, the task URI scheme is replaced by the provided value (HTTP or HTTPS). Possible values: ["HTTP", "HTTPS"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_cloud_tasks_queue#scheme GoogleCloudTasksQueue#scheme}
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
	// URI Override Enforce Mode.
	//
	// When specified, determines the Target UriOverride mode. If not specified, it defaults to ALWAYS. Possible values: ["ALWAYS", "IF_NOT_EXISTS"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_cloud_tasks_queue#uri_override_enforce_mode GoogleCloudTasksQueue#uri_override_enforce_mode}
	UriOverrideEnforceMode *string `field:"optional" json:"uriOverrideEnforceMode" yaml:"uriOverrideEnforceMode"`
}

