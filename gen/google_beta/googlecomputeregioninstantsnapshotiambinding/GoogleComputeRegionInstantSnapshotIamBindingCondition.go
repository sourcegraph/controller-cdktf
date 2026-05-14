package googlecomputeregioninstantsnapshotiambinding


type GoogleComputeRegionInstantSnapshotIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_compute_region_instant_snapshot_iam_binding#expression GoogleComputeRegionInstantSnapshotIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_compute_region_instant_snapshot_iam_binding#title GoogleComputeRegionInstantSnapshotIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_compute_region_instant_snapshot_iam_binding#description GoogleComputeRegionInstantSnapshotIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

