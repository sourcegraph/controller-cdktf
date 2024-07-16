package googlebinaryauthorizationattestoriammember


type GoogleBinaryAuthorizationAttestorIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_binary_authorization_attestor_iam_member#expression GoogleBinaryAuthorizationAttestorIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_binary_authorization_attestor_iam_member#title GoogleBinaryAuthorizationAttestorIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_binary_authorization_attestor_iam_member#description GoogleBinaryAuthorizationAttestorIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

