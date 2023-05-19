package templatesms


type TemplateSmsTranslations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/template_sms#language TemplateSms#language}.
	Language *string `field:"required" json:"language" yaml:"language"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/template_sms#template TemplateSms#template}.
	Template *string `field:"required" json:"template" yaml:"template"`
}

