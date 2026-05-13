package googleapihubplugininstance


type GoogleApihubPluginInstanceAuthConfigApiKeyConfigApiKey struct {
	// The resource name of the secret version in the format, format as: 'projects/* /secrets/* /versions/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin_instance#secret_version GoogleApihubPluginInstance#secret_version}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	SecretVersion *string `field:"required" json:"secretVersion" yaml:"secretVersion"`
}

