package googlecloudrunv2job


type GoogleCloudRunV2JobTemplateTemplateVolumesNfs struct {
	// Hostname or IP address of the NFS server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_job#server GoogleCloudRunV2Job#server}
	Server *string `field:"required" json:"server" yaml:"server"`
	// Path that is exported by the NFS server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_job#path GoogleCloudRunV2Job#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// If true, mount this volume as read-only in all mounts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_job#read_only GoogleCloudRunV2Job#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
}

