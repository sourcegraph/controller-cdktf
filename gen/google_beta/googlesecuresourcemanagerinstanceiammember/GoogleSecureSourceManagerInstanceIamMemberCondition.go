package googlesecuresourcemanagerinstanceiammember


type GoogleSecureSourceManagerInstanceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secure_source_manager_instance_iam_member#expression GoogleSecureSourceManagerInstanceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secure_source_manager_instance_iam_member#title GoogleSecureSourceManagerInstanceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secure_source_manager_instance_iam_member#description GoogleSecureSourceManagerInstanceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

