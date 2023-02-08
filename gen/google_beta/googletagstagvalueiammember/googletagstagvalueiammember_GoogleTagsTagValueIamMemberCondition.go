package googletagstagvalueiammember


type GoogleTagsTagValueIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_tags_tag_value_iam_member#expression GoogleTagsTagValueIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_tags_tag_value_iam_member#title GoogleTagsTagValueIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_tags_tag_value_iam_member#description GoogleTagsTagValueIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

