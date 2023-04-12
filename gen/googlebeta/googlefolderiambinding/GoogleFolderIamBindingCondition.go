package googlefolderiambinding


type GoogleFolderIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_folder_iam_binding#expression GoogleFolderIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_folder_iam_binding#title GoogleFolderIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_folder_iam_binding#description GoogleFolderIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

