package googlestoragemanagedfolderiammember


type GoogleStorageManagedFolderIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_managed_folder_iam_member#expression GoogleStorageManagedFolderIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_managed_folder_iam_member#title GoogleStorageManagedFolderIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_managed_folder_iam_member#description GoogleStorageManagedFolderIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

