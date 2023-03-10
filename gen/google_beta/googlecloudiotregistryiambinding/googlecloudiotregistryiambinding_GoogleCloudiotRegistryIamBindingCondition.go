package googlecloudiotregistryiambinding


type GoogleCloudiotRegistryIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudiot_registry_iam_binding#expression GoogleCloudiotRegistryIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudiot_registry_iam_binding#title GoogleCloudiotRegistryIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudiot_registry_iam_binding#description GoogleCloudiotRegistryIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

