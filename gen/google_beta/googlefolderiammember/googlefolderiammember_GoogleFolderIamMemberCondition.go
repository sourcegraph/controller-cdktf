package googlefolderiammember


type GoogleFolderIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_folder_iam_member#expression GoogleFolderIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_folder_iam_member#title GoogleFolderIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_folder_iam_member#description GoogleFolderIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

