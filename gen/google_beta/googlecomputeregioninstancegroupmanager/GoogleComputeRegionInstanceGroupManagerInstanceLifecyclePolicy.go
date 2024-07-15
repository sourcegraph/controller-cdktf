package googlecomputeregioninstancegroupmanager


type GoogleComputeRegionInstanceGroupManagerInstanceLifecyclePolicy struct {
	// Default behavior for all instance or health check failures.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_instance_group_manager#default_action_on_failure GoogleComputeRegionInstanceGroupManager#default_action_on_failure}
	DefaultActionOnFailure *string `field:"optional" json:"defaultActionOnFailure" yaml:"defaultActionOnFailure"`
	// Specifies whether to apply the group's latest configuration when repairing a VM.
	//
	// Valid options are: YES, NO. If YES and you updated the group's instance template or per-instance configurations after the VM was created, then these changes are applied when VM is repaired. If NO (default), then updates are applied in accordance with the group's update policy type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_instance_group_manager#force_update_on_repair GoogleComputeRegionInstanceGroupManager#force_update_on_repair}
	ForceUpdateOnRepair *string `field:"optional" json:"forceUpdateOnRepair" yaml:"forceUpdateOnRepair"`
}

