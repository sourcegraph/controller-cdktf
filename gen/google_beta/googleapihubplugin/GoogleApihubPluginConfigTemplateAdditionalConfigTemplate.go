package googleapihubplugin


type GoogleApihubPluginConfigTemplateAdditionalConfigTemplate struct {
	// ID of the config variable. Must be unique within the configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin#id GoogleApihubPlugin#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Type of the parameter: string, int, bool etc. Possible values: VALUE_TYPE_UNSPECIFIED STRING INT BOOL SECRET ENUM MULTI_SELECT MULTI_STRING MULTI_INT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin#value_type GoogleApihubPlugin#value_type}
	ValueType *string `field:"required" json:"valueType" yaml:"valueType"`
	// Description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin#description GoogleApihubPlugin#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// enum_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin#enum_options GoogleApihubPlugin#enum_options}
	EnumOptions interface{} `field:"optional" json:"enumOptions" yaml:"enumOptions"`
	// multi_select_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin#multi_select_options GoogleApihubPlugin#multi_select_options}
	MultiSelectOptions interface{} `field:"optional" json:"multiSelectOptions" yaml:"multiSelectOptions"`
	// Flag represents that this 'ConfigVariable' must be provided for a PluginInstance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin#required GoogleApihubPlugin#required}
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// Regular expression in RE2 syntax used for validating the 'value' of a 'ConfigVariable'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin#validation_regex GoogleApihubPlugin#validation_regex}
	ValidationRegex *string `field:"optional" json:"validationRegex" yaml:"validationRegex"`
}

