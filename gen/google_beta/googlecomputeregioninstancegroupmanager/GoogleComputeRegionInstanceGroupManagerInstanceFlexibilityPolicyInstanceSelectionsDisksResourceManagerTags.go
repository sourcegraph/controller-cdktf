package googlecomputeregioninstancegroupmanager


type GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksResourceManagerTags struct {
	// The unique key of the resource manager tag to assign to disks. Keys must be in the format tagKeys/{tag_key_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_instance_group_manager#key GoogleComputeRegionInstanceGroupManager#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value of the resource manager tag to assign to disks. Values must be in the format tagValues/456.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_instance_group_manager#value GoogleComputeRegionInstanceGroupManager#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

