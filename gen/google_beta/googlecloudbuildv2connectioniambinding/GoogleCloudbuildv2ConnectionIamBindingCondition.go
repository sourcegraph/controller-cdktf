package googlecloudbuildv2connectioniambinding


type GoogleCloudbuildv2ConnectionIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuildv2_connection_iam_binding#expression GoogleCloudbuildv2ConnectionIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuildv2_connection_iam_binding#title GoogleCloudbuildv2ConnectionIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuildv2_connection_iam_binding#description GoogleCloudbuildv2ConnectionIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

