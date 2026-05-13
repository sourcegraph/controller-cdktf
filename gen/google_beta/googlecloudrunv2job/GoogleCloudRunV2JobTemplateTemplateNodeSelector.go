package googlecloudrunv2job


type GoogleCloudRunV2JobTemplateTemplateNodeSelector struct {
	// The GPU to attach to an instance. See https://cloud.google.com/run/docs/configuring/jobs/gpu for configuring GPU.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_job#accelerator GoogleCloudRunV2Job#accelerator}
	Accelerator *string `field:"required" json:"accelerator" yaml:"accelerator"`
}

