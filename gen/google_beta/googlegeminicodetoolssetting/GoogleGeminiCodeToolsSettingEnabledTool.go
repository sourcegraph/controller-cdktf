package googlegeminicodetoolssetting


type GoogleGeminiCodeToolsSettingEnabledTool struct {
	// Handle used to invoke the tool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gemini_code_tools_setting#handle GoogleGeminiCodeToolsSetting#handle}
	Handle *string `field:"required" json:"handle" yaml:"handle"`
	// Link to the Tool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gemini_code_tools_setting#tool GoogleGeminiCodeToolsSetting#tool}
	Tool *string `field:"required" json:"tool" yaml:"tool"`
	// Link to the Dev Connect Account Connector that holds the user credentials. projects/{project}/locations/{location}/accountConnectors/{account_connector_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gemini_code_tools_setting#account_connector GoogleGeminiCodeToolsSetting#account_connector}
	AccountConnector *string `field:"optional" json:"accountConnector" yaml:"accountConnector"`
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gemini_code_tools_setting#config GoogleGeminiCodeToolsSetting#config}
	Config interface{} `field:"optional" json:"config" yaml:"config"`
	// Overridden URI, if allowed by Tool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gemini_code_tools_setting#uri_override GoogleGeminiCodeToolsSetting#uri_override}
	UriOverride *string `field:"optional" json:"uriOverride" yaml:"uriOverride"`
}

