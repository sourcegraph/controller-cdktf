package googlegeminicodetoolssetting


type GoogleGeminiCodeToolsSettingEnabledToolConfig struct {
	// Key of the configuration item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gemini_code_tools_setting#key GoogleGeminiCodeToolsSetting#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Value of the configuration item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gemini_code_tools_setting#value GoogleGeminiCodeToolsSetting#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

