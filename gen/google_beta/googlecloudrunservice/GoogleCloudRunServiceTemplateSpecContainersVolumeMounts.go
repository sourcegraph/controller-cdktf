package googlecloudrunservice


type GoogleCloudRunServiceTemplateSpecContainersVolumeMounts struct {
	// Path within the container at which the volume should be mounted.  Must not contain ':'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#mount_path GoogleCloudRunService#mount_path}
	MountPath *string `field:"required" json:"mountPath" yaml:"mountPath"`
	// This must match the Name of a Volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#name GoogleCloudRunService#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

