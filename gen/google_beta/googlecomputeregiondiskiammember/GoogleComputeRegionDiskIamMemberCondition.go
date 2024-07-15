package googlecomputeregiondiskiammember


type GoogleComputeRegionDiskIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_disk_iam_member#expression GoogleComputeRegionDiskIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_disk_iam_member#title GoogleComputeRegionDiskIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_disk_iam_member#description GoogleComputeRegionDiskIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

