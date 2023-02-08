package googleprivatecacertificatetemplateiammember


type GooglePrivatecaCertificateTemplateIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_privateca_certificate_template_iam_member#expression GooglePrivatecaCertificateTemplateIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_privateca_certificate_template_iam_member#title GooglePrivatecaCertificateTemplateIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_privateca_certificate_template_iam_member#description GooglePrivatecaCertificateTemplateIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

