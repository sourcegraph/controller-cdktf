package googlecloudrunv2workerpool


type GoogleCloudRunV2WorkerPoolTemplateVolumesNfs struct {
	// Path that is exported by the NFS server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#path GoogleCloudRunV2WorkerPool#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Hostname or IP address of the NFS server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#server GoogleCloudRunV2WorkerPool#server}
	Server *string `field:"required" json:"server" yaml:"server"`
	// If true, mount the NFS volume as read only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#read_only GoogleCloudRunV2WorkerPool#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
}

