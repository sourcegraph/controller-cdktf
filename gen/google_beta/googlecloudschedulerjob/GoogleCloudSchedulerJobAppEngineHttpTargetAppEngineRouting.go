package googlecloudschedulerjob


type GoogleCloudSchedulerJobAppEngineHttpTargetAppEngineRouting struct {
	// App instance. By default, the job is sent to an instance which is available when the job is attempted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_scheduler_job#instance GoogleCloudSchedulerJob#instance}
	Instance *string `field:"optional" json:"instance" yaml:"instance"`
	// App service.
	//
	// By default, the job is sent to the service which is the default service when the job is attempted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_scheduler_job#service GoogleCloudSchedulerJob#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
	// App version.
	//
	// By default, the job is sent to the version which is the default version when the job is attempted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_scheduler_job#version GoogleCloudSchedulerJob#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

