package googlespannerinstanceiambinding


type GoogleSpannerInstanceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_spanner_instance_iam_binding#expression GoogleSpannerInstanceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_spanner_instance_iam_binding#title GoogleSpannerInstanceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_spanner_instance_iam_binding#description GoogleSpannerInstanceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

