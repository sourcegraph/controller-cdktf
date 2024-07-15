package googlecloudrunv2service


type GoogleCloudRunV2ServiceTemplateContainersLivenessProbeHttpGetHttpHeaders struct {
	// The header field name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_service#name GoogleCloudRunV2Service#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The header field value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_service#value GoogleCloudRunV2Service#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

