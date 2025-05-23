package cloudrunv2service


type CloudRunV2ServiceScaling struct {
	// Minimum number of instances for the service, to be divided among all revisions receiving traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/cloud_run_v2_service#min_instance_count CloudRunV2Service#min_instance_count}
	MinInstanceCount *float64 `field:"optional" json:"minInstanceCount" yaml:"minInstanceCount"`
}

