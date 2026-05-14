package googlecomputeregioninstantsnapshotiammember


type GoogleComputeRegionInstantSnapshotIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_compute_region_instant_snapshot_iam_member#expression GoogleComputeRegionInstantSnapshotIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_compute_region_instant_snapshot_iam_member#title GoogleComputeRegionInstantSnapshotIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_compute_region_instant_snapshot_iam_member#description GoogleComputeRegionInstantSnapshotIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

