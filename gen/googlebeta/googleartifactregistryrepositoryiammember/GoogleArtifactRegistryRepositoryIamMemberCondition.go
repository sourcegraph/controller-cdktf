package googleartifactregistryrepositoryiammember


type GoogleArtifactRegistryRepositoryIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_artifact_registry_repository_iam_member#expression GoogleArtifactRegistryRepositoryIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_artifact_registry_repository_iam_member#title GoogleArtifactRegistryRepositoryIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_artifact_registry_repository_iam_member#description GoogleArtifactRegistryRepositoryIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

