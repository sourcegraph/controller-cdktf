package templatesms


type TemplateSmsTranslations struct {
	// The language to map the template to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/template_sms#language TemplateSms#language}
	Language *string `field:"required" json:"language" yaml:"language"`
	// The SMS message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/template_sms#template TemplateSms#template}
	Template *string `field:"required" json:"template" yaml:"template"`
}

