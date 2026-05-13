package googlecloudrunv2workerpool


type GoogleCloudRunV2WorkerPoolInstanceSplits struct {
	// Specifies percent of the instance split to this Revision. This defaults to zero if unspecified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#percent GoogleCloudRunV2WorkerPool#percent}
	Percent *float64 `field:"optional" json:"percent" yaml:"percent"`
	// Revision to which to assign this portion of instances, if split allocation is by revision.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#revision GoogleCloudRunV2WorkerPool#revision}
	Revision *string `field:"optional" json:"revision" yaml:"revision"`
	// The allocation type for this instance split. Possible values: ["INSTANCE_SPLIT_ALLOCATION_TYPE_LATEST", "INSTANCE_SPLIT_ALLOCATION_TYPE_REVISION"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#type GoogleCloudRunV2WorkerPool#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

