package googlecomputeregiondiskiambinding


type GoogleComputeRegionDiskIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_disk_iam_binding#expression GoogleComputeRegionDiskIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_disk_iam_binding#title GoogleComputeRegionDiskIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_disk_iam_binding#description GoogleComputeRegionDiskIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

