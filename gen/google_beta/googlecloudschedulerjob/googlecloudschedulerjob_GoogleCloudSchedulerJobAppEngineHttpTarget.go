package googlecloudschedulerjob


type GoogleCloudSchedulerJobAppEngineHttpTarget struct {
	// The relative URI.
	//
	// The relative URL must begin with "/" and must be a valid HTTP relative URL.
	// It can contain a path, query string arguments, and \# fragments.
	// If the relative URL is empty, then the root path "/" will be used.
	// No spaces are allowed, and the maximum length allowed is 2083 characters
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_scheduler_job#relative_uri GoogleCloudSchedulerJob#relative_uri}
	RelativeUri *string `field:"required" json:"relativeUri" yaml:"relativeUri"`
	// app_engine_routing block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_scheduler_job#app_engine_routing GoogleCloudSchedulerJob#app_engine_routing}
	AppEngineRouting *GoogleCloudSchedulerJobAppEngineHttpTargetAppEngineRouting `field:"optional" json:"appEngineRouting" yaml:"appEngineRouting"`
	// HTTP request body.
	//
	// A request body is allowed only if the HTTP method is POST or PUT.
	// It will result in invalid argument error to set a body on a job with an incompatible HttpMethod.
	//
	// A base64-encoded string.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_scheduler_job#body GoogleCloudSchedulerJob#body}
	Body *string `field:"optional" json:"body" yaml:"body"`
	// HTTP request headers.
	//
	// This map contains the header field names and values.
	// Headers can be set when the job is created.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_scheduler_job#headers GoogleCloudSchedulerJob#headers}
	Headers *map[string]*string `field:"optional" json:"headers" yaml:"headers"`
	// Which HTTP method to use for the request.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_scheduler_job#http_method GoogleCloudSchedulerJob#http_method}
	HttpMethod *string `field:"optional" json:"httpMethod" yaml:"httpMethod"`
}
