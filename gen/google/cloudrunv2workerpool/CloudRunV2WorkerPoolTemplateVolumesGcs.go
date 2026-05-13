package cloudrunv2workerpool


type CloudRunV2WorkerPoolTemplateVolumesGcs struct {
	// GCS Bucket name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_run_v2_worker_pool#bucket CloudRunV2WorkerPool#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// A list of flags to pass to the gcsfuse command for configuring this volume.
	//
	// Flags should be passed without leading dashes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_run_v2_worker_pool#mount_options CloudRunV2WorkerPool#mount_options}
	MountOptions *[]*string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
	// If true, mount the GCS bucket as read-only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_run_v2_worker_pool#read_only CloudRunV2WorkerPool#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
}

