package googlecloudrunv2service


type GoogleCloudRunV2ServiceScaling struct {
	// Total instance count for the service in manual scaling mode.
	//
	// This number of instances is divided among all revisions with specified traffic based on the percent of traffic they are receiving.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_service#manual_instance_count GoogleCloudRunV2Service#manual_instance_count}
	ManualInstanceCount *float64 `field:"optional" json:"manualInstanceCount" yaml:"manualInstanceCount"`
	// Combined maximum number of instances for all revisions receiving traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_service#max_instance_count GoogleCloudRunV2Service#max_instance_count}
	MaxInstanceCount *float64 `field:"optional" json:"maxInstanceCount" yaml:"maxInstanceCount"`
	// Minimum number of instances for the service, to be divided among all revisions receiving traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_service#min_instance_count GoogleCloudRunV2Service#min_instance_count}
	MinInstanceCount *float64 `field:"optional" json:"minInstanceCount" yaml:"minInstanceCount"`
	// The [scaling mode](https://cloud.google.com/run/docs/reference/rest/v2/projects.locations.services#scalingmode) for the service. Possible values: ["AUTOMATIC", "MANUAL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_service#scaling_mode GoogleCloudRunV2Service#scaling_mode}
	ScalingMode *string `field:"optional" json:"scalingMode" yaml:"scalingMode"`
}

