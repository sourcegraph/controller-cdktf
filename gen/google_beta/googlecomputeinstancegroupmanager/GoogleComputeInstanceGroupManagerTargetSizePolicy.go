package googlecomputeinstancegroupmanager


type GoogleComputeInstanceGroupManagerTargetSizePolicy struct {
	// The mode of target size policy based on which the MIG creates its VMs individually or all at once.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_instance_group_manager#mode GoogleComputeInstanceGroupManager#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

