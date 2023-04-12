package googlecloudschedulerjob


type GoogleCloudSchedulerJobTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_scheduler_job#create GoogleCloudSchedulerJob#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_scheduler_job#delete GoogleCloudSchedulerJob#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_scheduler_job#update GoogleCloudSchedulerJob#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

