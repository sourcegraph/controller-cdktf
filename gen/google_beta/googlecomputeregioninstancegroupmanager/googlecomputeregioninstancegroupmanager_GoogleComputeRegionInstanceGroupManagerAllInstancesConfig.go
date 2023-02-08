package googlecomputeregioninstancegroupmanager


type GoogleComputeRegionInstanceGroupManagerAllInstancesConfig struct {
	// The label key-value pairs that you want to patch onto the instance,.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_region_instance_group_manager#labels GoogleComputeRegionInstanceGroupManager#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The metadata key-value pairs that you want to patch onto the instance.
	//
	// For more information, see Project and instance metadata,
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_region_instance_group_manager#metadata GoogleComputeRegionInstanceGroupManager#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
}

