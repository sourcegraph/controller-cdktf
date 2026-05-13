package googleapihubplugin


type GoogleApihubPluginConfigTemplateAuthConfigTemplate struct {
	// The list of authentication types supported by the plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin#supported_auth_types GoogleApihubPlugin#supported_auth_types}
	SupportedAuthTypes *[]*string `field:"required" json:"supportedAuthTypes" yaml:"supportedAuthTypes"`
	// service_account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin#service_account GoogleApihubPlugin#service_account}
	ServiceAccount *GoogleApihubPluginConfigTemplateAuthConfigTemplateServiceAccount `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
}

