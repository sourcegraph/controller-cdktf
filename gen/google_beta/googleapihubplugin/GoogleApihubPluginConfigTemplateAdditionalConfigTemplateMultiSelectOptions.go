package googleapihubplugin


type GoogleApihubPluginConfigTemplateAdditionalConfigTemplateMultiSelectOptions struct {
	// Display name of the option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin#display_name GoogleApihubPlugin#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Id of the option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin#id GoogleApihubPlugin#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Description of the option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin#description GoogleApihubPlugin#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

