package googlecloudrunv2service


type GoogleCloudRunV2ServiceTemplateNodeSelector struct {
	// The GPU to attach to an instance. See https://cloud.google.com/run/docs/configuring/services/gpu for configuring GPU.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_cloud_run_v2_service#accelerator GoogleCloudRunV2Service#accelerator}
	Accelerator *string `field:"required" json:"accelerator" yaml:"accelerator"`
}

