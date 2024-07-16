package googlestoragemanagedfolderiambinding


type GoogleStorageManagedFolderIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_managed_folder_iam_binding#expression GoogleStorageManagedFolderIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_managed_folder_iam_binding#title GoogleStorageManagedFolderIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_managed_folder_iam_binding#description GoogleStorageManagedFolderIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

