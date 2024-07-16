package googleartifactregistryrepositoryiambinding


type GoogleArtifactRegistryRepositoryIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_artifact_registry_repository_iam_binding#expression GoogleArtifactRegistryRepositoryIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_artifact_registry_repository_iam_binding#title GoogleArtifactRegistryRepositoryIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_artifact_registry_repository_iam_binding#description GoogleArtifactRegistryRepositoryIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

