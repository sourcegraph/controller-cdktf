package googlecontaineranalysisnoteiammember


type GoogleContainerAnalysisNoteIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_analysis_note_iam_member#expression GoogleContainerAnalysisNoteIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_analysis_note_iam_member#title GoogleContainerAnalysisNoteIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_analysis_note_iam_member#description GoogleContainerAnalysisNoteIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

