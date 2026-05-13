package googlecomputeregioninstancegroupmanager


type GoogleComputeRegionInstanceGroupManagerTargetSizePolicy struct {
	// The mode of target size policy based on which the MIG creates its VMs individually or all at once.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_instance_group_manager#mode GoogleComputeRegionInstanceGroupManager#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

