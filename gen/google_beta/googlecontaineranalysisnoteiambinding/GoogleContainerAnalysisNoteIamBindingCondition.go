package googlecontaineranalysisnoteiambinding


type GoogleContainerAnalysisNoteIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_analysis_note_iam_binding#expression GoogleContainerAnalysisNoteIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_analysis_note_iam_binding#title GoogleContainerAnalysisNoteIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_analysis_note_iam_binding#description GoogleContainerAnalysisNoteIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

