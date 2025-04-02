package googlekmsekmconnectioniambinding


type GoogleKmsEkmConnectionIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_kms_ekm_connection_iam_binding#expression GoogleKmsEkmConnectionIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_kms_ekm_connection_iam_binding#title GoogleKmsEkmConnectionIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_kms_ekm_connection_iam_binding#description GoogleKmsEkmConnectionIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

