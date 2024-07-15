package googlebigtableinstanceiambinding


type GoogleBigtableInstanceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigtable_instance_iam_binding#expression GoogleBigtableInstanceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigtable_instance_iam_binding#title GoogleBigtableInstanceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigtable_instance_iam_binding#description GoogleBigtableInstanceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

