package googlecloudrunservice


type GoogleCloudRunServiceTemplateSpecVolumesNfs struct {
	// Path exported by the NFS server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#path GoogleCloudRunService#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// IP address or hostname of the NFS server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#server GoogleCloudRunService#server}
	Server *string `field:"required" json:"server" yaml:"server"`
	// If true, mount the NFS volume as read only in all mounts. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#read_only GoogleCloudRunService#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
}

