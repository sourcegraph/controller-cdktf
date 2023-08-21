package googlecomputeregioninstancetemplate


type GoogleComputeRegionInstanceTemplateConfidentialInstanceConfig struct {
	// Defines whether the instance should have confidential compute enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_compute_region_instance_template#enable_confidential_compute GoogleComputeRegionInstanceTemplate#enable_confidential_compute}
	EnableConfidentialCompute interface{} `field:"required" json:"enableConfidentialCompute" yaml:"enableConfidentialCompute"`
}

