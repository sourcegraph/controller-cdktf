package googlecloudiotregistryiammember


type GoogleCloudiotRegistryIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudiot_registry_iam_member#expression GoogleCloudiotRegistryIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudiot_registry_iam_member#title GoogleCloudiotRegistryIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudiot_registry_iam_member#description GoogleCloudiotRegistryIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

