package googletagstagkeyiammember


type GoogleTagsTagKeyIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_tags_tag_key_iam_member#expression GoogleTagsTagKeyIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_tags_tag_key_iam_member#title GoogleTagsTagKeyIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_tags_tag_key_iam_member#description GoogleTagsTagKeyIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

