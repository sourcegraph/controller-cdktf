package computeregioninstantsnapshot


type ComputeRegionInstantSnapshotTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_region_instant_snapshot#create ComputeRegionInstantSnapshot#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_region_instant_snapshot#delete ComputeRegionInstantSnapshot#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_region_instant_snapshot#update ComputeRegionInstantSnapshot#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

