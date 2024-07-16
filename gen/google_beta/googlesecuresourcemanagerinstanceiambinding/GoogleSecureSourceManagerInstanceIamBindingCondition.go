package googlesecuresourcemanagerinstanceiambinding


type GoogleSecureSourceManagerInstanceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secure_source_manager_instance_iam_binding#expression GoogleSecureSourceManagerInstanceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secure_source_manager_instance_iam_binding#title GoogleSecureSourceManagerInstanceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secure_source_manager_instance_iam_binding#description GoogleSecureSourceManagerInstanceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

