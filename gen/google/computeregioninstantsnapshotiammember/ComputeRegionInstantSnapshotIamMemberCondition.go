package computeregioninstantsnapshotiammember


type ComputeRegionInstantSnapshotIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_region_instant_snapshot_iam_member#expression ComputeRegionInstantSnapshotIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_region_instant_snapshot_iam_member#title ComputeRegionInstantSnapshotIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_region_instant_snapshot_iam_member#description ComputeRegionInstantSnapshotIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

