package googlecomputeregioninstancegroupmanager


type GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksLabels struct {
	// The unique key of the label to assign to disks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_instance_group_manager#key GoogleComputeRegionInstanceGroupManager#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value of the label to assign to disks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_instance_group_manager#value GoogleComputeRegionInstanceGroupManager#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

