package googlespannerinstanceiammember


type GoogleSpannerInstanceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_spanner_instance_iam_member#expression GoogleSpannerInstanceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_spanner_instance_iam_member#title GoogleSpannerInstanceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_spanner_instance_iam_member#description GoogleSpannerInstanceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

