package googlecolabruntimetemplateiammember


type GoogleColabRuntimeTemplateIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_colab_runtime_template_iam_member#expression GoogleColabRuntimeTemplateIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_colab_runtime_template_iam_member#title GoogleColabRuntimeTemplateIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_colab_runtime_template_iam_member#description GoogleColabRuntimeTemplateIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

