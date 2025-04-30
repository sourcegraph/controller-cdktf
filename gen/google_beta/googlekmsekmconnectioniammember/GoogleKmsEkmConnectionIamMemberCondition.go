package googlekmsekmconnectioniammember


type GoogleKmsEkmConnectionIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_kms_ekm_connection_iam_member#expression GoogleKmsEkmConnectionIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_kms_ekm_connection_iam_member#title GoogleKmsEkmConnectionIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_kms_ekm_connection_iam_member#description GoogleKmsEkmConnectionIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

