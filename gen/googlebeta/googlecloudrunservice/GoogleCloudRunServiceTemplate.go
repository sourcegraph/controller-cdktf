package googlecloudrunservice


type GoogleCloudRunServiceTemplate struct {
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_service#metadata GoogleCloudRunService#metadata}
	Metadata *GoogleCloudRunServiceTemplateMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_service#spec GoogleCloudRunService#spec}
	Spec *GoogleCloudRunServiceTemplateSpec `field:"optional" json:"spec" yaml:"spec"`
}

