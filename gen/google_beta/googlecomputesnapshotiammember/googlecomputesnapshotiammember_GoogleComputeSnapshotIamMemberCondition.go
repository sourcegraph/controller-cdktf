package googlecomputesnapshotiammember


type GoogleComputeSnapshotIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_snapshot_iam_member#expression GoogleComputeSnapshotIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_snapshot_iam_member#title GoogleComputeSnapshotIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_snapshot_iam_member#description GoogleComputeSnapshotIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

