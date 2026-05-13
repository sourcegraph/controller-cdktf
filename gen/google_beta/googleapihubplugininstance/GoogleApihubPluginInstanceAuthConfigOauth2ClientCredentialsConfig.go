package googleapihubplugininstance


type GoogleApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfig struct {
	// The client identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin_instance#client_id GoogleApihubPluginInstance#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// client_secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin_instance#client_secret GoogleApihubPluginInstance#client_secret}
	ClientSecret *GoogleApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecret `field:"required" json:"clientSecret" yaml:"clientSecret"`
}

