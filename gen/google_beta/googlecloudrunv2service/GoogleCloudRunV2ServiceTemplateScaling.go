package googlecloudrunv2service


type GoogleCloudRunV2ServiceTemplateScaling struct {
	// Maximum number of serving instances that this resource should have.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_service#max_instance_count GoogleCloudRunV2Service#max_instance_count}
	MaxInstanceCount *float64 `field:"optional" json:"maxInstanceCount" yaml:"maxInstanceCount"`
	// Minimum number of serving instances that this resource should have.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_service#min_instance_count GoogleCloudRunV2Service#min_instance_count}
	MinInstanceCount *float64 `field:"optional" json:"minInstanceCount" yaml:"minInstanceCount"`
}

