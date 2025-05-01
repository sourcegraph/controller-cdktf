package googlecloudtasksqueue


type GoogleCloudTasksQueueHttpTarget struct {
	// header_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_cloud_tasks_queue#header_overrides GoogleCloudTasksQueue#header_overrides}
	HeaderOverrides interface{} `field:"optional" json:"headerOverrides" yaml:"headerOverrides"`
	// The HTTP method to use for the request.
	//
	// When specified, it overrides HttpRequest for the task.
	// Note that if the value is set to GET the body of the task will be ignored at execution time. Possible values: ["HTTP_METHOD_UNSPECIFIED", "POST", "GET", "HEAD", "PUT", "DELETE", "PATCH", "OPTIONS"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_cloud_tasks_queue#http_method GoogleCloudTasksQueue#http_method}
	HttpMethod *string `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// oauth_token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_cloud_tasks_queue#oauth_token GoogleCloudTasksQueue#oauth_token}
	OauthToken *GoogleCloudTasksQueueHttpTargetOauthToken `field:"optional" json:"oauthToken" yaml:"oauthToken"`
	// oidc_token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_cloud_tasks_queue#oidc_token GoogleCloudTasksQueue#oidc_token}
	OidcToken *GoogleCloudTasksQueueHttpTargetOidcToken `field:"optional" json:"oidcToken" yaml:"oidcToken"`
	// uri_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_cloud_tasks_queue#uri_override GoogleCloudTasksQueue#uri_override}
	UriOverride *GoogleCloudTasksQueueHttpTargetUriOverride `field:"optional" json:"uriOverride" yaml:"uriOverride"`
}

