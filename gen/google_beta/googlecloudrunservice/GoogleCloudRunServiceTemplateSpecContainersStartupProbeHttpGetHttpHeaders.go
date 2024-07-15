package googlecloudrunservice


type GoogleCloudRunServiceTemplateSpecContainersStartupProbeHttpGetHttpHeaders struct {
	// The header field name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#name GoogleCloudRunService#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The header field value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#value GoogleCloudRunService#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

