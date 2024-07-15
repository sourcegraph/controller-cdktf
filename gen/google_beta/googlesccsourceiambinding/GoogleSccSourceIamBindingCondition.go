package googlesccsourceiambinding


type GoogleSccSourceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_scc_source_iam_binding#expression GoogleSccSourceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_scc_source_iam_binding#title GoogleSccSourceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_scc_source_iam_binding#description GoogleSccSourceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

