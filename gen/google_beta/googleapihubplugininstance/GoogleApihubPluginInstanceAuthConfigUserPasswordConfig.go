package googleapihubplugininstance


type GoogleApihubPluginInstanceAuthConfigUserPasswordConfig struct {
	// password block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin_instance#password GoogleApihubPluginInstance#password}
	Password *GoogleApihubPluginInstanceAuthConfigUserPasswordConfigPassword `field:"required" json:"password" yaml:"password"`
	// Username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin_instance#username GoogleApihubPluginInstance#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

