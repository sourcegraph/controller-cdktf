package googlecomputeinstancegroupmanager


type GoogleComputeInstanceGroupManagerInstanceLifecyclePolicyOnRepair struct {
	// Specifies whether the MIG can change a VM's zone during a repair.
	//
	// If "YES", MIG can select a different zone for the VM during a repair. Else if "NO", MIG cannot change a VM's zone during a repair. The default value of allow_changing_zone is "NO".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_instance_group_manager#allow_changing_zone GoogleComputeInstanceGroupManager#allow_changing_zone}
	AllowChangingZone *string `field:"optional" json:"allowChangingZone" yaml:"allowChangingZone"`
}

