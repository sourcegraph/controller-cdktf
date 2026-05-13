package googlecloudrunv2workerpool


type GoogleCloudRunV2WorkerPoolTemplateNodeSelector struct {
	// The GPU to attach to an instance. See https://cloud.google.com/run/docs/configuring/services/gpu for configuring GPU.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#accelerator GoogleCloudRunV2WorkerPool#accelerator}
	Accelerator *string `field:"required" json:"accelerator" yaml:"accelerator"`
}

