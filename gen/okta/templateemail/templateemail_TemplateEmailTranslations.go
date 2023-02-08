package templateemail


type TemplateEmailTranslations struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/template_email#language TemplateEmail#language}.
	Language *string `field:"required" json:"language" yaml:"language"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/template_email#subject TemplateEmail#subject}.
	Subject *string `field:"required" json:"subject" yaml:"subject"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/template_email#template TemplateEmail#template}.
	Template *string `field:"required" json:"template" yaml:"template"`
}

