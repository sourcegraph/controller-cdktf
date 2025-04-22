package googlecomputeinstantsnapshotiammember


type GoogleComputeInstantSnapshotIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_instant_snapshot_iam_member#expression GoogleComputeInstantSnapshotIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_instant_snapshot_iam_member#title GoogleComputeInstantSnapshotIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_instant_snapshot_iam_member#description GoogleComputeInstantSnapshotIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

