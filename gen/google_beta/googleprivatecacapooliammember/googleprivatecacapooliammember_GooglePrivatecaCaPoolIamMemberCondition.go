package googleprivatecacapooliammember


type GooglePrivatecaCaPoolIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_privateca_ca_pool_iam_member#expression GooglePrivatecaCaPoolIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_privateca_ca_pool_iam_member#title GooglePrivatecaCaPoolIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_privateca_ca_pool_iam_member#description GooglePrivatecaCaPoolIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}
