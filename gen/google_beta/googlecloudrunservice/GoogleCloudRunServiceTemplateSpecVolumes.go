package googlecloudrunservice


type GoogleCloudRunServiceTemplateSpecVolumes struct {
	// Volume's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_cloud_run_service#name GoogleCloudRunService#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_cloud_run_service#secret GoogleCloudRunService#secret}
	Secret *GoogleCloudRunServiceTemplateSpecVolumesSecret `field:"required" json:"secret" yaml:"secret"`
}

