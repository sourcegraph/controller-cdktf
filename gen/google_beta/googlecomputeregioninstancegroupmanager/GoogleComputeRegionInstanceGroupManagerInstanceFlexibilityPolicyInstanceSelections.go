package googlecomputeregioninstancegroupmanager


type GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelections struct {
	// Full machine-type names, e.g. "n1-standard-16".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_instance_group_manager#machine_types GoogleComputeRegionInstanceGroupManager#machine_types}
	MachineTypes *[]*string `field:"required" json:"machineTypes" yaml:"machineTypes"`
	// Instance selection name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_instance_group_manager#name GoogleComputeRegionInstanceGroupManager#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// disks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_instance_group_manager#disks GoogleComputeRegionInstanceGroupManager#disks}
	Disks interface{} `field:"optional" json:"disks" yaml:"disks"`
	// Name of the minimum CPU platform to be used by this instance selection. e.g. 'Intel Ice Lake'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_instance_group_manager#min_cpu_platform GoogleComputeRegionInstanceGroupManager#min_cpu_platform}
	MinCpuPlatform *string `field:"optional" json:"minCpuPlatform" yaml:"minCpuPlatform"`
	// Preference of this instance selection.
	//
	// Lower number means higher preference. MIG will first try to create a VM based on the machine-type with lowest rank and fallback to next rank based on availability. Machine types and instance selections with the same rank have the same preference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_instance_group_manager#rank GoogleComputeRegionInstanceGroupManager#rank}
	Rank *float64 `field:"optional" json:"rank" yaml:"rank"`
}

