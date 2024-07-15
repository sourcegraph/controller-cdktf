package googlecloudrunservice


type GoogleCloudRunServiceTemplateSpecVolumes struct {
	// Volume's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#name GoogleCloudRunService#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// csi block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#csi GoogleCloudRunService#csi}
	Csi *GoogleCloudRunServiceTemplateSpecVolumesCsi `field:"optional" json:"csi" yaml:"csi"`
	// empty_dir block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#empty_dir GoogleCloudRunService#empty_dir}
	EmptyDir *GoogleCloudRunServiceTemplateSpecVolumesEmptyDir `field:"optional" json:"emptyDir" yaml:"emptyDir"`
	// nfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#nfs GoogleCloudRunService#nfs}
	Nfs *GoogleCloudRunServiceTemplateSpecVolumesNfs `field:"optional" json:"nfs" yaml:"nfs"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#secret GoogleCloudRunService#secret}
	Secret *GoogleCloudRunServiceTemplateSpecVolumesSecret `field:"optional" json:"secret" yaml:"secret"`
}

