package googlecomputediskiammember


type GoogleComputeDiskIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_disk_iam_member#expression GoogleComputeDiskIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_disk_iam_member#title GoogleComputeDiskIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_disk_iam_member#description GoogleComputeDiskIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

