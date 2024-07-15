package googlecloudrunv2service


type GoogleCloudRunV2ServiceScaling struct {
	// Minimum number of instances for the service, to be divided among all revisions receiving traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_service#min_instance_count GoogleCloudRunV2Service#min_instance_count}
	MinInstanceCount *float64 `field:"optional" json:"minInstanceCount" yaml:"minInstanceCount"`
}

