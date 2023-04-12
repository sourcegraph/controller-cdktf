package googlecomputeregiondiskiammember


type GoogleComputeRegionDiskIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_region_disk_iam_member#expression GoogleComputeRegionDiskIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_region_disk_iam_member#title GoogleComputeRegionDiskIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_region_disk_iam_member#description GoogleComputeRegionDiskIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

