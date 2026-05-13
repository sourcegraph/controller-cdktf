package computeregioninstancegroupmanager


type ComputeRegionInstanceGroupManagerTargetSizePolicy struct {
	// The mode of target size policy based on which the MIG creates its VMs individually or all at once.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_region_instance_group_manager#mode ComputeRegionInstanceGroupManager#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

