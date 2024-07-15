package googledataplextaskiambinding


type GoogleDataplexTaskIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_task_iam_binding#expression GoogleDataplexTaskIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_task_iam_binding#title GoogleDataplexTaskIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_task_iam_binding#description GoogleDataplexTaskIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

