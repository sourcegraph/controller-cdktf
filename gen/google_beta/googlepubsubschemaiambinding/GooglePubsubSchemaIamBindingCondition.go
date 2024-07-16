package googlepubsubschemaiambinding


type GooglePubsubSchemaIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_pubsub_schema_iam_binding#expression GooglePubsubSchemaIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_pubsub_schema_iam_binding#title GooglePubsubSchemaIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_pubsub_schema_iam_binding#description GooglePubsubSchemaIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

