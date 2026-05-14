package computeregioninstantsnapshotiambinding


type ComputeRegionInstantSnapshotIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_region_instant_snapshot_iam_binding#expression ComputeRegionInstantSnapshotIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_region_instant_snapshot_iam_binding#title ComputeRegionInstantSnapshotIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_region_instant_snapshot_iam_binding#description ComputeRegionInstantSnapshotIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

