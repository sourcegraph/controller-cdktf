package googlecomputeregioninstancegroupmanager


type GoogleComputeRegionInstanceGroupManagerNamedPort struct {
	// The name of the port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_instance_group_manager#name GoogleComputeRegionInstanceGroupManager#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The port number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_instance_group_manager#port GoogleComputeRegionInstanceGroupManager#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

