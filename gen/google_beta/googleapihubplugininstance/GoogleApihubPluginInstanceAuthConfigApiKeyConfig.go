package googleapihubplugininstance


type GoogleApihubPluginInstanceAuthConfigApiKeyConfig struct {
	// api_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin_instance#api_key GoogleApihubPluginInstance#api_key}
	ApiKey *GoogleApihubPluginInstanceAuthConfigApiKeyConfigApiKey `field:"required" json:"apiKey" yaml:"apiKey"`
	// The location of the API key. The default value is QUERY. Possible values: HTTP_ELEMENT_LOCATION_UNSPECIFIED QUERY HEADER PATH BODY COOKIE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin_instance#http_element_location GoogleApihubPluginInstance#http_element_location}
	HttpElementLocation *string `field:"required" json:"httpElementLocation" yaml:"httpElementLocation"`
	// The parameter name of the API key. E.g. If the API request is "https://example.com/act?api_key=", "api_key" would be the parameter name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin_instance#name GoogleApihubPluginInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

