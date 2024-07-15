package googlespannerinstanceiambinding


type GoogleSpannerInstanceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_spanner_instance_iam_binding#expression GoogleSpannerInstanceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_spanner_instance_iam_binding#title GoogleSpannerInstanceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_spanner_instance_iam_binding#description GoogleSpannerInstanceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

