package googletagstagvalueiambinding


type GoogleTagsTagValueIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_tags_tag_value_iam_binding#expression GoogleTagsTagValueIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_tags_tag_value_iam_binding#title GoogleTagsTagValueIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_tags_tag_value_iam_binding#description GoogleTagsTagValueIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

