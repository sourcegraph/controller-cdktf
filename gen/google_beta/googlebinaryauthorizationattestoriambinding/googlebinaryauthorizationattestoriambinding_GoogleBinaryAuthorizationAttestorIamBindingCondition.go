package googlebinaryauthorizationattestoriambinding


type GoogleBinaryAuthorizationAttestorIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_binary_authorization_attestor_iam_binding#expression GoogleBinaryAuthorizationAttestorIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_binary_authorization_attestor_iam_binding#title GoogleBinaryAuthorizationAttestorIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_binary_authorization_attestor_iam_binding#description GoogleBinaryAuthorizationAttestorIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

