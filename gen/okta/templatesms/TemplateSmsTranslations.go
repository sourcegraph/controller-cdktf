package templatesms


type TemplateSmsTranslations struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/template_sms#language TemplateSms#language}.
	Language *string `field:"required" json:"language" yaml:"language"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/template_sms#template TemplateSms#template}.
	Template *string `field:"required" json:"template" yaml:"template"`
}

