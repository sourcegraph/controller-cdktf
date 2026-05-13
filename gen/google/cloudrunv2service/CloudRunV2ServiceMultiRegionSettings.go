package cloudrunv2service


type CloudRunV2ServiceMultiRegionSettings struct {
	// The list of regions to deploy the multi-region Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_run_v2_service#regions CloudRunV2Service#regions}
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
}

