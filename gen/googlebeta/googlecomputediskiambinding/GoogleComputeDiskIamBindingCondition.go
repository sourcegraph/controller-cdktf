package googlecomputediskiambinding


type GoogleComputeDiskIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_disk_iam_binding#expression GoogleComputeDiskIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_disk_iam_binding#title GoogleComputeDiskIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_disk_iam_binding#description GoogleComputeDiskIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

